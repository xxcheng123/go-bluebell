package mysql

import (
	"errors"
	"go-generator/models"
	"go-generator/pkg/encrypt"
)

var (
	ErrorUserExist         = errors.New("用户已存在")
	ErrorIncorrectPassword = errors.New("密码错误")
)

// QueryUserExistByUsername 根据用户名查询用户是否存在
func QueryUserExistByUsername(username string) error {
	sqlStr := `SELECT COUNT(user_id) FROM user WHERE username=?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

func InsertUser(user *models.User) (err error) {
	sqlStr := `INSERT INTO user(user_id, username, password) values(?,?,?)`
	encryptPassword := encrypt.MD5(user.Password)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, encryptPassword)
	return err
}

func CompareUserPassword(username string, password string) error {
	encryptPassword := encrypt.MD5(password)
	sqlStr := "SELECT COUNT(user_id) FROM user WHERE username=? AND password=?"
	compareCount := 0
	if err := db.QueryRow(sqlStr, username, encryptPassword).Scan(&compareCount); err != nil {
		return err
	}
	if compareCount == 0 {
		return ErrorIncorrectPassword
	}
	return nil
}
