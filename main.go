package unlimitedchannel

type UnLimitedChannel[T any] struct {
	in, out chan T
}

func NewUnLimitedChannel[T any]() *UnLimitedChannel[T] {
	in, out := make(chan T), make(chan T)
	go func() {
		buffer := []T{}
		for {
			// 如果缓冲区没有消息，则尝试读取写入通道。这一步主要防止切割空缓冲区导致 panic。
			if len(buffer) == 0 {
				e, ok := <-in
				// 如果写入通道已经被关闭，则将读出通道也关闭
				if !ok {
					close(out)
					return
				}
				// 能读到消息则放入缓冲区
				buffer = append(buffer, e)
			} else {
				// 缓冲区不为空，尝试将数据发送到输出通道
				select {
				case out <- buffer[0]:
					buffer = buffer[1:]
				// 继续读取写入通道，避免因输出通道停止接收或消费堆积导致阻塞
				case e, ok := <-in:
					if !ok {
						// 清空缓冲区后再关闭
						for _, b := range buffer {
							out <- b
						}
						close(out)
						return
					}
					buffer = append(buffer, e)
				}
			}
		}
	}()

	return &UnLimitedChannel[T]{
		in:  in,
		out: out,
	}
}

func (ch *UnLimitedChannel[T]) In() chan<- T {
	return ch.in
}

func (ch *UnLimitedChannel[T]) Out() <-chan T {
	return ch.out
}

func (ch *UnLimitedChannel[T]) Close() {
	close(ch.in)
}
