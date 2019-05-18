package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

//定义全局的pool
var pool *redis.Pool

func initPool(address string,maxIdle,maxActive int,idleTimeOut time.Duration){

	pool=&redis.Pool{
		MaxIdle:maxIdle,//最大空闲连接数
		MaxActive:maxActive,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout:idleTimeOut,//最大空闲时间单位：秒
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp",address)
		},
	}
}
