package model

type ConnMsg struct {
	Users []User `json:"users"` //用户列表
	Msg   string `json:"msg"`   //通知消息
}

type User struct {
	Name string `json:"name"` //用户名
	Ip   string `json:"ip"`   //用户IP
}

type MsgType int

const (
	TYPE_DATA MsgType = 0
	TYPE_CONN MsgType = 1
	TYPE_USER MsgType = 2
)

type Data struct {
	Type    MsgType `json:"type"`     //数据类型 0->Data 1->ConnMsg 2->UserMsg
	DataMsg string  `json:"data_msg"` //路径数据
	ConnMsg ConnMsg `json:"conn_msg"` //用户信息
	UserMsg UserMsg `json:"user_msg"` //用户发送之消息
}

type UserMsg struct {
	Users User `json:"users"` //用户列表
	Msg   string `json:"msg"`   //通知消息
}
