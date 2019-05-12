package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位
	file := "D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Sprintf("read file err = %v", err)
	}
	//把读取到的内容显示到终端
	fmt.Printf("%v\n", content)//[]byte
	//因为没有显式的Open文件，因此也不需要显式的Close
	//文件的Open和Close已经被封装到ReadFiled的函数内部
	//因策这种方式就只能读取长度较短的文件
	fmt.Printf("%v\n",string(content))
}
