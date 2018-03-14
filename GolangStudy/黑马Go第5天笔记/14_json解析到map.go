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



func main() {
	jsonBuf := `{
			"company": "itcast",
			"isOk": true,
			"price": 666.666,
			"subject": [
			"go",
			"c++",
			"java"
			]
			}`
	var tmp map[string]interface{}

	err:=json.Unmarshal([]byte(jsonBuf),&tmp)
	if err!=nil {
		fmt.Println("err = ",err)
		return
	}
	fmt.Println("tem = ",tmp)

}
