package testing_demo_3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int 
	Skill string
} 

//绑定方法store:序列化病保存到文件中
func (this *Monster) Store() bool{
	//序列化
	data,err:=json.Marshal(this)
	if err != nil {
		fmt.Printf("marshal err=%v/n",err)
		return false
	}
	//保存到文件
	filePath:="D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/testing.txt"
	err=ioutil.WriteFile(filePath,data,0666)
	if err != nil {
		fmt.Printf("writeFile err=%v\n",err)
		return false
	}
	return true
}

//绑定Restore方法 从文件中读取，并反序列化为结构体对象
func (this *Monster)Restore() bool{
	//读取序列化后的字符串
	filePath:="D:/GIT/GoProject/src/GolangStudy/GolangStudy/files/testing.txt"
	data,err:=ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file err=%v",err)
		return false
	}

	//反序列化
	err=json.Unmarshal(data,this)
	if err != nil {
		fmt.Printf("Unmarshal err=%v",err)
		return false
	}
	return true
}
