package main

import "fmt"

//冒泡排序
func bsortAsc(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

//排序算法
func ssort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] > a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

func main() {
	b := [...]int{8, 6, 5, 9, 12, 3, 1}
	bsortAsc(b[:])
	fmt.Println(b)
	ssort(b[:])
	fmt.Println(b)
}
