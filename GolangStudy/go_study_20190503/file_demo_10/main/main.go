package main

import (
	"fmt"
	"os"
)

func main() {
	//判断文件是否存在
	filePath_1 := "D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/abc.txt"
	filePath_2 := "D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/qqq.txt"
	exist1, err1 := PathExist(filePath_1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(exist1)
	exist2, err2 := PathExist(filePath_2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(exist2)
}

//根据文件路径判断文件是否存在
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
