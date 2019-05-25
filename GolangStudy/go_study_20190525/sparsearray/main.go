package main

import "fmt"

type ValNode struct {
	row    int
	column int
	val    interface{}
}

func main() {
	//1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //蓝子

	//输出看看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//转成稀疏数组
	//1.遍历chessMap 如果有一个元素的值！=0，就创建一个node节点
	//2.将其放入到对应的切片中即可

	var sparseArr []ValNode

	//标准的一个稀疏数组应该还有一个记录元素的二维数组规模（行和列 默认值）
	valNode := ValNode{
		row:    11,
		column: 11,
		val:    0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个valNode节点
				valNode := ValNode{
					row:    i,
					column: j,
					val:    v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前的稀疏数组是...")
	for i,valNode:=range sparseArr{
		fmt.Printf("%d: %d %d %d\n",i,valNode.row,valNode.column,valNode.val)
	}

	//将这个稀疏数组，存盘

	//如何恢复原始的数组

	//1.打开文件，恢复数组

	//先创建一个原始数组
	var chessMap2 [11][11]int

	//遍历稀疏数组
	for i,valNode:=range sparseArr{
		if i==0 {
			continue
		}
		chessMap2[valNode.row][valNode.column]=valNode.val.(int)
	}

	//看看chessMap2是不是恢复了
	fmt.Println("恢复后的原始数据是...")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
