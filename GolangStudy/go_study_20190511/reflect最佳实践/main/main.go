package main

import (
	"fmt"
	"reflect"
)

//定义一个结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

//显示
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("----end----")
}

//返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1+n2
}

//Set方法
func (s Monster)Set(name string,age int,score float32,sex string)  {
	s.Name=name
	s.Age=age
	s.Score=score
	s.Sex=sex
}

func TestStruct(a interface{}){
	//获取reflect.Type()
	typ:=reflect.TypeOf(a)
	//获取到a对应的value
	val:=reflect.ValueOf(a)
	//获取a对应的类别
	kd:=val.Kind()
	//如果V换入的不是struct，就退出
	if kd!=reflect.Struct{
		return
	}

	//获取到该结构体有几个字段
	num:=val.NumField()
	fmt.Printf("该结构体有%v个字段\n",num)

	//遍历结构体所有的字段
	for i := 0; i < num; i++{
		fmt.Printf("字段 %d 值为=%v\n",i,val.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取Tag的值
		tagVal:=typ.Field(i).Tag.Get("json")
		//如果该字段有标签,就显示，否则不显示
		if tagVal!="" {
			fmt.Printf("字段 %v tag为=%v\n",i,tagVal)
		}
	}

	//获取到这个结构体有多少个方法
	numOfMethod:=val.NumMethod()
	fmt.Printf("该结构体有%v个方法\n",numOfMethod)

	//获取到第二个方法，并调用它
	//函数方法的排序默认是按照函数名排序的
	val.Method(1).Call(nil)

	//调用结构体第一个方法Mothed(0)
	var params []reflect.Value//声明一个reflect.Value切片
	params=append(params, reflect.ValueOf(10))
	params=append(params,reflect.ValueOf(40))
	res:=val.Method(0).Call(params)
	fmt.Printf("res = %v\n",res[0].Int())
}

func main() {
	var a Monster=Monster{
		Name:"黄鼠狼精",
		Age:400,
		Score:30.0,
	}

	TestStruct(a)
}
