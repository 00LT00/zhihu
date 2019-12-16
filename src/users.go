package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 注册
func(s *Service)Register(c *gin.Context)(int,interface{}){
	UserName := c.Request.FormValue("UserName")
	PassWord := c.Request.FormValue("PassWord")
	Phone := c.Request.FormValue("Phone")
	Email := c.Request.FormValue("Email")
	NickName := c.Request.FormValue("NickName")
	if UserName == ""||PassWord == ""||Phone == ""||Email == ""||NickName == "" {
		fmt.Println(map[string]string{
			"name":UserName,
			"pass":PassWord,
			"phone":Phone,
			"email":Email,
			"nickname":NickName,
		})
		return s.makeErrJSON(403,40301,"query null")
	}
	//	防止用户名重复
	count:=0
	s.DB.Where(user{UserName:UserName}).Count(&count)
	if count !=0 {
		return s.makeErrJSON(403,40302,"Duplicate username")
	}

	UserID := s.makeID([]string{UserName})
	fmt.Println(UserID)
	tx := s.DB.Begin()
	if err:= tx.Create(&user{
		UserID:         UserID,
		UserName:       UserName,
		PassWord:       PassWord,
		Phone:          Phone,
		Email:          Email,
		NickName:       NickName,
		FollowersCount: 0,
		FavoritesCount: 0,
		FansCount:      0,
	}).Error ;err !=nil {
		tx.Rollback()
		return s.makeErrJSON(500,50000,err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(gin.H{
		"UserName":UserName,
		"UserID":UserID,
	})
}

//	获取用户详情
func(s *Service)GetUser(c *gin.Context)(int,interface{}){
	UserID:=c.Param("id")
	AccessToken := c.GetHeader("Authorization")
	loginUserID,err:=s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403,40301,err.Error())
	}
	var User user
	//	本人获取
	if UserID == "" {
		s.DB.Where(user{UserID:loginUserID}).Find(&User)
		return s.makeSuccessJSON(User)
	}

	s.DB.Where(user{UserID: UserID}).Find(&User)
	//	本人获取
	if loginUserID == UserID {
		return s.makeSuccessJSON(User)
	}
	//	他人获取
	//	昵称，手机号，邮箱
	return s.makeSuccessJSON(user{NickName:User.NickName,Phone:User.Phone,Email:User.Email})
}

//	更改个人信息
func(s *Service)Change(c *gin.Context)(int,interface{}){
	AccessToken:=c.GetHeader("Authorization")
	UserID:=c.Param("id")
	PassWord:=c.Request.FormValue("password")
	Phone:=c.Request.FormValue("phone")
	Email:=c.Request.FormValue("email")
	Avatar:=c.Request.FormValue("avatar")
	NickName:=c.Request.FormValue("nickname")
	Introduction:=c.Request.FormValue("Introduction")
	loginUserID,err:=s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403,40301,err.Error())
	}
	if loginUserID != UserID {
		return s.makeErrJSON(403,40302,"loginID != ID")
	}
	tx:=s.DB.Begin()
	var User user
	if tx.Model(&User).Updates(user{PassWord:PassWord,Phone:Phone,Email:Email,Avatar:Avatar,NickName:NickName,Introduction:Introduction}).RowsAffected !=1 {
		tx.Rollback()
		return s.makeErrJSON(500,50000,"update error")
	}
	tx.Commit()
	s.DB.Where(user{UserID:UserID}).Find(&User)
	return s.makeSuccessJSON(User)
}