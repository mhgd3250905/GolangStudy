package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	w.Write([]byte("login"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}

const form = `<html>
<body>
<form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form>
<body>
</html>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>Hello World</h1>")
	panic("test test")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		request.ParseForm()
		io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, request.FormValue("in"))
	}
}

/**
处理Panic错误
*/
func logPanics(handle func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caugh panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

func main() {
	http.HandleFunc("/", logPanics(Hello))
	http.HandleFunc("/user/login", logPanics(Login))
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http_server listen failed,err=", err)
	}

}
