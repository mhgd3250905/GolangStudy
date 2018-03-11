package main

import(
	"fmt"
)

func main() {
	//数组：同意类型的集合
	var ids [50]int
	//数组操作，通过下标，从0开始，到len()-1
	for i := 0; i < len(ids); i++ {
		ids[i]=i+1
		fmt.Printf("id[%d]=%d\n",i,ids[i])
	}
}
