package main

import (
	"GolangStudy/GolangStudy/go_study_20190617/collectors/chule"
	"GolangStudy/GolangStudy/go_study_20190617/collectors/huxiu"
	"GolangStudy/GolangStudy/go_study_20190617/collectors/ifanr"
	"github.com/gomodule/redigo/redis"
	"os"
)

//定义一个全局的pool
var pool *redis.Pool


func init() {

	pool = &redis.Pool{
		MaxIdle:     0,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间单位：秒
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func push2RedisList(c redis.Conn, key string, content string) (err error) {
	_, err = c.Do("RPUSH", key, content)
	return
}

var conn redis.Conn

func main() {
	conn = pool.Get()
	defer conn.Close()

	chule.ChuleSpider(conn)
	huxiu.HuxiuSpider(conn,OnHuxiuSpiderFinish)

}

func OnHuxiuSpiderFinish(){
	ifanr.ChuleSpider(conn,OnIfanrSpiderFinish)
}

func OnIfanrSpiderFinish(){
	os.Exit(0)
}


