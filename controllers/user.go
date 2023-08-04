package controllers

import (
	"errors"
	"fmt"
	"go-generator/dao/mysql"
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
		//msg := models.ResponseMessageFailedBindParam
		msg := ""
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, validationError := range errs {
				msg = fmt.Sprintf("%s,%s", msg, validationError.Translate(trans))
			}
		}
		ResponseErrorWithMsg(ctx, CodeFailedBindParam, msg)
		//ctx.JSON(200, gin.H{
		//	"code":    models.ResponseCodeFailedBindParam,
		//	"message": msg,
		//})
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
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(ctx, CodeUserExist)
		} else {
			ResponseError(ctx, CodeRegisterFailed)
		}
		//ctx.JSON(200, gin.H{
		//	"code":    models.ResponseCodeRegisterFailed,
		//	"message": models.ResponseMessageRegisterFailed,
		//})
		return
	}
	zap.L().Info("SignUp with Success")
	ResponseSuccess(ctx, nil)
	//ctx.JSON(200, gin.H{
	//	"code":    models.ResponseCodeSuccess,
	//	"message": models.ResponseMessageSuccess,
	//})
}

// LoginHandler 处理用户登录
func LoginHandler(ctx *gin.Context) {
	p := new(models.ParamLogin)
	//参数校验
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with FailedBindParam", zap.Error(err))
		msg := ""
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, validationError := range errs {
				msg = fmt.Sprintf("%s,%s", msg, validationError.Translate(trans))
			}
		}
		ResponseErrorWithMsg(ctx, CodeFailedBindParam, msg)
		//ctx.JSON(200, gin.H{
		//	"code":    models.ResponseCodeFailedBindParam,
		//	"message": msg,
		//})
		return
	}
	//对比账号密码
	if err := logic.Login(p); err != nil {
		zap.L().Info("Login with Failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorIncorrectPassword) {
			ResponseError(ctx, CodePasswordIncorrect)
		} else {
			ResponseError(ctx, CodeBusy)
		}

		//ctx.JSON(200, gin.H{
		//	"code":    models.ResponseCodePasswordIncorrect,
		//	"message": models.ResponseMessagePasswordIncorrect,
		//})
		return
	}
	//返回结果
	zap.L().Info("Login with Success")
	ResponseSuccess(ctx, nil)
	//ctx.JSON(200, gin.H{
	//	"code":    models.ResponseCodeSuccess,
	//	"message": models.ResponseMessageSuccess,
	//})

}
