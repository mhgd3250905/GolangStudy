package main

import "testing"

func TestGetSub(t *testing.T) {
	//调用
	res:=getSub(3,5)
	if res != -2 {
		//fmt.Printf("AddUpper(10) 执行错误，期望值=%v,实际值=%v",55,res)
		t.Fatalf("getSub(3,5) 执行错误，期望值=%v,实际值=%v",-2,res)
	}

	//如果正确，输出日志
	t.Logf("getSub(3,5) 执行正确，期望值=%v,实际值=%v",-2,res)
}
