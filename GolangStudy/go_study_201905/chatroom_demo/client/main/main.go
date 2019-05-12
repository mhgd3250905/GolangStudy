package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	Start(":9999")
}

func Start(tcpAddrStr string) {
	tcpAddr,err:=net.ResolveTCPAddr("tcp4",tcpAddrStr)
	if err != nil {
		log.Printf("Resolve tcp addr failed:%v\n",err)
		return
	}

	//向服务器拨号
	conn,err:=net.DialTCP("tcp",nil,tcpAddr)
	if err != nil {
		log.Printf("Dial to server failed:%v\n",err)
		return
	}

	//向服务器发送消息
	go SendMsg(conn)

	//接收来自服务器端的广播消息
	buf:=make([]byte,1024)
	for{
		length,err:=conn.Read(buf)
		if err!=nil{
			log.Printf("recv server msg failed:%v\n",err)
			conn.Close()
			os.Exit(0)
			break
		}

		fmt.Println(string(buf[:length]))
	}
}

//向服务器发送消息
func SendMsg(conn *net.TCPConn) {
	scanner:=bufio.NewScanner(os.Stdin)
	for{
		var input string

		//接收输入的消息，放入到input变量中
		scanner.Scan()
		input=scanner.Text()

		if input =="/q"||input=="/quit"{
			fmt.Println("Byebye...")
			conn.Close()
			os.Exit(0)
		}

		//log.Print(input)

		//只处理有内容的消息
		if len(input) > 0 {
			_,err:=conn.Write([]byte(input))
			if err != nil {
				log.Printf("conn write fialed:%v\n",err)
				conn.Close()
				break
			}
		}
	}
}
