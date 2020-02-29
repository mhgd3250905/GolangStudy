package main

import (
	"GolangStudy/GolangStudy/go_study_20200219/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

const USER_MSG_FLAG = "[-MSG-]"

var server = Server{"Painting", make(map[string]*Client, 0), make(chan error)}

type Server struct {
	Pattern string
	Clients map[string]*Client
	ErrCh   chan error
}

type Client struct {
	Id     string
	RoomId string
	Msg    chan []byte
	Ws     *websocket.Conn
	Server *Server
	DoneCh chan bool
}

func main() {
	r := gin.Default()
	r.GET("/chatroom/:roomId/listen", listen)
	r.Run(":80")
}

func NewClient(ws *websocket.Conn, server *Server, roomId string) *Client {
	if ws == nil {
		panic("ws cannot be nil!")
	}
	if server == nil {
		panic("server cannot be nil!")
	}
	ID := ws.RemoteAddr().String()
	doneCh := make(chan bool)
	msg := make(chan []byte)
	return &Client{ID, roomId, msg, ws, server, doneCh}
}

func listen(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	//获取房间Id
	roomId := c.Param("roomId")

	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	ID := ws.RemoteAddr().String()

	var client *Client
	if server.Clients[ID] == nil {
		client = NewClient(ws, &server, roomId)
		server.Clients[ID] = client

		go func(client *Client) {
			for {
				var message []byte
				message = <-client.Msg
				fmt.Printf("%s 接收数据：%s\n", client.Id, string(message))
				err = ws.WriteMessage(1, message)
				if err != nil {
					close(client.Msg)
					fmt.Printf("%s 协程关闭！", client.Id)
					break
				}
			}
		}(client)

		users := collectUsersInfo()
		connMsg := model.ConnMsg{Users: users, Msg: fmt.Sprintf("%s %s", ID, "连接成功。")}
		data := model.Data{Type: model.TYPE_CONN, ConnMsg: connMsg}

		b, err := json.Marshal(data)
		if err != nil {
			fmt.Println("json 解析错误")
		}
		sendMessage(b, ID, roomId)
	}

	for {
		//读取ws中的数据
		_, message, err := ws.ReadMessage()
		if err != nil {
			delete(server.Clients, ID)
			//message = makeNoticeMessage(ID, "退出了连接")
			users := collectUsersInfo()
			connMsg := model.ConnMsg{Users: users, Msg: fmt.Sprintf("%s %s", ID, "退出连接。")}
			data := model.Data{Type: model.TYPE_CONN, ConnMsg: connMsg}
			b, err := json.Marshal(data)
			if err != nil {
				fmt.Println("json 解析错误")
			}
			sendMessage(b, ID, roomId)
			break
		}

		//fmt.Printf("mt: %d\n",mt)

		//fmt.Printf("%s ,Message: %s\n", ws.RemoteAddr().String(), string(message))
		if string(message) == "ping" {
			message = []byte("pong")
		}

		var b []byte
		fmt.Printf("%s ,Message: %s\n", ws.RemoteAddr().String(), string(message)[:8])
		if string(message)[:7] == USER_MSG_FLAG {

			//转发用户消息
			splits := strings.Split(ws.RemoteAddr().String(), ":")
			user := model.User{Name: splits[len(splits)-1], Ip: ws.RemoteAddr().String()}
			userMsg := model.UserMsg{Users: user, Msg: string(message)[7:]}
			data := model.Data{Type: model.TYPE_USER, UserMsg: userMsg}
			b, err = json.Marshal(data)

			if err != nil {
				println("转换json错误！")
				break
			}

			sendMessage(b, "", roomId)
		} else {

			//转发路径消息
			data := model.Data{Type: model.TYPE_DATA, DataMsg: string(message)}
			b, err = json.Marshal(data)

			if err != nil {
				println("转换json错误！")
				break
			}

			sendMessage(b, ID, roomId)
		}
	}
}

//收集用户信息
func collectUsersInfo() []model.User {
	users := make([]model.User, 0)
	for k, _ := range server.Clients {
		splits := strings.Split(k, ":")
		users = append(users, model.User{Name: splits[len(splits)-1], Ip: k})
	}
	return users
}

func sendMessage(message []byte, exceptId string, roomId string) {
	for k, v := range server.Clients {
		if k == exceptId || v.RoomId != roomId {
			continue
		}
		//发送消息到通道
		v.Msg <- message
	}
}
