package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//	根据问题id获取回答
func (s *Service) GetAnswersByQuestionID(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	QuestionID := c.Query("questionid")
	flag := c.Param("flag")
	_, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var db *gorm.DB
	switch flag {
	case "new": //根据时间，新的在上
		db = s.DB.Where(answer{QuestionID: QuestionID}).Order("updated_at desc")
		break
	case "hot": //根据赞数
		db = s.DB.Where(answer{QuestionID: QuestionID}).Order("up_vote_count-down_vote_count desc")
	default:
		db = s.DB.Where(answer{QuestionID: QuestionID})
	}
	var Answers []answer
	err = db.Find(&Answers).Error
	if err != nil {
		return s.makeErrJSON(500, 50000, err.Error())
	}
	for i := 0; i < len(Answers); i++ {
		s.DB.Select("user_id,nick_name,avatar,introduction").Where(user{UserID: Answers[i].UserID}).Find(&Answers[i].User)
	}
	return s.makeSuccessJSON(Answers)
}

//获取用户的回答
func (s *Service) GetAnswersByToken(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	UserID, err := s.GetUserID(AccessToken)
	flag := c.Param("flag")
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var db *gorm.DB
	switch flag {
	case "new": //根据时间，新的在上
		db = s.DB.Where(answer{UserID: UserID}).Order("updated_at desc")
		break
	case "hot": //根据赞数
		db = s.DB.Where(answer{UserID: UserID}).Order("up_vote_count-down_vote_count desc")
	default:
		db = s.DB.Where(answer{UserID: UserID})
	}
	var Answers []answer
	err = db.Find(&Answers).Error
	if err != nil {
		return s.makeErrJSON(500, 50000, err.Error())
	}
	for i := 0; i < len(Answers); i++ {
		s.DB.Select("user_id,nick_name,avatar,introduction").Where(user{UserID: Answers[i].UserID}).Find(&Answers[i].User)
	}
	return s.makeSuccessJSON(Answers)
}

//	发表回答
func (s *Service) AddAnswer(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	UserID, err := s.GetUserID(AccessToken)
	QuestionID := c.Request.FormValue("questionid")
	Content := c.Request.FormValue("content")
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	AnswerID := s.makeID([]string{
		QuestionID[:9],
		UserID[:4],
	})
	tx := s.DB.Begin()
	if err := tx.Create(&answer{AnswerID: AnswerID, UserID: UserID, Content: Content, QuestionID: QuestionID}).Error; err != nil {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(gin.H{
		"AnswerID": AnswerID,
	})
}

//	根据回答id获取回答
func (s *Service) GetAnswerByID(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	AnswerID := c.Query("answerid")
	_, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var Answer answer
	if s.DB.Where(answer{AnswerID: AnswerID}).Find(&Answer).RowsAffected != 1 {
		return s.makeErrJSON(500, 50000, "AnswerID error")
	}
	return s.makeSuccessJSON(Answer)
}
