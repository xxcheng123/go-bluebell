package mysql

import (
	"go-generator/models"
	"go-generator/pkg/encrypt"
)

// QueryUserExistByUsername 根据用户名查询用户是否存在
func QueryUserExistByUsername(username string) (isExist bool, err error) {
	sqlStr := `SELECT COUNT(user_id) FROM user WHERE username=?`
	var count int64
	if err = db.Get(&count, sqlStr, username); err != nil {
		return true, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func InsertUser(user *models.User) (err error) {
	sqlStr := `INSERT INTO user(user_id, username, password) values(?,?,?)`
	encryptPassword := encrypt.MD5(user.Password)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, encryptPassword)
	return err
}
