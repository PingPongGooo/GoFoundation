package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)


var LessThanTwoError = errors.New("n should be not less than 2")
var LargeThenHundreError = errors.New("n should be not larger than 100")

func GetFibonacci2(n int) ([]int,error) {
	if n < 2{
		return nil,LessThanTwoError
	}
	if n > 100{
		return nil,LargeThenHundreError
	}
	fibList := []int{1,1}
	for i := 2; i< n; i++{
		fibList = append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList,nil
}

func TestGetFibonacci(t *testing.T) {
	if v,err :=  GetFibonacci2(10); err !=nil{
		if  LessThanTwoError == err{
			fmt.Println("It is less.")
		}
		t.Error(err)
	}else {
		t.Log(v)
	}
}

func GetFibonacci3(str string)  {
	var (
		i int
		err error
		list []int
	)
	if i,err = strconv.Atoi(str); err!= nil{
		fmt.Println("Error", err)
		return
	}
	if list,err = GetFibonacci2(i); err!= nil{
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}