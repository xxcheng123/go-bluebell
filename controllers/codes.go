package controllers

type ResponseCode uint16
type ResponseMessage string
type ResponseData any

type ResponseWrapper struct {
	Code    ResponseCode    `json:"code"`
	Message ResponseMessage `json:"message"`
	Data    ResponseData    `json:"data"`
}

func (respC ResponseCode) GetMsg() ResponseMessage {
	if msg, ok := MessageMap[respC]; ok {
		return msg
	}
	return MessageMap[CodeBusy]
}

const (
	CodeSuccess           ResponseCode = 0
	CodeFailedBindParam   ResponseCode = 1001
	CodeFailedVerifyParam ResponseCode = 1002

	CodeRegisterFailed    ResponseCode = 2001
	CodePasswordIncorrect ResponseCode = 2002
	CodeUserExist         ResponseCode = 2003

	CodeBusy ResponseCode = 9999
)

var MessageMap = map[ResponseCode]ResponseMessage{
	CodeSuccess:           "success",
	CodeFailedBindParam:   "参数解析绑定失败",
	CodeFailedVerifyParam: "参数验证失败",

	CodeRegisterFailed:    "注册失败",
	CodePasswordIncorrect: "密码错误",
	CodeUserExist:         "用户已存在",

	CodeBusy: "内部繁忙",
}
