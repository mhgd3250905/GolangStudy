package process

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/client/utils"
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {
}

//发送群聊的消息
func (this *SmsProcess) sendGroup(content string) (err error) {

	//1.创建一个Msg
	var msg message.Message
	msg.Type = message.SmsMsgType

	//2.创建一个SmsMsg
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.User.UserId = CurUser.User.UserId
	smsMsg.User.UserStatus = CurUser.User.UserStatus

	//3.序列化smsMsg
	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	msg.Data = string(data)

	//4.再次序列化对msg
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//5.将sms发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMsg fail err=", err)
		return
	}
	return
}
