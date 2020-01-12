package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T)  {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt = MyInt(b)
	t.Log(a,b,c)
}

func TestPoint(t *testing.T)  {
	a:=1
	aPtr := &a
	t.Log(a,aPtr)
	t.Logf("%T %T", a,aPtr)
	//aPtr = aPtr + 1 //会编译报错。go 不支持 指针运算。
}

func  TestString(t *testing.T)  {
	var s string // 默认值的初始化值为空字符串 ，而不是nil
	t.Log("*"+s+"*")
	t.Log(len(s))
}



