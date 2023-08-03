package controllers

import (
	"fmt"
	"go-generator/logic"
	"go-generator/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理用户注册
func SignUpHandler(ctx *gin.Context) {
	p := new(models.ParamSignUp)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with FailedBindParam", zap.Error(err))
		msg := models.ResponseMessageFailedBindParam
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, validationError := range errs {
				msg = fmt.Sprintf("%s,%s", msg, validationError.Translate(trans))
			}
		}
		ctx.JSON(200, gin.H{
			"code":    models.ResponseCodeFailedBindParam,
			"message": msg,
		})
		return
	}
	/*
		if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 {
			zap.L().Error("SignUp with FailedVerifyParam")
			ctx.JSON(200, gin.H{
				"code":    models.ResponseCodeFailedVerifyParam,
				"message": models.ResponseMessageFailedVerifyParam,
			})
			return
		}
	*/
	if err := logic.SignUp(p); err != nil {
		zap.L().Info("SignUp with Failed", zap.Error(err))
		ctx.JSON(200, gin.H{
			"code":    models.ResponseCodeRegisterFailed,
			"message": models.ResponseMessageRegisterFailed,
		})
		return
	}
	zap.L().Info("SignUp with Success")
	ctx.JSON(200, gin.H{
		"code":    models.ResponseCodeSuccess,
		"message": models.ResponseMessageSuccess,
	})
}
