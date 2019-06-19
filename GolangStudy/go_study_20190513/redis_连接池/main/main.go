package main

import (
	"GolangStudy/GolangStudy/go_study_20190617/modles/bookSet"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool

func init() {

	pool=&redis.Pool{
		MaxIdle:8,//最大空闲连接数
		MaxActive:0,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout:100,//最大空闲时间单位：秒
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp","localhost:6379")
		},
	}
}

func main() {
	//先从pool中取出一个连接
	conn:=pool.Get()
	defer conn.Close()

	_,err:=conn.Do("Set","name","tom cat")
	if err != nil {
		fmt.Println("conn.Do err=",err)
	}

	//取出
	r,err:=redis.String(conn.Do("Get","name"))
	if err != nil {
		fmt.Println("conn.Do err=",err)
	}

	fmt.Println("r=",r)

	//如果我们要从pool中取出链接 一定保证连接池是没有关闭
	pool.Close()
	//conn2:=pool.Get()//这个时候获取到的conn使用时会报错


	arr,_:=redis.Strings(conn.Do("LRANGE","book","0","5"))

	for i,_:= range arr  {
		book:=bookSet.Book{}
		err=json.Unmarshal([]byte(arr[i]),&book)
		if err != nil {
			fmt.Println("json.Unmarshal failed,err= ",err)
			continue
		}
		fmt.Println(book)
	}

	//fmt.Println(arr)


}
