package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (s *Service) message(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	ID := c.Request.FormValue("id")
	Type := c.Request.FormValue("type")
	action := c.Request.FormValue("action")
	UserID, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var column string
	switch action {
	case "up":
		column = "up_vote_count"
		break
	case "down":
		column = "down_vote_count"
		break
	default:
		return s.makeErrJSON(403, 40302, "action error")
	}
	Tx := s.DB.Begin()
	var tx *gorm.DB
	Tx.Create(&message{UserID: UserID, Action: action, TargetID: ID})
	switch Type {
	case "article":
		tx = Tx.Model(article{}).Where(article{ArticleID: ID})
		break
	case "comment":
		tx = Tx.Model(comment{}).Where(comment{CommentID: ID})
		break
	case "answer":
		tx = Tx.Model(answer{}).Where(answer{AnswerID: ID})
		break
	default:
		return s.makeErrJSON(403, 40303, "type error")
	}
	err = tx.Update(column, gorm.Expr(column+"+1")).Error
	if err != nil {
		tx.Callback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON("successful")
}
