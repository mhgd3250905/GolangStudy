package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	//输入下文件，file就是一个指针
	fmt.Printf("file = %v", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err = ", err)
	}
}
