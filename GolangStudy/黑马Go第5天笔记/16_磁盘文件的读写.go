package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"bufio"
)



func main() {
	path:="./demo.txt"

	//WriteFile(path)
	//ReadFile(path)
	ReadFileLine(path)
}

//每次读取一行
func ReadFileLine(path string){
	//打开文件
	f,err:=os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先换在缓冲区
	r:=bufio.NewReader(f)

	//遇到\n就结束读取，但是'\n'也读取了进来
	for{
		buf,err:=r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {//文件已经结束
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("buf = #%s#\n",string(buf))
	}
}


func WriteFile(path string) {
	//打开文件，新建文件
	f,err:=os.Create(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	//使用完毕，需要关闭文件
	defer f.Close()

	var buf string

	for i := 0; i < 10; i++ {
		buf=fmt.Sprintf("i = %d\n",i)
		n,err:=f.WriteString(buf)
		if err!=nil{
			log.Fatal(err)
			return
		}
		fmt.Printf("写了%d个字符",n)
	}
}

func ReadFile(path string)  {
	//打开文件
	f,err:=os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	//关闭文件
	defer f.Close()

	buf:=make([]byte,1024*2) //2k大小

	n,err:=f.Read(buf)
	if err != nil && err!=io.EOF{//文件出错 同时没到结尾
		log.Fatal(err)
		return
	}
	fmt.Printf("读了%d个字符\n",n)
	fmt.Println("buf = ",string(buf[:n]))


}
