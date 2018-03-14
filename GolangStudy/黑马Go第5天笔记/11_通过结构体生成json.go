package main

import (
	"encoding/json"
	"fmt"
)

//首字母必须大写，否则无法使用
//type IT struct {
//	Company  string
//	Subjects []string
//	IsOk     bool
//	Price    float64
//}

type IT struct {
	Company  string   `json:"-"`       //此字段不会输出到屏幕
	Subjects []string `json:"subject"` //二次编码
	IsOk     bool     `json:",string"` //转化为字符串再编码
	Price    float64  `json:",string"`
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"go", "c++", "java"}, true, 666.666}
	//根据内容生成json文本
	//buf,err:=json.Marshal(s)

	buf, err := json.MarshalIndent(s, "", " ")

	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}
