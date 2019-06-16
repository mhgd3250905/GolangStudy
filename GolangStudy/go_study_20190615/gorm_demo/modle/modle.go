package modle

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)



type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

type Email struct {
	ID         int
	UserID     int    `gorm_demo:"index"`                           //外键（属于）,tag `index`是为该列创建索引
	Email      string `gorm_demo:"type:varchar(100);unique_index;"` //`type`设置sql类型,`unique_index`为该列设置唯一索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm_demo:"not null;unique"` //设置字段为非空并唯一
	Address2 string         `gorm_demo:"type:varchar(100);unique"`
	Post     sql.NullString `gorm_demo:"not null"`
}

type Languages struct {
	ID   int
	Name string `gorm_demo:"index:idx_name_code"` //创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm_demo:"index:idx_name_code"` //`unique_index` also works
}



