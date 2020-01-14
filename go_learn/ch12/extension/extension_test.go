package extension

import (
	"fmt"
	"testing"
)

type Pet struct {

}

func (p *Pet) Speak(){
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string){
	p.Speak()
	fmt.Println(" ",host)
}

type Dog struct {
	Pet   //
}

func TestDog(t *testing.T){
	dog := new(Dog)
	dog.SpeakTo("Chao")
	dog.Speak()  // 好像获得了继承的感觉
}

