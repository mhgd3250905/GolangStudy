package process

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"fmt"
)

//客户端要维护map
var onlineUsers = make(map[int]*message.User, 10)

//在客户端显示当前在线用户
func outputOnlineUser() {
	//遍历onlineUsers
	fmt.Println("当前在线用户列表：")
	for id,_:=range onlineUsers{
		fmt.Println("用户id:\t",id)
	}
}


//编写一个方法，处理返回的NotifuUserStatusMsg
func updateUserStatus(notifyUserStatusMsg *message.NotifyUserStatusMsg)  {
	//适当的优化，考虑是否已经存在
	user,ok:=onlineUsers[notifyUserStatusMsg.UserId]
	if !ok {
		user=&message.User{
			UserId:notifyUserStatusMsg.UserId,
		}
	}
	user.UserStatus=notifyUserStatusMsg.Status
	onlineUsers[notifyUserStatusMsg.UserId]=user
	outputOnlineUser()
}
