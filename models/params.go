package models

type ParamSignUp struct {
	Username   string `json:"username" binding:"required,gte=3,lte=8"`
	Password   string `json:"password" binding:"required,gte=8,lte=16"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
