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

	user_1:=modle.User{}
	db.First(&user_1)

	user_1.Name="shengkai_update_0"
	user_1.Age=100

	db.Save(&user_1)

	user_2:=modle.User{}

	db.Where("name = ?","shengkai_update_0").First(&user_2)

	user_2.Show()

}