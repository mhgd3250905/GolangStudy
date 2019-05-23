package process

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"encoding/json"
	"fmt"
)

func OutputGroupMsg(msg *message.Message) (err error){
	//显示即可
	//1.反序列化msg.Data
	var smsMsg message.SmsMsg
	err=json.Unmarshal([]byte(msg.Data),&smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail,err= ",err)
		return
	}

	//显示信息
	info:=fmt.Sprintf("用户id:\t%d 对大家说:\t%s",smsMsg.User.UserId,smsMsg.Content)
	fmt.Println(info)
	fmt.Println()
	return
}
