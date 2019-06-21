package main

import (
	"GolangStudy/GolangStudy/go_study_20190621/modle"
	"GolangStudy/GolangStudy/go_study_20190621/modle/bookSet"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
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

func main() {
	conn := pool.Get()
	defer conn.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/spider/bookset/:key", func(c *gin.Context) {
		key := c.Param("key")
		start := c.DefaultQuery("start", "0")
		end := c.Query("end")

		booksArr, err := redis.Strings(conn.Do("lrange", key, start, end))

		books := make([]bookSet.Book, 0)
		for i, _ := range booksArr {
			bookStr := booksArr[i]
			book := bookSet.Book{}
			json.Unmarshal([]byte(bookStr), &book)
			books = append(books, book)
		}
		if err != nil {
			msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_FAILED,
				Error:   err.Error(),
				Data: "",
			}
			msgByte, err := json.Marshal(&msg)
			if err != nil {

			} else {
				c.String(400, string(msgByte))
			}
		} else {
			contentByte, err := json.Marshal(&books)
			fmt.Println(string(contentByte))
			if err != nil {

			} else {
				msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_SUCCESS,
					Error:   "",
					Data: books}
				fmt.Println(msg)
				msgStr, err := json.Marshal(&msg)
				fmt.Println(string(msgStr))
				if err != nil {

				} else {
					c.String(200, string(msgStr))
				}
			}
		}
	})
	r.Run()
}
