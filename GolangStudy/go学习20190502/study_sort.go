package main

import (
	"fmt"
	"sort"
)

func main() {
	//testMapSort()
	testMapReverse()
}

func testMapReverse() {
	var a map[string]int
	var b map[int] string

	a=make(map[string]int)
	b=make(map[int]string)

	a["abc"]=101
	a["efg"]=10

	for k,v:=range a{
		b[v]=k
	}

	fmt.Println(b)

}

func testMapSort(){
	var a map[int]int
	a=make(map[int]int,5)

	a[0]=10
	a[3]=10
	a[2]=10
	a[1]=10
	a[18]=10

	var keys []int
	for k,v:=range a{
		//keys=append(keys,k)
		fmt.Println(k,v)
	}

	sort.Ints(keys)

	for _,v:=range keys{
		fmt.Println(v,a[v])
	}
	//if a[0]==nil {
	//	a[0]=make(map[int]int)
	//}
	//a[0][10]=10
	//fmt.Println(a)
}

