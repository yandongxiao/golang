package merge

func merge(ch1, ch2 <-chan int) <-chan int {
	if ch1 == nil && ch2 == nil {
		ch := make(chan int)
		close(ch)
		return ch
	} else if ch1 == nil {
		return ch2
	} else if ch2 == nil {
		return ch1
	}

	output := make(chan int)
	oldCh1, oldCh2 := ch1, ch2
	go func() {
		var v1, v2 *int
		var ch1Closed, ch2Closed bool

		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1Closed = true
				} else {
					v1 = &v
				}
				ch1 = nil
			case v, ok := <-ch2:
				if !ok {
					ch2Closed = true
				} else {
					v2 = &v
				}
				ch2 = nil
			}

			// 控制信号和数据处理必须在一起!
			if ch1Closed && ch2Closed {
				close(output)
				break
			} else if ch1Closed {
				if v2 != nil {
					output <- *v2
					v2 = nil
				}
				ch2 = oldCh2
			} else if ch2Closed {
				if v1 != nil {
					output <- *v1
					v1 = nil
				}
				ch1 = oldCh1
			} else {
				if v1 != nil && v2 != nil {
					if *v1 <= *v2 {
						output <- *v1
						v1 = nil
						ch1 = oldCh1
					} else {
						output <- *v2
						v2 = nil
						ch2 = oldCh2
					}
				}
			}
		}
	}()
	return output
}
