package main

import "github.com/gin-gonic/gin"

// 获取token
func (s *Service) Auth(c *gin.Context)(int,interface{}){
	UserName:= c.Query("username")
	PassWord:= c.Query("password")
	if UserName == ""|| PassWord == ""{
		return s.makeErrJSON(404,40400,"query null")
	}
	User:= new(user)
	s.DB.Where(user{UserName:UserName}).Find(User)
	if User.PassWord != PassWord {
		return s.makeErrJSON(403,40301,"pass error")
	}
	token,err :=s.makeAuth(User.UserID)
	if err != nil {
		return s.makeErrJSON(500,50000,err.Error())
	}
	return s.makeSuccessJSON(gin.H{
		"Auth":token,
		"UserID":User.UserID,
	})
}