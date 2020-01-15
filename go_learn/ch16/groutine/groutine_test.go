package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T)  {
	for i:=0; i< 10; i++ {
		go func(i int) { // 报 打印数字的 函数 放到协程中 执行
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond *50)
}