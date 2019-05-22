package message

//确定一些消息类型常量
const (
	LoginMsgType            = "LoginMsg"
	LoginResMsgType         = "LoginResMsg"
	RegisterMsgType         = "RegisterMsg"
	RegisterResMsgType      = "RegisterResMsgType"
	NotifyUserStatusMsgType = "NotifyUserStatusMsgType"
	SmsMsgType              = "SmsMsgType"
)

//定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的内容
}

//定义登陆消息

type LoginMsg struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

//定义登陆反馈消息
type LoginResMsg struct {
	Code    int    `json:"code"`    //返回状态吗 500 表示该用户未注册 200表示注册成功
	UsersId []int  `json:"usersId"` //增加字段,保存用户id的切片
	Error   string `json:"error"`   //返回错误信息
}

//定义注册消息
type RegisterMsg struct {
	User User //类型就是User结构体
}

type RegisterResMsg struct {
	Code  int    `json:"code"`  //返回状态码 400 表示用户已存在，200表示注册成功
	Error string `json:"error"` //返回错误信息
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMsg struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户状态
}

//正价一个SmsMes ,发送的
type SmsMsg struct {
	Content string `json:"content"` //内容
	User    User   `json:"user"`
}
