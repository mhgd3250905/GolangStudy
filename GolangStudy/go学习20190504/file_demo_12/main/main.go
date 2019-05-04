package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
创建exe go build -o test.exe main.go
*/


//定义一个结构体用于保存统计结果
type CharCount struct {
	ChCount    int //英文个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其他字符个数
}


func main() {
	//思路，打开文件，创建一个Reader
	//每读取一行，就去统计有多少个英文、数字、空格、和其他字符
	fileName := "D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/test.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历str,进行统计
		for _, v := range str {
			//fmt.Println(v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v >= ' ' || v <= '\t':
				count.SpaceCount++
			case v >= 0 && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	//输出统计的结果
	fmt.Printf("字符的个数为%v,数字的个数为%v，空格的个数为%v，其他个数为%v", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
