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
	//增加一个字段，表示该Conn是哪个用户的
	UserId int
}

//编写通知所有的在线用户的方法
//userId要通知其他的在线用户 我上线了
func (this *UserProcess)NotifyOthersOnlineUser(userId int)  {
	//遍历onlineUsers,然后一个一个发送
	for id,up:=range userMgr.onlineUsers{
		//过滤掉自己
		if id == userId {
			continue
		}
		//开始通知 单独开辟一个方法
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess)NotifyMeOnline(userId int){

	//组装我们的NotifyUserStatusMes
	var msg message.Message
	msg.Type=message.NotifyUserStatusMsgType

	var notifyUserStatusMsg message.NotifyUserStatusMsg
	notifyUserStatusMsg.UserId=userId
	notifyUserStatusMsg.Status=message.UserOnline

	//将NotifuUserStatusMsg序列化
	data,err:=json.Marshal(notifyUserStatusMsg)
	if err != nil {
		fmt.Println("json.Marshal fail err=",err)
	}

	msg.Data=string(data)

	//将msg序列化
	data,err=json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal fail err=",err)
	}

	//发送,创建一个Transfer实例发送
	tf:=&utils.Transfer{
		Conn:this.Conn,
	}
	err=tf.WritePkg(data)
	if err != nil {
		fmt.Println("notify me online err=",err)
		return
	}
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
		} else {
			loginResMsg.Code = 505 //500 状态码表示该用户不存在
			loginResMsg.Error = "服务器内部错误..."
		}
		//这里我们先测试成功，然后我们可以根据错误返回具体的信息
	} else {
		loginResMsg.Code = 200
		//因为用户已经登录成功,我们就把登录成功的用户放入到UserMgr中
		//将登录成功的UserId赋值给this
		this.UserId = loginMsg.UserId
		userMgr.AddOnlineUser(this)

		//通知其他在线的用户，我上线了
		this.NotifyOthersOnlineUser(loginMsg.UserId)

		//将当前在线用户的id放入到loginResMsg的UserIds
		//遍历 userMgr.OnlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMsg.UsersId = append(loginResMsg.UsersId, id)
		}
		fmt.Println(user, " 登录成功")
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

func (this *UserProcess) ServerProcessRegister(msg *message.Message) (err error) {
	//1.先从msg中取出msg.Data,并反序列化为RegisterMsg
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1.先声音一个 reMsg
	var resMsg message.Message
	resMsg.Type = message.RegisterResMsgType

	//2.再声明一个LoginResMessage
	var registerResMsg message.RegisterResMsg

	//我们需要先到redis数据库完成注册
	//1.使用model.MyUserDao 到redis验证
	err = model.MyUserDao.Register(&registerMsg.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMsg.Code = 505
			registerResMsg.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMsg.Code = 506
			registerResMsg.Error = "注册发生未知错误..."
		}
	} else {
		registerResMsg.Code = 200
	}

	//3.将registerResMsg序列化
	data, err := json.Marshal(registerResMsg)
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
