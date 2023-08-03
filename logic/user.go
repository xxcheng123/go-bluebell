package logic

import (
	"errors"
	"fmt"
	"go-generator/dao/mysql"
	"go-generator/models"
	"go-generator/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	fmt.Println("处理用户注册...")
	//判断用户是否存在
	if isExist, err := mysql.QueryUserExistByUsername(p.Username); err != nil {
		return err
	} else if isExist {
		return errors.New("用户名已存在")
	}
	//生成UID
	UID := snowflake.GenID()
	//存入数据库
	user := &models.User{
		UserID:   UID,
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)
}
