package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/test.txt")
	if err != nil {
		fmt.Println("open file err = ", err)
	}
	//当函数退出时，要及时关闭file
	//否则会有内存泄漏
	defer file.Close()

	//创建一个*Reader，带缓冲的
	//默认缓冲区为4096
	reader:=bufio.NewReader(file)
	//循环读取文件内容
	for{
		//读到一个换行符就跳转下一个循环
		str,err:=reader.ReadString('\n')
		if err==io.EOF{//io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}

	fmt.Println("文件读取结束...")
}
