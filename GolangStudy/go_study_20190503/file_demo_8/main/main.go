package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开一个存在的文件在文件内容后面追加内容
	//创建一个新文件，写入内容 5句 "hello world"
	//1.打开一个已经存在的文件 d:/abc.txt
	filePath:="D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/abc.txt"
	//os.O_APPEND：追加内容 os.O_RDWR:读写方式
	file,err:=os.OpenFile(filePath,os.O_RDWR|os.O_APPEND,0666)
	if err != nil {
		fmt.Printf("open file err = %v",err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	//先读取文件内容，显示在终端
	reader:=bufio.NewReader(file)
	for{
		str,err:=reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//显示到终端
		fmt.Print(str)
	}

	//准备写内容
	str:="ABC ENGLISH\r\n"
	//写入时使用带缓存的*Writer
	writer:=bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用WriterString方法时
	//其实内容是先写入到缓存，所以需要调用Flush方法将缓存的数据真正写入到文件中
	//否则文件中会没有数据
	writer.Flush()
}