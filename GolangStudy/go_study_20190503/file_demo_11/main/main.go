package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写一个函数接收两个文件路径
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file fail err =%v", err)
	}

	defer srcFile.Close()
	//通过srcFile,获取到Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	//通过dstFile获取到writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main() {
	srcFile := "D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/abc.txt"
	dstFile := "D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/copy/ccc.txt"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Printf("拷贝失败 err=%v", err)
	}
}
