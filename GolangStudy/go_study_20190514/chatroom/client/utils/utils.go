package utils

import (
	"GolangStudy/GolangStudy/go_study_20190514/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	//它应该有哪些字段
	Conn net.Conn
	Buf [8096]byte //这是传输时使用的缓冲

}

func (this *Transfer)ReadPkg() (msg message.Message, err error) {

	fmt.Println("等待读取客户端发送的数据...")
	n, err := this.Conn.Read(this.Buf[:4])
	//conn.Read 在从你没有被关闭的情况下才会阻塞
	//如果客户端关闭了conn连接，就不会堵塞了
	if n != 4 || err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}

	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen读取消息内容
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}

	//把获取数据反序列化为->message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	return
}

func (this *Transfer)WritePkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(bytes) err=", err)
		return err
	}

	//发送消息本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.write(data) fail,err= ", err)
		return
	}

	return
}