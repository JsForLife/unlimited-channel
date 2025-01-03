package unlimitedchannel

import (
	"sync"
	"testing"
)

func TestNewUnLimitedChannel(t *testing.T) {
	ch := NewUnLimitedChannel[int]()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch.In() <- i
		}
		ch.Close()
	}()

	go func() {
		defer wg.Done()
		for val := range ch.Out() {
			t.Log(val)
		}
	}()

	wg.Wait()
}
