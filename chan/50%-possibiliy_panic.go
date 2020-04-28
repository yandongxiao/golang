package main

// close: nil(panic), closed(panic)
// write: nil(wait for ever), closed(panic)
// read: nil(wait for ever), closed(读取完chan中有效的值以后，开始读取到零值)
func main() {
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}: // panic if this case is selected.
	case <-c:
	}
}
