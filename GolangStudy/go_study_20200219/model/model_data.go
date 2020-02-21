package model

type UsersMsg struct {
	Users []User `json:"users"` //用户列表
	Msg   string `json:"msg"`   //通知消息
}

type User struct {
	Name string `json:"name"` //用户名
	Ip   string `json:"ip"`   //用户IP
}
