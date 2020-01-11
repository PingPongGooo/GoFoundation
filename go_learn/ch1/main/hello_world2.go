package main
// 包， 表明代码所在的模块

import (
	"fmt"
	"os"
)

// 功能实现
func main() {

	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}

	os.Exit(0)
}

// 应用程序入口

//1. 必须是 main 包： package main
//2. 必须是 main 方法： func main()
//3. 文件名不一定是main.go

// 退出返回值
// 与其他主要编程语言的差异
// Go 中main函数不支持任何返回值
// 通过os.Exit来返回状态

//  获取命令行参数
// 与其他主要编程语言的差异
// main函数不支持传入参数
// 在程序中直接通过os.Args获取命令行参数



