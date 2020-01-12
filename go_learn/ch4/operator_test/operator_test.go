package operator_test

import "testing"

const(
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T)  {
	a:=[...]int{1,2,3,4}
	b:=[...]int{1,3,4,5}
	//c:=[...]int{1,2,3,4,5}
	d:=[...]int{1,2,3,4}
	t.Log(a == b)
	//t.Log(a == c) // 编译错误
	t.Log(a == d)
}

func TestBitClear(t *testing.T){
	a:=7 // 0111
	a = a&^Readable
	b := a &^Executable
	t.Log(a&Readable == Readable,a&Writable == Writable,a&Executable == Executable)
	t.Log(b&Readable == Readable,b&Writable == Writable,b&Executable == Executable)
}

//位运算符
//与其他主要编程语言的差异，Go语言为了提高生产力
//&^ 按位置零
//
//1 &^ 0  -- 1
//1 &^ 1  -- 0
//0 &^ 1  -- 0
//0 &^ 0  -- 0

