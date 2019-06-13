package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

type Person struct {
	Name string
	Age  int
}

type Result struct {
	output string
}

func (this *Result) Write(p []byte) (n int, err error) {
	fmt.Println("call by template")
	this.output+=string(p)
	return len(p),nil
}

func initTemplate(fileName string) (err error) {
	t, err = template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("parse file err =", err)
		return
	}
	return
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	//fmt.Fprint(w, "hello")


	m:=make(map[string]interface{})
	m["name"]="mary"
	m["age"]=18
	m["title"]="我的个人网站"

	var arr []Person
	p := Person{Name: "Mary", Age: 31}
	p2 := Person{Name: "Mary002", Age: 31}
	p3 := Person{Name: "Mary003", Age: 31}
	arr=append(arr,p)
	arr=append(arr,p2)
	arr=append(arr,p3)

	err:=t.Execute(w, arr)
	if err != nil {
		fmt.Println("template execute err = ",err)
	}


	////渲染到文件中
	//file, err := os.OpenFile("C:/Users/admin/go/src/GolangStudy/GolangStudy/go_study_20190613/template_http/test.txt",
	//	os.O_CREATE|os.O_WRONLY, 0755)
	//if err != nil {
	//	fmt.Println("open file err= ",err)
	//	return
	//}
	//t.Execute(file, p)

	////渲染到实现接口的结构体中
	//result:=&Result{}
	//t.Execute(result,p)
	//fmt.Println(result.output)

}

func main() {
	http.HandleFunc("/", userInfo)
	err := initTemplate("C:/Users/admin/go/src/GolangStudy/GolangStudy/go_study_20190613/template_http/index.html")
	if err != nil {
		fmt.Println("init template err= ", err)
	}

	err = http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http_server listen failed,err=", err)
	}

}
