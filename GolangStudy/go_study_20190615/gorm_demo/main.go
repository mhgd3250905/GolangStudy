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
	//fmt.Println("db open success -> ",db)

	//for i := 0; i < 20; i++ {
	//	//创建一个表并写入数据
	//	user := modle.User{Name: fmt.Sprintf("shengkai No.%d",i), Age: i, Birthday: time.Now()}
	//	user.CreateRecord(db)
	//}

	//查询第一条
	firstUser := modle.User{}
	db.First(&firstUser)

	fmt.Println("user query first record: ", firstUser)

	//查询最后一条
	lastUser := modle.User{}
	db.Last(&lastUser)

	fmt.Println("user query last record: ", lastUser)

	//查询所有记录
	users := make([]modle.User, 20)
	db.Find(&users)

	fmt.Println("user query all record: ", users)

	//根据主键获取记录
	fifthUser := modle.User{}
	db.First(&fifthUser, 5)

	fmt.Println("user query first id =10: ", fifthUser)
}
