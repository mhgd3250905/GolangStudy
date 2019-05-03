package modle

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//编写一个工厂模式返回一个Customer
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//第二种创建的方法
func NewCustomerWithoutId(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//返回用户的信息
func (this *Customer)GetInfo() string {
	info:=fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
		this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Email)
	return info
}