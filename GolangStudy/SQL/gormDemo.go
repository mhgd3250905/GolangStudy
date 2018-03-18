package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Product struct {
	ID uint `gorm:"primary_key"`
	Code string
	Price uint
}

type Student struct {
	Id uint `gorm:"primary_key"`
	Name string
	Age uint
	Address string
}

func main() {
	db, err := gorm.Open("mysql", "root:sk3250905@tcp(localhost:3306)/person")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// 自动迁移表，生成的表名为 products
	db.AutoMigrate(&Student{})

	// Create
	db.Create(&Student{Name: "jack", Age: 18,Address:"bj"})

	// Read
	var student Student
	db.First(&student, 1)                   // find product with id 1
	fmt.Println(student)
	db.First(&student, "name = ?", "jack") // find product with code l1212
	fmt.Println(student)


}