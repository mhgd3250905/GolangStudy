package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将abc.txt的文件内容导入到 kkk.txt
	
	//1.首先将 abc文件内容读取到内存，将读取到的内容写入到kkk文件
	filePath_1:="D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/abc.txt"
	filePath_2:="D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/kkk.txt"

	content,err:=ioutil.ReadFile(filePath_1)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("read file err= %v\n",err)
		return
	}

	err=ioutil.WriteFile(filePath_2,content,0666)
	if err != nil {
		fmt.Printf("write file err= %v\n",err)
	}
}
