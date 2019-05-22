package model

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"net"
)

//因为在客户端我们很多地方会用到当前用户，所以我们将其作为全局
type CurUser struct {
	Conn net.Conn
	User message.User
}
