package process2

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

//写方法转发
func (this *SmsProcess) SendGroupMsg(msg *message.Message) (err error) {

	//遍历服务器端的onlineUsers
	//将消息转发出去
	//取出msg中的内容
	var smsMsg message.SmsMsg
	err = json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail,err=", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal fail,err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//这里，还需要过滤到自己，即不要在发给自己
		if id == smsMsg.User.UserId {
			continue
		}
		err = this.SendMsgToEachUser(data, up.Conn)
		if err != nil {
			println("SendMsgToEachUser fail,err=", err)
			continue
		}
	}
	return
}

//转发信息给每一个用户
func (this *SmsProcess) SendMsgToEachUser(data []byte, conn net.Conn) (err error) {

	//创建一个Transfer 发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
	return
}
