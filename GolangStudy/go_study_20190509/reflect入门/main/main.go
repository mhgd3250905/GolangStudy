package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取到传入的变量的type,kind,value
	//1.获取type
	rType:=reflect.TypeOf(b)
	fmt.Println("type = ",rType)

	//2.获取到reflect.Value
	rVal:=reflect.ValueOf(b)
	fmt.Printf("value= %v,value type = %T\n",rVal,rVal)
	//这个结果只是reflect.Value并不是真的int，无法完成int的操作
	fmt.Printf("value= %v,value type = %T\n",rVal.Int(),rVal.Int())

	//3.将rVal转成 interface{}
	iv:=rVal.Interface()
	//再将interface{}断言为我们需要的乐行
	num2:=iv.(int)
	fmt.Printf("num2= %v,num2 type = %T\n",num2,num2)
}

func main() {
	//1.定义一个int
	var num int =100
	reflectTest01(num)
}
