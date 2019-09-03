package redis_utils

import "github.com/gomodule/redigo/redis"

func Push2RedisList(c redis.Conn, key string, content string) (err error) {
	_, err = c.Do("RPUSH", key, content)
	return
}

func Push2RedisSortedSet(c redis.Conn, newsId string, idKey string,infoKey string, score string, content string) (err error) {
	err = SaveNewsId(c, idKey, score, newsId)
	if err != nil {
		return
	}
	err = SaveHashMap(c, infoKey, newsId, content)
	return
}

func SaveHashMap(c redis.Conn, key string, mapKey string, content string) (err error) {
	_, err = c.Do("HMSET", key, mapKey, content)
	return
}

func SaveNewsId(c redis.Conn, key string, score string, newsId string) (err error) {
	_, err = c.Do("ZADD", key, score, newsId)
	return
}
