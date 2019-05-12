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
	//2.通过go向redis写入数据
	//因为返回的r是interface{}
	//因为name对应的值为string，所以我们需要转换
	_,err=conn.Do("HSet","user01","name","john")
	if err != nil {
		fmt.Println("set err= ",err)
		return
	}

	_,err=conn.Do("HSet","user01","age",10)
	if err != nil {
		fmt.Println("set err= ",err)
		return
	}

	//3 通过go向redis读取数据，string [key-val]
	r1,err:=redis.String(conn.Do("HGet","user01","name"))
	if err != nil {
		fmt.Println("HGet err= ",err)
		return
	}
	r2,err:=redis.Int(conn.Do("HGet","user01","age"))
	if err != nil {
		fmt.Println("HGet err= ",err)
		return
	}

	var user01Strs []string
	user01Strs,err=redis.Strings(conn.Do("HGetall","user01"))
	if err!=nil {
		fmt.Println("HGetall err= ",err)
		return
	}

	fmt.Printf("操作ok r1=%v,r2=%v,user01Strs=%v\n",r1,r2,user01Strs)
}
