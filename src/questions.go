package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//	发问题
func (s *Service) AddQuestion(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	UserID, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	Title := c.Request.FormValue("title")
	var QuestionID string
	if len(Title) >= 9 {
		QuestionID = s.makeID([]string{
			Title[:9],
			UserID[:4],
		})
	} else {
		QuestionID = s.makeID([]string{
			Title,
			UserID[:4],
		})
	}
	Content := c.Request.FormValue("content")
	//Topic:=c.Request.FormValue("topic")
	tx := s.DB.Begin()
	if err := tx.Create(&question{QuestionID: QuestionID, UserID: UserID, Title: Title, Content: Content}).Error; err != nil {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(gin.H{
		"QuestionID": QuestionID,
	})
}

//	获取某人问题
func (s *Service) GetQuestion(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	UserID, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var Question []question
	err = s.DB.Where(question{UserID: UserID}).Find(&Question).Error
	if err != nil {
		return s.makeErrJSON(500, 50000, err.Error())
	}
	return s.makeSuccessJSON(Question)
}

//	根据问题id获取
func (s *Service) GetQuestionByID(c *gin.Context) (int, interface{}) {
	QuestionID := c.Param("id")
	AccessToken := c.GetHeader("Authorization")
	_, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var Question question
	s.DB.Where(question{QuestionID: QuestionID}).Find(&Question)
	return s.makeSuccessJSON(Question)
}

//	获取全部问题
func (s *Service) GetQuestions(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	flag := c.Param("flag")
	_, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var db *gorm.DB
	switch flag {
	case "":
		db = s.DB
		break
	case "hot":
		db = s.DB.Order("view_count+followers_count+favorite_count desc").Limit(10)
		break
	default:
		db = s.DB
	}
	var Question []question
	err = db.Find(&Question).Error
	if err != nil {
		return s.makeErrJSON(500, 50000, err.Error())
	}
	return s.makeSuccessJSON(Question)
}
