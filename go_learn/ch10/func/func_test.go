package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues()(int, int){
	return rand.Intn(10), rand.Intn(20)
}
func TestFn(t *testing.T)  {
	//a,b := returnMultiValues()
	//t.Log(a,b)
	//
	//c,_:= returnMultiValues()
	//t.Log(c)
	//
	//d,_:= returnMultiValues()
	//t.Log(d)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))

}

// 有点像装饰器模式
func timeSpent(inner func (op int) int) func (op int) int  {
	return func (n int) int{
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int{
	time.Sleep(time.Second * 1)
	return op
}

func sum(ops ...int) int {
	ret := 0
	for _,op := range  ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T){
	t.Log(sum(1,2,3,4))
	t.Log(sum(1,2,3,4,5))
}

func TestDefer(t *testing.T) {
	defer Clear() // Started 打印之后，Clear才会执行
	t.Log("Started")
}

func TestDefer2(t *testing.T) {
	defer Clear()
	t.Log("Started")
	panic("Fatal error") //defer 仍会执行     panic 代表程序异常中断。即便程序中断，defer里的程序依然会执行。
}

func Clear(){
	fmt.Println("Clear Resopurce")
}





