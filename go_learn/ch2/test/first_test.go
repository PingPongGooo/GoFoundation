package try_test

import (
	"testing"
)

func TestFirstTry(t *testing.T) {
	t.Log("My first try!!")
}


// The master has failed more times than the beginner has tried.

// 编写测试程序
// 1. 源码文件以_test结尾：XXX_test.go
// 2. 测试方法名以Test 开头： func TestXXX(t *testing.T) {...}

// 大写的方法，代表包外可以访问