package main

import (
	"encoding/json"
	"fmt"
)

/*
monster序列化后={"name":"牛魔王","age":500,"birthday":"2011-11-11","sal":8000,"skill":"牛魔拳"}
a map序列化后={"address":"红云洞","age":30,"name":"红孩儿"}
slice序列化后=[{"address":"北京","age":7,"name":"jack"},{"address":"墨西哥","age":20,"name":"tom"}]

*/

type Monster struct {
	Name     string  `json:"name"` //反射机制
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func unmarshalStruct() {
	str := "{\"name\":\"牛魔王\",\"age\":500,\"birthday\":\"2011-11-11\",\"sal\":8000,\"skill\":\"牛魔拳\"}";
	//定义一个monster的实例
	var monster Monster
	err:=json.Unmarshal([]byte(str),&monster)
	if err != nil {
		fmt.Printf("反序列化失败 err=%v",err)
	}
	fmt.Printf("反序列化结构体为：%v\n",monster)
}

func unmarshalMap() {
	str:="{\"address\":\"红云洞\",\"age\":30,\"name\":\"红孩儿\"}"

	//定义一个map
	var a map[string]interface{}

	//反序列化
	err:=json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Printf("反序列化失败 err=%v",err)
	}
	fmt.Printf("反序列化后a = %v\n",a)
}

func unmarshalSlice() {
	str:="[{\"address\":\"北京\",\"age\":7,\"name\":\"jack\"},{\"address\":\"墨西哥\",\"age\":20,\"name\":\"tom\"}]"

	//定义一个切片
	var slice []map[string]interface{}
	//反序列化
	err:=json.Unmarshal([]byte(str),&slice)
	if err != nil {
		fmt.Printf("反序列化失败 err=%v",err)
	}
	fmt.Printf("反序列化后slice = %v\n",slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
