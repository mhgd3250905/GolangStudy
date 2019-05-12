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
	_,err=conn.Do("HMSet","user02","name","john","age",19)
	if err != nil {
		fmt.Println("HMSet err= ",err)
		return
	}



	var user01Strs []string
	user01Strs,err=redis.Strings(conn.Do("HMGet","user02","name","age"))
	if err!=nil {
		fmt.Println("HGetall err= ",err)
		return
	}

	for i,v:=range user01Strs  {
		fmt.Printf("r[%d]=%v\n",i,v)
	}

}
