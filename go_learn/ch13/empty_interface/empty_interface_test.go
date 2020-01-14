package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomingthing(p interface{}) {
	//if i, ok:=p.(int); ok{
	//	fmt.Println("Integer",i)
	//	return
	//}
	//
	//if s, ok:=p.(string); ok{
	//	fmt.Println("string",s)
	//	return
	//}
	//
	//fmt.Println("unknow Type")
	switch v:=p.(type) {
	case int:
		fmt.Println("Integer",v)
	case string:
		fmt.Println("string",v)
	default:
		fmt.Println("unknow type")
	}
}

func TestEmptyinterfaceAssertion(t *testing.T){
	DoSomingthing(10)
	DoSomingthing("10")
}