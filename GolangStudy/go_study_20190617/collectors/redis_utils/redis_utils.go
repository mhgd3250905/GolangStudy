package redis_utils

import "github.com/gomodule/redigo/redis"

func Push2RedisList(c redis.Conn, key string, content string) (err error) {
	_, err = c.Do("RPUSH", key, content)
	return
}

func Push2RedisSortedSet(c redis.Conn, key string, score string, content string) (err error) {
	_, err = c.Do("ZADD", key, score, content)
	return
}

func SaveHashMap(c redis.Conn, key string, mapKey string, content string) (err error) {
	_, err = c.Do("HMSET", key, mapKey, content)
	return
}
