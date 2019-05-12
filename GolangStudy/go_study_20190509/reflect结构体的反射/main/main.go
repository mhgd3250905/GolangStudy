package main

import (
	"fmt"
	"reflect"
)

func reflectTest02(b interface{}) {
	//通过反射获取到传入的变量的type,kind,value
	//1.获取type
	rType:=reflect.TypeOf(b)
	fmt.Println("type = ",rType)

	//2.获取到reflect.Value
	rVal:=reflect.ValueOf(b)
	fmt.Printf("value= %v,value type = %T\n",rVal,rVal)

	//获取变量对应的kind
	fmt.Printf("rVal kind =%v,rType kind =%v\n",rVal.Kind(),rType.Kind())

	//3.将rVal转成 interface{}
	iv:=rVal.Interface()
	//再将interface{}断言为我们需要的类型
	//虽然这时候获取到的反射类型就已经是Student类型了，但是编译阶段无法调用其内部的方法
	//所以还是断言类型之后再使用
	fmt.Printf("iv= %v,iv type = %T\n",iv,iv)

	num2:=iv.(Student)
	//可以使用switch来进行类型判断
	fmt.Printf("num2= %v,num2 type = %T\n",num2,num2)
}

type Student struct {
	Name string
	Age int
}

func main() {
	stu:=Student{
		Name:"tom",
		Age:18,
	}
	reflectTest02(stu)
}
