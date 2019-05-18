package process2

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/server/model"
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	//
}

//编写一个函数serverProcessLogin专门，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	//核心代码
	//1.先从msg中取出msg.Data,并反序列化为LoginMeg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1.先声音一个 reMsg
	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType

	//2.再声明一个LoginResMessage
	var loginResMsg message.LoginResMsg

	//我们需要先到redis数据库完成验证
	//1.使用model.MyUserDao 到redis验证
	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMsg.Code = 500 //500 状态码表示该用户不存在
			loginResMsg.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMsg.Code = 403 //500 状态码表示该用户不存在
			loginResMsg.Error = err.Error()
		}else {
			loginResMsg.Code = 505 //500 状态码表示该用户不存在
			loginResMsg.Error = "服务器内部错误..."
		}
		//这里我们先测试成功，然后我们可以根据错误返回具体的信息
	} else {
		loginResMsg.Code = 200
		fmt.Println(user," 登录成功")
	}

	//如果用户id为100，密码为123456，认为合法，否则不合法
	//if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
	//	//合法
	//	loginResMsg.Code = 200
	//} else {
	//	//不合法
	//	loginResMsg.Code = 500 //500 状态码表示该用户不存在
	//	loginResMsg.Error = "该用户不存在，请注册再使用..."
	//}

	//3.将loginResMsg序列化
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//4.将data赋值给resMsg
	resMsg.Data = string(data)

	//5.对resMsg序列化并准备发送
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//6.发送data,将其封装到writePkg函数中
	//因为使用了分层模式（MVC）先创建一个Transfer,然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
