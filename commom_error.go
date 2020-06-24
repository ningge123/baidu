package baidu

type CommonError struct {
	ErrCode int `json:"error_code"`  // 错误码
	ErrMsg string `json:"error_msg"` // 错误描述
}