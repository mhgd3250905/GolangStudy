package main

import (
	"fmt"
)

func main() {
	arr:=[]int{0,1,2,3,4,5,6,7,8,9}
	//[low:high:max]取下标从low开始的元素，len=high-low,cap=max-low
	s1:=arr[:]//[0:len(arr):len(arr)]
	fmt.Println("s1= ",s1)
	fmt.Printf("len= %d,cap= %d\n",len(s1),cap(s1))

	//操作某个元素
	data:=arr[0]
	fmt.Println("data = ",data)

	s2:=arr[3:6:7]//a[3],a[4],a[5] len=6-3=3  cap=7-3=4
	fmt.Println("s2= ",s2)
	fmt.Printf("len= %d,cap= %d\n",len(s2),cap(s2))

	s3:=arr[:6]//从0开始取6个 容量也是6
	fmt.Println("s3= ",s3)
	fmt.Printf("len= %d,cap= %d\n",len(s3),cap(s3))

	s4:=arr[3:]//从3开始到结尾
	fmt.Println("4= ",s4)
	fmt.Printf("len= %d,cap= %d\n",len(s4),cap(s4))

}