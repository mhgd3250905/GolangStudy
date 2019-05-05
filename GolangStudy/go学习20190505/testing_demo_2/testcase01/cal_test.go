package main

import (
	"fmt"
	"testing"
)

//编写测试用例，去测试addUpper是否正确
func TestAddUpper(t *testing.T)  {
	//调用
	res:=addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v,实际值=%v",55,res)
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v,实际值=%v",55,res)
	}

	//如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确，期望值=%v,实际值=%v",55,res)
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello被调用...")
}
