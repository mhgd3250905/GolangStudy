package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	age string
}

func main() {
	t,err:=template.ParseFiles("D:/index_room_1_1.html")
	if err!=nil{
		fmt.Println("parse file err: ",err)
		return
	}
	p:=Person{Name:"Mary",age:"31"}
	if err := t.Execute(os.Stdin, p);err!=nil {
		fmt.Println("There was an error",err.Error())
	}
}
