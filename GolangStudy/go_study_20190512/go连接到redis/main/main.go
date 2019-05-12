package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	//1.连接到redis
	conn,err:=redis.Dial("tcp","0.0.0.0:6379")
	if err != nil {
		fmt.Println("redis Dial err=",err)
		return
	}

	defer conn.Close()
	//通过go向redis写入数据
	//因为返回的r是interface{}
	//因为name对应的值为string，所以我们需要转换
	_,err=conn.Do("Set","name","tomjerry")
	if err != nil {
		fmt.Println("set err= ",err)
		return
	}

	r,err:=redis.String(conn.Do("Get","name"))
	if err != nil {
		fmt.Println("set err= ",err)
		return
	}


	fmt.Println("操作ok r=",r)
}
