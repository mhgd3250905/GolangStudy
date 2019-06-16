package modle

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"` //设置默认值
	Age  int64
}

/**
创建一个modle到表中
*/
func (this *Animal)CreateRecord(db *gorm.DB) (hasCreated bool,err error) {

	if !db.HasTable(&this) {
		if err := db.CreateTable(&this).Error; err != nil {
			panic(err)
		}
	}

	isNewRecord := db.NewRecord(this) //主键为空返回`true`
	if !isNewRecord {
		return false,fmt.Errorf("this modle %v has created!",this)
	}

	db.Create(&this)

	isNewRecord = db.NewRecord(this) //主键为空返回`true`
	if !isNewRecord {
		return true,nil
	}else {
		return false,fmt.Errorf("this modle %v creare fail!",this)
	}
}