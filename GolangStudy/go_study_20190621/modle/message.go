package modle

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
