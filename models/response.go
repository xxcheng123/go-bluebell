package models

type ResponseCode uint16
type ResponseMessage string

type ResponseCommon struct {
	Code    ResponseCode    `json:"code"`
	Message ResponseMessage `json:"message"`
}

const (
	ResponseCodeSuccess           = 0
	ResponseCodeFailedBindParam   = 1001
	ResponseCodeFailedVerifyParam = 1002
)

const (
	ResponseMessageSuccess           = "success"
	ResponseMessageFailedBindParam   = "参数解析绑定失败"
	ResponseMessageFailedVerifyParam = "参数验证失败"
)
