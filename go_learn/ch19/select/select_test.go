package _select

import (
"fmt"
"testing"
"time"
)

func service() string{
	time.Sleep(time.Millisecond * 500) // 耗时操作
	//time.Sleep(time.Millisecond * 50) // 耗时操作

	return "Done"
}

func AsyncService() chan string { // 返回 channel
	retCh := make(chan string,1)
	go func() { // 1. 启动协程
		ret := service()
		fmt.Println("returned result")
		retCh <- ret // 把结果放回到channel中
		fmt.Println("service exited")
	}()
	return retCh
}

func TestSelect(t *testing.T)  {

	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
