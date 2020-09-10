### 基本信息

context package 为我们提供了以下关键能力：
- cancellation
    - 可重入性：After the first call, subsequent calls to a CancelFunc do nothing.
    - Calling the CancelFunc cancels the child and its children, removes the parent's
    reference to the child, and stops any associated timers.
    - NOTE: Failing to call the CancelFunc leaks the child and its children
- cancellation propagation
    - cancellation propagation 的实现依赖：
    - Context对象在函数之间的传递
    - 所有的Context对象构成了Tree Of Dependency
    - 父节点的Cancel动作会递归地传播给所有子节点
- pass value
    - 谨慎使用WithValue，要求：
    - value必须是Request级别，比如RequestID
    - 即使是Request级别的value，如果它很重要，你应该通过参数传递。提高代码的可读性

### 注意事项

Programs that use Contexts should follow these rules to keep interfaces
consistent across packages and enable static analysis tools to check **context propagation**:

1. Do not store Contexts inside a struct type; instead, pass a Context explicitly to
   each function that needs it. The Context should be the first parameter.
2. Do not pass a nil Context, even if a function permits it. Pass context.TODO
   if you are unsure about which Context to use.
3. Use context Values only for request-scoped data that transits processes and APIs,
   not for passing optional parameters to functions.
4. Contexts are safe for simultaneous use by multiple goroutines

### WithCancel

```go
func ExampleWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			case <-time.After(time.Second):
				fmt.Println("working")
			}
		}
	}(ctx)

	time.Sleep(2500 * time.Millisecond)
	cancel()
	// Output:
	// working
	// working
}
```

### WithTimeout

```go
func ExampleWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 1秒钟后cancel会自动执行
	// 可重入、协程安全的函数
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("work done")
	}

	// Output:
	// context deadline exceeded
}
```

### WithDeadline

````go
func ExampleWithDeadline() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	ch := make(chan int)
	go func(ctx context.Context) {
		select {
		case <-time.NewTimer(time.Millisecond * 1050).C:
			ch <- 1
		case <-ctx.Done():
			ch <- 0
		}
	}(ctx)

	fmt.Println(<-ch)
	// Output:
	// 0
}
````

### WithValue

```go
func ExampleWithValue() {
	ctx := context.WithValue(context.Background(), 1, 100)
	fmt.Println(ctx.Value(1))

	ctx2 := context.WithValue(ctx, 1, 200)
	fmt.Println(ctx.Value(1))
	fmt.Println(ctx2.Value(1))

	// Output:
	// 100
	// 100
	// 200
}
```