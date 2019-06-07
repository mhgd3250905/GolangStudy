package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

//websocket 接收数据并且全部大写之后返回
func upper(ws *websocket.Conn) {
	for{
		var replay string
		if err:=websocket.Message.Receive(ws,&replay);err != nil {
			log.Printf("websocket receive failed:%v/n",err)
			continue
		}

		if err := websocket.Message.Send(ws, strings.ToUpper(replay));err!=nil {
			log.Printf("websocket send failed:%v/n",err)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	t,err:=template.ParseFiles("E:/go/GoStudy_001/websocket_demo/index.html")
	if err!=nil{
		log.Printf("template parsefiles failed:%v/n",err)
	}
	t.Execute(w,nil)
}

func main() {
	http.Handle("/upper",websocket.Handler(upper))
	http.HandleFunc("/",index)

	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Printf("http_server listenAndServe failed:%v/n",err)
		os.Exit(1)
	}
}
