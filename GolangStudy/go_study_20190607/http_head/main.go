package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url=[]string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {
	for _,v:= range url{
		//设置超时
		c:=http.Client{
			Transport:&http.Transport{
				Dial: func(network, addr string) (conn net.Conn, e error) {
					timeout:=time.Second*2
					return net.DialTimeout(network,addr,timeout)
				},
			},
		}
		resp,err:=c.Head(v)
		if err!=nil{
			fmt.Println("head %s failed,err=%v\n",v,err)
			continue
		}
		fmt.Printf("head succ,status:%v\n",resp.Status)
	}
}
