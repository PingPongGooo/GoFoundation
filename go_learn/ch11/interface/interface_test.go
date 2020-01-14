package _interface

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {  // 没有 使用implement 去和 programmer 绑定.

}

func (g *GoProgrammer) WriteHelloWorld() string{  // 方法签名看起来是一样的。 就是接口实现。  duck type
	return "Hello world"
}

func TestClient(t *testing.T)  {
	var p Programmer
	p =  new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
