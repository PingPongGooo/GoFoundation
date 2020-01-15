package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string{
	time.Sleep(time.Millisecond * 50) // 耗时操作
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)  // 延迟输出信息
	fmt.Println("Task is done")
}

func TestService(t *testing.T)  {
	fmt.Println(service())  // 串行
	otherTask()
}

// 第一种 channel 通信双方阻塞等待
func AsyncService() chan string { // 返回 channel
	retCh := make(chan string) // 声明一个channel 指明 channel类型，放入string，就只能放string
	go func() { // 1. 启动协程
		ret := service()
		fmt.Println("returned result")
		retCh <- ret // 把结果放回到channel中
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T){
	retCh := AsyncService() // 另外一个任务。 放到另外一个协程中执行。
	otherTask()
	fmt.Println(<-retCh)  // 当我们需要结果的时候，去channel中把数据读出来。
}

// 第2种 channel buff 当前设置为 1
func AsyncService2() chan string { // 返回 channel
	retCh := make(chan string,1) // 声明一个channel 指明 channel类型，放入string，就只能放string  //
	go func() { // 1. 启动协程
		ret := service()
		fmt.Println("returned result")
		retCh <- ret // 把结果放回到channel中
		fmt.Println("service exited")
	}()
	return retCh
}

func TestAsyncService2(t *testing.T){
	retCh := AsyncService2() // 另外一个任务。 放到另外一个协程中执行。
	otherTask()
	fmt.Println(<-retCh)  // 当我们需要结果的时候，去channel中把数据读出来。
}