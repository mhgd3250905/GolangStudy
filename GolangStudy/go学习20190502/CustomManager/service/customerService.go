package service

import "GolangStudy/GolangStudy/go学习20190502/CustomManager/modle"

//该结构体完成对Customer的操作，包括
//增删改查
type CustomerService struct {
	customers []modle.Customer
	//声明一个字段,表示当前切片含有多少个客户
	//该字段后面还可以作为新客户的id
	customerNum int
}

//获取到CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := modle.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []modle.Customer {
	return this.customers
}

//添加客户
func (this *CustomerService) Add(customer modle.Customer) bool {
	//确定一个Id的分配规则，就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

//根据Id删除客户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		//没有这个客户
		return false
	}
	//删除客户
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

//根据Id更新用户数据
func (this *CustomerService) Update(id int,customer modle.Customer) bool{
	index:=this.FindById(id)
	if index == -1 {
		//没有这个客户
		return false
	}
	this.customers[index]=customer
	return true
}

//根据Id查找客户在切片中对应下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {
	//遍历
	index := -1
	//遍历this.customers 切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			//找到了
			index = i
		}
	}
	return index
}
