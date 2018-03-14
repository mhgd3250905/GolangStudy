package main

import (
	"github.com/Unknwon/macaron"
	"encoding/json"
	"fmt"
)

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return getJson()
	})
	m.Run()
}

type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

func getJson() string {
	//创建一个map
	m := make(map[string]interface{}, 4)

	m["company"] = "itcast"
	m["subject"] = []string{"go", "c++", "java"}
	m["isOk"] = true
	m["price"] = 666.666

	//编码成json
	result, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		fmt.Println("err = ", err)
		return ""
	}

	resultStr:=string(result)

	fmt.Println(resultStr)

	return resultStr
}
