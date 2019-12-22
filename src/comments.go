package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func (s *Service) AddComment(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	UserID, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	ParentID := c.Request.FormValue("parentid")
	Content := c.Request.FormValue("content")
	CommentID := s.makeID([]string{
		ParentID[:9],
		UserID[:4],
	})
	var tempComment comment
	var TargetID string
	s.DB.Where(comment{CommentID: ParentID}).Find(&tempComment)
	fmt.Println(tempComment)
	if tempComment.TargetID == "" {
		TargetID = CommentID
		fmt.Println("aaa")
	} else {
		TargetID = tempComment.TargetID
		fmt.Println("bbb")
	}
	tx := s.DB.Begin()
	if err := tx.Create(&comment{CommentID: CommentID, TargetID: TargetID, UserID: UserID, ParentID: ParentID, Content: Content}).Error; err != nil {
		tx.Rollback()
		return s.makeErrJSON(500, 50000, err.Error())
	}
	tx.Commit()
	return s.makeSuccessJSON(gin.H{
		"CommentID": CommentID,
	})
}

func (s *Service) GetAllComments(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	ParentID := c.Query("id")
	flag := c.Param("flag")
	_, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var db *gorm.DB
	switch flag {
	case "new":
		db = s.DB.Where(comment{ParentID: ParentID}).Or("target_id in (SELECT comment_id from comments WHERE parent_id = ?)", ParentID).Order("updated_at desc")
		break
	case "hot":
		db = s.DB.Where(comment{ParentID: ParentID}).Or("target_id in (SELECT comment_id from comments WHERE parent_id = ?)", ParentID).Order("up_vote_count - down_vote_count desc").Limit(10)
		break
	default:
		db = s.DB.Where(comment{ParentID: ParentID}).Or("target_id in (SELECT comment_id from comments WHERE parent_id = ?)", ParentID).Order("updated_at")
	}
	var Comments []comment
	err = db.Find(&Comments).Error
	if err != nil {
		return s.makeErrJSON(500, 50000, err.Error())
	}
	for i := 0; i < len(Comments); i++ {
		s.DB.Select("user_id,nick_name,avatar,introduction").Where(user{UserID: Comments[i].UserID}).Find(&Comments[i].User)
	}
	return s.makeSuccessJSON(Comments)
}

func (s *Service) GetCommentsByTargetID(c *gin.Context) (int, interface{}) {
	AccessToken := c.GetHeader("Authorization")
	TargetID := c.Query("targetid")
	_, err := s.GetUserID(AccessToken)
	if err != nil {
		return s.makeErrJSON(403, 40301, err.Error())
	}
	var Comments []comment
	s.DB.Where(comment{TargetID: TargetID}).Order("updated_at").Find(&Comments)
	for i := 0; i < len(Comments); i++ {
		s.DB.Select("user_id,nick_name,avatar,introduction").Where(user{UserID: Comments[i].UserID}).Find(&Comments[i].User)
	}
	return s.makeSuccessJSON(Comments)
}
