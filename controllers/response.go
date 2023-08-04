package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(ctx *gin.Context, code ResponseCode) {
	ctx.JSON(http.StatusOK, &ResponseWrapper{
		Code:    code,
		Message: code.GetMsg(),
		Data:    nil,
	})
}

func ResponseErrorWithMsg(ctx *gin.Context, code ResponseCode, msg string) {
	ctx.JSON(http.StatusOK, &ResponseWrapper{
		Code:    code,
		Message: ResponseMessage(fmt.Sprintf("%s,%s", code.GetMsg(), msg)),
		Data:    nil,
	})
}

func ResponseSuccess(ctx *gin.Context, data ResponseData) {
	ctx.JSON(http.StatusOK, &ResponseWrapper{
		Code:    CodeSuccess,
		Message: CodeSuccess.GetMsg(),
		Data:    data,
	})
}
