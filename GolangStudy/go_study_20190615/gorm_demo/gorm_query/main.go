package main

import (
	"GolangStudy/GolangStudy/go_study_20190615/gorm_demo/modle"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func main() {
	db, err := gorm.Open("sqlite3", "C:/Users/admin/go/src/GolangStudy/GolangStudy/go_study_20190615/sqlite3/gorm_demo.db")

	if err != nil {
		fmt.Println("sqlite3 db open fail err= ", err)
		return
	}
	defer db.Close()
	fmt.Println("db open success -> ", db)

	//查询第一条
	firstUser := modle.User{}
	db.First(&firstUser)

	//fmt.Println("user query first record: ", firstUser)

	//查询最后一条
	lastUser := modle.User{}
	db.Last(&lastUser)

	//fmt.Println("user query last record: ", lastUser)

	//查询所有记录
	users := make([]modle.User, 20)
	db.Find(&users)

	//fmt.Println("user query all record: ", users)

	//根据主键获取记录
	fifthUser := modle.User{}
	db.First(&fifthUser, 5)

	//fmt.Println("user query first id =10: ", fifthUser)

	//where 查询
	whereUser_1 := modle.User{}
	db.Where("name = ?", "shengkai No.1").First(&whereUser_1)
	//whereUser_1.Show()

	whereUsers := make([]modle.User, 20)
	db.Where("name = ?", "shengkai No.2").Find(&whereUsers)
	//for _,user:=range whereUsers{
	//	user.Show()
	//}

	whereUsers_2 := make([]modle.User, 20)
	db.Where("name in (?)", []string{"shengkai No.1", "shengkai No.2"}).Find(&whereUsers_2)
	//for _,user:=range whereUsers_2{
	//	user.Show()
	//}

	whereUsers_3 := make([]modle.User, 20)
	db.Where("name LIKE ?", "%shengkai%").Find(&whereUsers_3)
	//for _,user:=range whereUsers_3{
	//	user.Show()
	//}

	whereUsers_4 := make([]modle.User, 20)
	db.Where("name LIKE ? And age >=?", "%shengkai%", 15).Find(&whereUsers_4)
	//for _, user := range whereUsers_4 {
	//	user.Show()
	//}

	whereUsers_5 := make([]modle.User, 20)
	db.Where("ID BETWEEN ? And ?", 5,10).Find(&whereUsers_5)
	//for _, user := range whereUsers_5 {
	//	user.Show()
	//}

	//struct 条件查询
	whereUser_2:=modle.User{}
	db.Where(&modle.User{Name:"shengkai No.3"}).First(&whereUser_2)
	//whereUser_2.Show()

	//map 条件查询
	whereUsers_6 := make([]modle.User, 20)
	db.Where(map[string]interface{}{"name":"shengkai No.4"}).Find(&whereUsers_6)
	//for _, user := range whereUsers_6 {
	//	user.Show()
	//}

	//主键的Slice 条件查询
	whereUsers_7 := make([]modle.User, 20)
	db.Where([]int64{6,7,8}).Find(&whereUsers_7)
	//for _, user := range whereUsers_7 {
	//	user.Show()
	//}




}
