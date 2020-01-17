package channel_close

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func service() string{
	time.Sleep(time.Millisecond * 500) // 耗时操作
	//time.Sleep(time.Millisecond * 50) // 耗时操作

	return "Done"
}

func dataProducer(ch chan int, wg *sync.WaitGroup) {

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg
	}()

}

func TestSelect(t *testing.T)  {

	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
