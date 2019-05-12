package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
)

type Monster struct {
	Name string `json:"name"`//反射机制
	Age int `json:"age"`
	Birthday string `json:"birthday"`
	Sal float64`json:"sal"`
	Skill string`json:"skill"`
}

func testStruct(){
	//演示
	monster:=Monster{
		Name:"牛魔王",
		Age:500,
		Birthday:"2011-11-11",
		Sal:8000.0,
		Skill:"牛魔拳",
	}
	//将monster序列化
	data,err:=json.Marshal(&monster)
	if err!=nil{
		fmt.Printf("序列化错误 err=%v",err)
	}
	//输出序列化结果
	fmt.Printf("monster序列化后=%v\n",string(data))
}

//将map进行序列化
func testMap() {
	//定义一个Map
	var a map[string]interface{}
	//使用map之前需要make
	a=make(map[string]interface{})
	a["name"]="红孩儿"
	a["age"]=30
	a["address"]="红云洞"

	//将map进行序列化
	data,err:=json.Marshal(a)
	if err!=nil{
		fmt.Printf("序列化错误 err=%v",err)
	}
	//输出序列化结果
	fmt.Printf("a map序列化后=%v\n",string(data))
}

//演示对切片进行序列化
func testSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1=make(map[string]interface{})
	m1["name"]="jack"
	m1["age"]=7
	m1["address"]="北京"
	slice=append(slice,m1)

	var m2 map[string]interface{}
	m2=make(map[string]interface{})
	m2["name"]="tom"
	m2["age"]=20
	m2["address"]="墨西哥"
	slice=append(slice,m2)

	//将slice进行序列化
	data,err:=json.Marshal(slice)
	if err!=nil{
		fmt.Printf("序列化错误 err=%v",err)
	}
	//输出序列化结果
	fmt.Printf("slice序列化后=%v\n",string(data))
}

//对基本类型序列化：序列化只是转化为一个字符串而已
func testFloat64(){
	var num1 float64=2345.6
	//将float64进行序列化
	data,err:=json.Marshal(num1)
	if err!=nil{
		fmt.Printf("序列化错误 err=%v",err)
	}
	//输出序列化结果
	fmt.Printf("float64序列化后=%v\n",string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
