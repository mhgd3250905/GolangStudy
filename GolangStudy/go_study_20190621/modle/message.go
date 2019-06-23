package modle

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

const (
	MESSAGE_CODE_QUERY_SUCCESS = 0
	MESSAGE_CODE_QUERY_FAILED  = 101
	MESSAGE_CODE_REDIS_FAILED  = 102
)

type Message struct {
	ErrCode int         `json:"err_code"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func (this *Message) Send(c *gin.Context){
	msgStr, err := json.Marshal(&this)
	if err != nil {

	} else {
		c.String(200, string(msgStr))
	}
}