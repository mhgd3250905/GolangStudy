package model

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//我们在服务器启动后，就初始化一个userDao实例
//将其做成全局的变量，在需要和redis操作时，就直接使用可以
var (
	MyUserDao *UserDao
)

//定义一个UserDao 结构体
//完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}

//1.根据一个用户id 返回 一个User实例+err
func (this *UserDao) getUserByid(conn redis.Conn, id int) (user *message.User, err error) {
	//通过给定id去 redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误！
		if err == redis.ErrNil { //表示在 users 哈希中没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &message.User{}

	//需要把res反序列化为User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//完成登录的校验
//1.Login完成对用户的验证
//2.如果用户的id和pwd都正确，则返回一个user实例
//3.如果用户id或pwd,则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {

	//先从UserDao的连接池中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserByid(conn, userId)
	if err != nil {
		return
	}

	//这时候证明用户获取到了
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {

	//先从UserDao的连接池中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserByid(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	//这是说明id在redis中还没有，则可以完成注册
	data, err := json.Marshal(user) //序列化
	if err != nil {
		return
	}

	//入库
	_,err=conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil{
		fmt.Println("保存注册用户错误 err=",err)
		return
	}
	return
}
