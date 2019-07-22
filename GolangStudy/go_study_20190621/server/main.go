package main

import (
	"GolangStudy/GolangStudy/go_study_20190621/modle"
	"GolangStudy/GolangStudy/go_study_20190621/modle/bookSet"
	"GolangStudy/GolangStudy/go_study_20190621/modle/huxiu"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
)

const (
	KEY_BOOK_DETAIL_IN_REDIS  = "book_detail"
	KEY_BOOK_IN_REDIS         = "book"
	KEY_HUXIU_IN_REDIS        = "huxiu"
	KEY_CHULE_IN_REDIS        = "chule"
	KEY_HUXIU_DETAIL_IN_REDIS = "huxiu_detail"
	KEY_CHULE_DETAIL_IN_REDIS = "chule_detail"
)

//定义一个全局的pool
var pool *redis.Pool
var conn redis.Conn

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
	conn = pool.Get()
	defer conn.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/spider/bookset/:key", getBooks)
	r.GET("/spider/huxiu/:key", getHuxius)
	r.GET("/spider/detail/:mapKey/:key", getDetail)
	r.Run(":8880")
}

func getBooks(c *gin.Context) {
	key := c.Param("key")
	start := c.DefaultQuery("start", "0")
	end := c.Query("end")

	//获取结果
	result, err := redis.Strings(conn.Do("lrange", key, start, end))

	if err != nil {
		msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_FAILED,
			Error: err.Error(),
			Data:  "",
		}
		msg.Send(c)
	} else {
		//反序列化到数组中
		if key == KEY_BOOK_IN_REDIS {
			books := make([]bookSet.Book, 0)
			for i, _ := range result {
				bookStr := result[i]
				book := bookSet.Book{}
				json.Unmarshal([]byte(bookStr), &book)
				books = append(books, book)
			}
			//设置到消息类中
			msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_SUCCESS,
				Error: "",
				Data:  books}
			msg.Send(c)
		} else if key == KEY_BOOK_DETAIL_IN_REDIS {
			bookDetails := make([]bookSet.BookDetail, 0)
			for i, _ := range result {
				bookDetailStr := result[i]
				bookDetail := bookSet.BookDetail{}
				json.Unmarshal([]byte(bookDetailStr), &bookDetail)
				bookDetails = append(bookDetails, bookDetail)
			}
			//设置到消息类中
			msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_SUCCESS,
				Error: "",
				Data:  bookDetails}
			msg.Send(c)
		} else if key == KEY_HUXIU_IN_REDIS {
			huxiuNewsList := make([]huxiu.HuxiuNews, 0)
			for i, _ := range result {
				huxiuNewsStr := result[i]
				huxiu := huxiu.HuxiuNews{}
				json.Unmarshal([]byte(huxiuNewsStr), &huxiu)
				huxiuNewsList = append(huxiuNewsList, huxiu)
			}
			//设置到消息类中
			msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_SUCCESS,
				Error: "",
				Data:  huxiuNewsList}
			msg.Send(c)
		}

	}
}

func getHuxius(c *gin.Context) {
	key := c.Param("key")
	start := c.DefaultQuery("start", "0")
	end := c.Query("end")

	//获取结果
	//ZRANGE w3ckey 0 10 WITHSCORES
	result, err := redis.Strings(conn.Do("ZREVRANGE", key, start, end, "WITHSCORES"))

	if err != nil {
		msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_FAILED,
			Error: err.Error(),
			Data:  "",
		}
		msg.Send(c)
	} else {
		//反序列化到数组中
		if key == KEY_HUXIU_IN_REDIS || key == KEY_CHULE_IN_REDIS {
			huxiuNewsList := make([]huxiu.HuxiuNews, 0)
			for i, _ := range result {
				huxiuNewsStr := result[i]
				huxiu := huxiu.HuxiuNews{}
				json.Unmarshal([]byte(huxiuNewsStr), &huxiu)
				huxiuNewsList = append(huxiuNewsList, huxiu)
			}
			//设置到消息类中
			msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_SUCCESS,
				Error: "",
				Data:  huxiuNewsList}
			msg.Send(c)
		} else if key == KEY_HUXIU_DETAIL_IN_REDIS {
			detailList := make([]huxiu.HuxiuDetail, 0)
			for i, _ := range result {
				huxiuDetailStr := result[i]
				detail := huxiu.HuxiuDetail{}
				json.Unmarshal([]byte(huxiuDetailStr), &detail)
				detailList = append(detailList, detail)
			}
			//设置到消息类中
			msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_SUCCESS,
				Error: "",
				Data:  detailList}
			msg.Send(c)
		}

	}
}

func getDetail(c *gin.Context) {
	mapKey := c.Param("mapKey")
	key := c.Param("key")

	//获取结果
	result, err := redis.Strings(conn.Do("HMGET", mapKey, key))

	if err != nil {
		msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_FAILED,
			Error: err.Error(),
			Data:  "",
		}
		msg.Send(c)
	} else {
		//反序列化到数组中
		if mapKey == KEY_HUXIU_DETAIL_IN_REDIS || mapKey == KEY_CHULE_DETAIL_IN_REDIS {
			detailList := make([]huxiu.HuxiuDetail, 0)
			for i, _ := range result {
				huxiuDetailStr := result[i]
				detail := huxiu.HuxiuDetail{}
				json.Unmarshal([]byte(huxiuDetailStr), &detail)
				detailList = append(detailList, detail)
			}
			//设置到消息类中
			msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_SUCCESS,
				Error: "",
				Data:  detailList}
			msg.Send(c)
		}

	}
}
