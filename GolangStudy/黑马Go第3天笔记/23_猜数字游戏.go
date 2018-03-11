package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randNum int
	//产生一个4位的随机数
	CreatNum(&randNum)
	//fmt.Println("randNum = ", randNum)

	randSlice := make([]int, 4)
	//保存这个4位数的每一位
	GetNum(randSlice, randNum)
	//fmt.Println("randSlice = ", randSlice)

	OnGame(randSlice)
}

func OnGame(randSlice []int) {
	var num int
	keySlice := make([]int, 4)

	for {
		for {
			fmt.Println("请输入一个4位数")
			fmt.Scan(&num)
			if num < 10000 && num > 999 {
				break
			} else {
				fmt.Println("输入的数不符合要求")
			}
		}

		GetNum(keySlice, num)
		//fmt.Println("keySlice = ",keySlice)

		n:=0
		for i := 0; i < 4; i++ {
			if randSlice[i]<keySlice[i] {
				fmt.Printf("第%d位大了一点\n",i+1)
			}else if randSlice[i]>keySlice[i] {
				fmt.Printf("第%d位小了一点\n",i+1)

			}else if randSlice[i]==keySlice[i] {
				fmt.Printf("第%d位猜对了\n",i+1)
				n++
			}
		}

		if n == 4 {//4位都猜对了
			fmt.Println("全部猜对了！！！")
			break//挑出循环
		}
	}

}

//获取每一位
func GetNum(s []int, num int) {
	s[0] = num / 1000       //取千位
	s[1] = num % 1000 / 100 //取百位
	s[2] = num % 100 / 10   //取十位
	s[3] = num % 10         //取十位
}

//获取随机数
func CreatNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000) //一定要是四位数
		if num >= 1000 {
			break
		}
	}
	//fmt.Println("num = ", num)

	*p = num
}
