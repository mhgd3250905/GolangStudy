package modle

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm_demo:"size:255;"`      //string默认长度为255，使用tag重新设置
	Num      int    `gorm_demo:"AUTO_INCREMENT"` //自增

	CreditCard CreditCard //One-to-One(拥有一个-CreditCard表的UserID作为外键)
	Emails     []Email    //One-to-Many(拥有多个-Email表的UserID作为外键)

	BillingAddress   Address //One-to-One（属于-本表的BillingAddressID作为外键）
	BillingAddressID int

	IgnoreMe  int         `gorm_demo："-"` //忽略这个字段
	Languages []Languages `gorm_demo:"mang2mang:user_languages;"`
}

/**
创建一个modle到表中
*/
func (this *User)CreateRecord(db *gorm.DB) (hasCreated bool,err error) {

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

func (this *User)Show(){
	fmt.Printf("Name: %v , Age: %v\n",this.Name,this.Age)
}