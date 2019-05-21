package message

//定义一个用户结构体
type User struct {
	//为了能够序列化反序列化成功
	//必须加入对应的tag
	UserId     int    `json:userId`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`//用户状态..
}
