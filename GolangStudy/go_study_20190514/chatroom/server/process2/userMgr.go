package process2

import "fmt"

//因为UserMgr实例在服务器端有且只有一个
//因为在很多地方都会使用到，因此，我们将其定义为全局变量

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成哪个队onlineUser的添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DeleteOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//查询所有在线的用户
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

//查询指定UserId的用户
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess,err error) {
	//如何从map中取出一个值，待检测方式
	up,ok:=this.onlineUsers[userId]
	if !ok {
		//说明该用户是不在map中的
		//说明你要查找的这个用户当前不在线
		err=fmt.Errorf("用户%d不存在\n",userId)
		return
	}
	return
}