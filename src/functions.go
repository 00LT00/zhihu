package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"time"
)

// 生成token
func (s *Service) makeAuth(UserID string) (string, error) {
	AccessToken := uuid.NewV4().String()
	tx := s.DB.Begin()
	var Token token
	tx.Where(token{UserID: UserID}).Assign(token{AccessToken: AccessToken}).FirstOrCreate(&Token)
	if Token.UserID == "" || Token.AccessToken == "" {
		tx.Rollback()
		return "", errors.New("insert null")
	}
	tx.Commit()
	return AccessToken, nil
}

//	生成id
func (s *Service) makeID(strs []string) string {
	//	UserID :UserName
	//	ArticleID: UserID[:4],Title

	sum := ""
	for _, str := range strs {
		sum += str
		sum += "."
	}
	sum += strconv.FormatInt(time.Now().Unix(), 36)
	fmt.Println(sum)
	id := base64.StdEncoding.EncodeToString([]byte(sum))
	fmt.Println(id)
	return id
}

// 	获取UserID
func (s *Service) GetUserID(AccessToken string) (string, error) {
	if AccessToken == "" {
		return "", errors.New("token null")
	}
	var Token token
	s.DB.Where(token{AccessToken: AccessToken}).Find(&Token)
	if Token.UserID == "" {
		return "", errors.New("token error")
	}
	return Token.UserID, nil
}
