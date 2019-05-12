package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64=3.4
	fmt.Println("type: ",reflect.TypeOf(x))
	v:=reflect.ValueOf(x)
	fmt.Println("v value:",v)

	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n",v.Interface())
	y:=v.Interface().(float64)
	fmt.Println(y)

	v2:=reflect.ValueOf(&x)
	*(v2.Interface().(*float64))=255.0
	fmt.Println("x value:",x)
}
