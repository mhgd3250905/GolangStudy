package main

import (
	"fmt"
	"regexp"
)

func main() {

	buf := "abc azc a7c aac 888 a9c tac  "

	//1 解释规则 解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`a.c`)
	if reg1 == nil {
		fmt.Println("err = ", reg1)
		return
	}

	//2 根据规则提取关键信息
	result1:=reg1.FindAllStringSubmatch(buf,-1)
	fmt.Println("result1 = ",result1)
}
