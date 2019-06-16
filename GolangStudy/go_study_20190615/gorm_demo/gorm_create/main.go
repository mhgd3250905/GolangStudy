package main

import (
	"GolangStudy/GolangStudy/go_study_20190615/gorm_demo/modle"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

var db *gorm.DB

func main() {
	db, err := gorm.Open("sqlite3", "C:/Users/admin/go/src/GolangStudy/GolangStudy/go_study_20190615/sqlite3/gorm_demo.db")

	if err != nil {
		fmt.Println("sqlite3 db open fail err= ", err)
		return
	}
	defer db.Close()
	fmt.Println("db open success -> ",db)

	for i := 0; i < 20; i++ {
		//创建一个表并写入数据
		user := modle.User{Name: fmt.Sprintf("shengkai No.%d",i), Age: i, Birthday: time.Now()}
		user.CreateRecord(db)
	}
}
