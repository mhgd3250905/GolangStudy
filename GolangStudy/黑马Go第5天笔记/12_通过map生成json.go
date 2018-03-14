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
	//创建一个map
	m := make(map[string]interface{}, 4)

	m["company"] = "itcast"
	m["subject"] = []string{"go", "c++", "java"}
	m["isOk"] = true
	m["price"] = 666.666

	//编码成json
	result, err := json.MarshalIndent(m,"","	")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("result = ", string(result))
}