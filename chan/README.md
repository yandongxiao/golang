data races
One of the duties in concurrent programming is to control resource sharing among concurrent computations, so that data races will not happen.
The ways to achieve this duty are called concurrency synchronization, or data synchronization.
Go supports several data synchronization techniques. Below will introduce one of them, channel.

Most operations in Go are not synchronized. In other words, they are not concurrency-safe.
These operations include value assignments, argument passing and container element manipulations, etc.
There are only a few operations which are synchronized, including the several to be introduced channel operations below.

One suggestion (made by Rob Pike) for concurrent programming is don't (let computations) communicate by sharing memory, (let them) share memory by communicating (through channels).

chan T
chan<- T
<-chan T

All channel types are comparable types.

close(ch)
ch <- v
<-ch
len(ch)
cap(ch)
All these operations are already synchronized, so no further synchronizations are needed to safely perform these operations.
However, like most other operations in Go, channel value assignments are not synchronized.
Similarly, assigning the received value to another value is also not synchronized


Operation	A Nil Channel	A Closed Channel	A Not-Closed Non-Nil Channel
Close	panic	panic	succeed to close (C)
Send Value To	block for ever	panic	block or succeed to send (B)
Receive Value From	block for ever	never block (D)	block or succeed to receive (A)

* Closing a nil or a already closed channel produces a panic in the current goroutine.
*  Sending a value to a closed channel also produces a panic in the current goroutine.
*  Sending a value to or receiving a value from a nil channel makes the current goroutine enter and stay in blocking state for ever.


Channel operation case C: when a goroutine tries to close a not-closed non-nil channel, once the goroutine has acquired the lock of of the channel, both of the following two steps will be performed by the following order.
1.  If the receiving goroutine queue of the channel is not empty, in which case the value buffer of the channel must be empty, all the goroutines in the receiving goroutine queue of the channel will be unshifted one by one,
each of them will receive a zero value of the element type of the channel and be resumed to running state.
2.  If the sending goroutine queue of the channel is not empty, all the goroutines in the sending goroutine queue of the channel will be unshifted one by one and
each of them will produce a panic for sending on a closed channel. The values which have been already pushed into the value buffer of the channel are still there.

In the above explanations, if a goroutine is unshifted out of a queue (either the sending goroutine queue or the receiving goroutine queue) of a channel, and the goroutine was blocked for being pushed into the queue at a select control flow code block, then the goroutine will be resumed to running state at the step 9 of the select control flow code block execution. It may be dequeued from the corresponding goroutine queues of several channels involved in the select control flow code block.
If the channel is unbufferd, then at any time, generally one of its sending goroutine queue and the receiving goroutine queue must be empty, but with an exception that a goroutine may be pushed into both of the two queues when executing a select control flow code block.
https://go101.org/article/channel-use-cases.html

Note, a channel is referenced by all the goroutines in either the sending or the receiving goroutine queue of the channel, so if neither of the two queues of the channel is empty, the channel will not be garbage collected for sure.
On the other hand, if a goroutine is blocked and stays in either the sending or the receiving goroutine queue of a channel, then the goroutine will also not be garbage collected for sure, even if the channel is referenced only by this goroutine. In fact, a goroutine can be only garbage collected when it has already exited.

Channel send operations and receive operations are simple statements.
for ; y < (1 << 63); c <- y {
for x, ok := <-c; ok; x, ok = <-c {

select-case Control Flow Code Blocks
    No expressions and statements are allowed to follow the select keyword (before {).
    No fallthrough statements are allowed to be used in case branches.
    Each statement following a case keyword in a select-case code block must be either a channel receive operation or a channel send operation statement.
    In case of there are some non-blocking case operations, Go runtime will randomly select one of them to execute, then continue to execute the corresponding case branch.
    In case of all the case operations in a select-case code block are blocking operations, the default branch will be selected to execute if the default branch is present. If the default branch is absent, the current goroutine will be pushed into the corresponding sending goroutine queue or receiving goroutine queue of every channel involved in all case operations, then enter blocking state.

Select Mechanism:
    1. evaluate all involved channel expressions and value expressions to be potentially sent in case operations, from top to bottom and left to right. Destination values for receive operations (as source values) in assignments needn't to be evaluated at this time.
    2. randomize the branch orders for polling in step 5. The default branch is always put at the last position in the result order. Channels may be duplicate in the case operations.
    3. sort all involved channels in the case operations to avoid deadlock in the next step. No duplicate channels stay in the first N channels of the sorted result, where N is the number of involved channels in the case operations. Below, the channel lock order is a concept for the first N channels in the sorted result.
    4. lock (a.k.a., acquire the locks of) all involved channels by the channel lock order produced in last step.
    5. poll each branch in the select block by the randomized order produced in step 2:
        1. if this is a case branch and the corresponding channel operation is a send-value-to-closed-channel operation, unlock all channels by the inverse channel lock order and make the currrent goroutine panic. Go to step 12.
        2. if this is a case branch and the corresponding channel operation is non-blocking, perform the channel operation and unlock all channels by the inverse channel lock order, then execute the corresponding case branch body. The channel operation may wake up another goroutine in blocking state. Go to step 12.
        3. if this is the default branch, then unlock all channels by the inverse channel lock order and execute the default branch body. Go to step 12.
    (Up to here, the default branch is absent and all case operations are blocking operations.)
    6. push (enqueue) the current goroutine (along with the information of the corresponding case branch) into the receiving or sending goroutine queue of the involved channel in each case operation. The current goroutine may be pushed into the queues of a channel for multiple times, for the involved channels in multiple cases may be the same one.
    7. make the current goroutine enter blocking state and unlock all channels by the inverse channel lock order.
    8. ..., in blocking state, waiting other channel operations to wake up the current goroutine, ...
    9. the current goroutine is waken up by another channel operation in another goroutine. The other operation may be a channel close operation or a channel send/receive operation. If it is a channel send/receive operation, there must be a case channel receive/send operation (in the current being explained select-case block) cooperating with it (by transferring a value). In the cooperation, the current goroutine will be dequeued from the receiving/sending goroutine queue of the channel.
    10. lock all involved channels by the channel lock order.
    11. dequeue the current goroutine from the receiving goroutine queue or sending goroutine queue of the involved channel in each case operation,
        1. if the current goroutine is waken up by a channel close operation, go to step 5.
        2. if the current goroutine is waken up by a channel send/receive operation, the corresponding case branch of the cooperating receive/send operation has already been found in the dequeuing process, so just unlock all channels by the inverse channel lock order and execute the corresponding case branch.
    12. Done

    if the current goroutine is waken up by a channel send/receive operation, the corresponding case branch of the cooperating receive/send operation has already been found in the dequeuing process, so just unlock all channels by the inverse channel lock order and execute the corresponding case branch.
    when a goroutine being blocked at a select-case code block gets resumed later, it will be removed from all the sending goroutine queues and the receiving goroutine queues of every channels involved in the channel operations followed case keywords in the select-case code block.
