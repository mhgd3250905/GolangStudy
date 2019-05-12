package main

import "fmt"

func main() {
	var a [10]int

	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}

	for index,value:=range a{
		fmt.Printf("a[%d]=%d\n",index,value)
	}
}
