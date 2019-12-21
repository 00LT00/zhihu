package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//	发文章
func (s *Service) AddArticle(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	UserID, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	Title := c.Request.FormValue("title")
	var ArticleID string
	if len(Title) >= 9 {
		ArticleID = s.makeID([]string{
			Title[:9],
			UserID[:4],
		})
	} else {
		ArticleID = s.makeID([]string{
			Title,
			UserID[:4],
		})
	}
	Content := c.Request.FormValue("content")
	//Topic:=c.Request.FormValue("topic")
	tx := s.DB.Begin()
	if err := tx.Create(&article{ArticleID: ArticleID, UserID: UserID, Title: Title, Content: Content}).Error; err != nil {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(gin.H{
		"ArticleID": ArticleID,
	})
}

//	获取某人文章
func (s *Service) GetArticle(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	UserID, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var Articles []article
	err = s.DB.Where(article{UserID: UserID}).Find(&Articles).Error
	if err != nil {
		return s.makeErrJSON(500, 50000, err.Error())
	}
	return s.makeSuccessJSON(Articles)
}

//	根据文章id获取
func (s *Service) GetArticleByID(c *gin.Context) (int, interface{}) {
	ArticleID := c.Param("id")
	AccessToken := c.GetHeader("Authorization")
	_, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var Article article
	s.DB.Where(article{ArticleID: ArticleID}).Find(&Article)
	return s.makeSuccessJSON(Article)
}

//	获取全部文章
func (s *Service) GetArticles(c *gin.Context) (int, interface{}) {
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
		db = s.DB.Order("up_vote_count+followers_count-down_vote_count desc").Limit(10)
		break
	default:
		db = s.DB
	}
	var Articles []article
	err = db.Find(&Articles).Error
	if err != nil {
		return s.makeErrJSON(500, 50000, err.Error())
	}
	return s.makeSuccessJSON(Articles)
}
