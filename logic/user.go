package logic

import (
	"go-generator/dao/mysql"
	"go-generator/models"
	"go-generator/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) error {
	//fmt.Println("处理用户注册...")
	//判断用户是否存在
	if err := mysql.QueryUserExistByUsername(p.Username); err != nil {
		return err
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

func Login(p *models.ParamLogin) error {
	err := mysql.CompareUserPassword(p.Username, p.Password)
	if err != nil {
		return err
	}
	return nil
}
