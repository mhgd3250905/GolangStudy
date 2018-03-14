package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("useage: xxx srcFile dstFile")
		return
	}
	//fmt.Println("list = ", list)
	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("目的文件和源文件名字不能相同！")
		return
	}

	//只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	//新建目的文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//操作完毕需要关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理，从源文件读取内容，往目的文件写，读多少写多少
	buf := make([]byte, 4*1024) //4k大小
	for {
		n, err3 := sF.Read(buf)
		if err3 != nil {
			fmt.Println("err3 = ", err3)
			if err3 == io.EOF {
				break
			}
		}
		//往目的文件写
		dF.Write(buf[:n])
	}

}
