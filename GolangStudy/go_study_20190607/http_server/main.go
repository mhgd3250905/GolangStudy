package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/",Hello)
	http.HandleFunc("/user/login",Login)
	err:=http.ListenAndServe("0.0.0.0:8880",nil)
	if err != nil {
		fmt.Println("http_server listen failed,err=",err)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	w.Write([]byte("login"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w,"hello")
}
