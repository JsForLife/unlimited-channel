package unlimitedchannel

import (
	"sync"
	"testing"
)

func TestNewUnLimitedChannel(t *testing.T) {
	in, out := NewUnLimitedChannel[int]()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		defer wg.Done()
		for val := range out {
			t.Log(val)
		}
	}()

	wg.Wait()
}
