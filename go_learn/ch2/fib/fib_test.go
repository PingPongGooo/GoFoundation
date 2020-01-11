package fib

import (
	"testing"
)

// var a int // 全局变量

func TestFibList(t *testing.T) {
	// var a int = 1 // 声明变量不使用，编译会报错。
	// var b int = 1

	// var (
	// 	a int = 1
	// 	b     = 1
	// )

	// a = 1
	a := 1
	b := 1

	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)

		tmp := a
		a = b
		b = tmp + a

	}
	t.Log()
}

func TestExchange(t *testing.T) {
	a := 2
	b := 3

	t.Log(a, b)
	a, b = b, a

	t.Log(a, b)
}

// // 变量赋值
// 与其他主要编程语言的差异
// 赋值可以自动类型推断
// 在一个赋值语句中可以对多个变量进行同时赋值
