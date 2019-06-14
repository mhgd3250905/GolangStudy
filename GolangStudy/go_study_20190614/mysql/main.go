package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId int `db:user_id`
	Username string `db:username`
	Sex string `db:sex"`
	Email string `db:email`
}

type Place struct {
	Country string `db:country`
	City string `db:city`
	TelCode int `db:telcode`
}

//线程安全的内置了连接池
var Db *sqlx.DB


func init(){
	database,err:=sqlx.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed err= ",err)
		return
	}
	Db=database
}

func main() {
	r,err:=Db.Exec("insert into person(username,sex,email)values(?,?,?)","stu001","man","sty001@qq.com")
	if err != nil {
		fmt.Println("exec failed err= ",err)
		return
	}

	id,err:=r.LastInsertId()
	if err!=nil{
		fmt.Println("exec failed err= ",err)
		return
	}

	fmt.Println("insert succ ",id)
}