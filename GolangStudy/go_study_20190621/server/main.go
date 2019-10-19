package main

import (
	"GolangStudy/GolangStudy/go_study_20190621/modle"
	"GolangStudy/GolangStudy/go_study_20190621/modle/comic"
	"GolangStudy/GolangStudy/go_study_20190621/modle/huxiu"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
)

const (
	KEY_HUXIU_IN_REDIS        = "huxiu"
	KEY_HUXIU_INFO_IN_REDIS   = "huxiu_info"
	KEY_CHULE_IN_REDIS        = "chule"
	KEY_CHULE_INFO_IN_REDIS   = "chule_info"
	KEY_IFANR_IN_REDIS        = "ifanr"
	KEY_IFANR_INFO_IN_REDIS   = "ifanr_info"
	KEY_HUXIU_DETAIL_IN_REDIS = "huxiu_detail"
	KEY_CHULE_DETAIL_IN_REDIS = "chule_detail"
	KEY_IFANR_DETAIL_IN_REDIS = "ifanr_detail"

	KEY_COMIC_BOOK_ID_IN_REDIS   = "COMIC_BOOK_ID"
	KEY_COMIC_BOOK_INFO_IN_REDIS = "COMIC_BOOK_INFO"
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
	r.GET("/spider/news/:key", getNews)
	r.GET("/spider/detail/:mapKey/:key", getDetail)
	r.GET("/spider/comic/book", getComicList)
	r.GET("/spider/comic/chapter", getChapterInfo)
	r.GET("/spider/comic/chapter/image", getChapterImage)
	r.Run(":8080")
}

func getNews(c *gin.Context) {
	key := c.Param("key")
	start := c.DefaultQuery("start", "0")
	end := c.Query("end")

	//获取结果
	//ZRANGE w3ckey 0 10 WITHSCORES
	result, err := redis.Strings(conn.Do("ZREVRANGE", key, start, end))

	//取出来的是一串id,要分别获取保存的信息
	infoKey := KEY_HUXIU_INFO_IN_REDIS
	if key == KEY_HUXIU_IN_REDIS {

		infoKey = KEY_HUXIU_INFO_IN_REDIS

	} else if key == KEY_CHULE_IN_REDIS {

		infoKey = KEY_CHULE_INFO_IN_REDIS

	} else if key == KEY_IFANR_IN_REDIS {

		infoKey = KEY_IFANR_INFO_IN_REDIS

	}

	var id string
	var newsInfo string
	newsInfoArr := make([]string, 0)
	for i, _ := range result {
		id = result[i]
		newsInfo, err = redis.String(conn.Do("HGET", infoKey, id))
		if err != nil {
			continue
		}
		newsInfoArr = append(newsInfoArr, newsInfo)
	}

	if err != nil {
		msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_FAILED,
			Error: err.Error(),
			Data:  "",
		}
		msg.Send(c)
	} else {
		//反序列化到数组中
		if key == KEY_HUXIU_IN_REDIS || key == KEY_CHULE_IN_REDIS || key == KEY_IFANR_IN_REDIS {
			huxiuNewsList := make([]huxiu.HuxiuNews, 0)
			for i, _ := range newsInfoArr {
				huxiuNewsStr := newsInfoArr[i]
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
			for i, _ := range newsInfoArr {
				huxiuDetailStr := newsInfoArr[i]
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
		if mapKey == KEY_HUXIU_DETAIL_IN_REDIS || mapKey == KEY_CHULE_DETAIL_IN_REDIS || mapKey == KEY_IFANR_DETAIL_IN_REDIS {
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

func getComicList(c *gin.Context) {
	start := c.DefaultQuery("start", "0")
	end := c.Query("end")

	//获取结果
	//ZRANGE w3ckey 0 10 WITHSCORES
	result, err := redis.Strings(conn.Do("ZREVRANGE", KEY_COMIC_BOOK_ID_IN_REDIS, start, end))

	//取出来的是一串id,要分别获取保存的信息
	infoKey := KEY_COMIC_BOOK_INFO_IN_REDIS

	var id string
	var newsInfo string
	comicInfoArr := make([]string, 0)
	for i, _ := range result {
		id = result[i]
		newsInfo, err = redis.String(conn.Do("HGET", infoKey, id))
		if err != nil {
			continue
		}
		comicInfoArr = append(comicInfoArr, newsInfo)
	}

	if err != nil {
		msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_FAILED,
			Error: err.Error(),
			Data:  "",
		}
		msg.Send(c)
	} else {
		//反序列化到数组中

		comicList := make([]comic.ComicBook, 0)
		for i, _ := range comicInfoArr {
			comicStr := comicInfoArr[i]
			comic := comic.ComicBook{}
			json.Unmarshal([]byte(comicStr), &comic)
			comicList = append(comicList, comic)
		}
		//设置到消息类中
		msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_SUCCESS,
			Error: "",
			Data:  comicList}
		msg.Send(c)
	}

}

func getChapterInfo(c *gin.Context) {

	id := c.Query("id")

	//获取结果
	//ZRANGE w3ckey 0 10 WITHSCORES
	result, err := redis.Strings(conn.Do("LRANGE", id, "0", "-1"))

	////取出来的是一串id,要分别获取保存的信息
	//infoKey := KEY_COMIC_BOOK_INFO_IN_REDIS
	//
	//var newsInfo string
	//comicInfoArr := make([]string, 0)
	//for i, _ := range result {
	//	id = result[i]
	//	newsInfo, err = redis.String(conn.Do("HGET", infoKey, id))
	//	if err != nil {
	//		continue
	//	}
	//	comicInfoArr = append(comicInfoArr, newsInfo)
	//}

	if err != nil {
		msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_FAILED,
			Error: err.Error(),
			Data:  "",
		}
		msg.Send(c)
	} else {
		//反序列化到数组中

		//comicList := make([]comic.ComicBook, 0)
		//for i, _ := range comicInfoArr {
		//	comicStr := comicInfoArr[i]
		//	comic := comic.ComicBook{}
		//	json.Unmarshal([]byte(comicStr), &comic)
		//	comicList = append(comicList, comic)
		//}
		//设置到消息类中
		msg := modle.Message{ErrCode: modle.MESSAGE_CODE_QUERY_SUCCESS,
			Error: "",
			Data:  result}
		msg.Send(c)
	}

}

func getChapterImage(c *gin.Context){
	path := c.Query("path")

	c.File(string(path))
}
