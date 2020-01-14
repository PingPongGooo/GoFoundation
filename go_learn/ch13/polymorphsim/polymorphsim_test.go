package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {  // 没有 使用implement 去和 programmer 绑定.
}

func (p *GoProgrammer) WriteHelloWorld() Code{
	return "hello world go"
}

type JavaProgrammer struct {  // 没有 使用implement 去和 programmer 绑定.
}


func (p *JavaProgrammer) WriteHelloWorld() Code{
	return "hello world java"
}

func writeFirstProgram(p Programmer)  {
	fmt.Printf("%T %v\n",p,p.WriteHelloWorld())
}


func TestPolymorphism(t *testing.T)  {

	goProg :=  new(GoProgrammer)
	javaProg :=  new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
