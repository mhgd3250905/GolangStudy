package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var server Server = Server{"Painting", make(map[string]*Client, 0), make(chan error)}

type Server struct {
	Pattern string
	Clients map[string]*Client
	ErrCh   chan error
}

type Client struct {
	Id     string
	Msg    chan []byte
	Ws     *websocket.Conn
	Server *Server
	DoneCh chan bool
}

func NewClient(ws *websocket.Conn, server *Server) *Client {
	if ws == nil {
		panic("ws cannot be nil!")
	}
	if server == nil {
		panic("server cannot be nil!")
	}
	ID := ws.RemoteAddr().String()
	doneCh := make(chan bool)
	msg := make(chan []byte)
	return &Client{ID, msg, ws, server, doneCh}
}

//webSocket请求ping返回pong
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	ID := ws.RemoteAddr().String()

	var client *Client
	if server.Clients[ID] == nil {
		client = NewClient(ws, &server)
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

		message := makeNoticeMessage(ID, "连接...")
		sendMessage(message, ID)
	}

	for {
		//读取ws中的数据
		_, message, err := ws.ReadMessage()
		if err != nil {
			delete(server.Clients, ID)
			message = makeNoticeMessage(ID, "退出了连接")
			sendMessage(message, ID)
			break
		}

		//fmt.Printf("mt: %d\n",mt)

		fmt.Printf("%s ,Message: %s\n", ws.RemoteAddr().String(), string(message))
		if string(message) == "ping" {
			message = []byte("pong")
		}

		sendMessage(message, ID)
	}
}

func makeNoticeMessage(userId string, message string) []byte {
	return []byte(fmt.Sprintf("Notice:%s %s", userId, message))
}

func sendMessage(message []byte, exceptId string) {
	for k, v := range server.Clients {
		if k == exceptId {
			continue
		}
		//发送消息到通道
		v.Msg <- message
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run(":80")
}
