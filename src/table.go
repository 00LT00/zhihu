package main

import "github.com/jinzhu/gorm"

type user struct {
	gorm.Model
	UserID         string `gorm:"index;unique"`
	UserName       string `gorm:"not null;unique"`
	PassWord       string `gorm:"not null"`
	Phone          string `gorm:"not null"`
	Email          string `gorm:"not null"`
	Avatar         string
	NickName       string `gorm:"not null"`
	Introduction   string
	FollowersCount int64
	FavoritesCount int64
	FansCount      int64
}

type question struct {
	gorm.Model
	QuestionID     string `gorm:"index;unique"`
	UserID         string `gorm:"not null"`
	User           user   `gorm:"ForeignKey:UserID"`
	Title          string `gorm:"not null"`
	Content        string `gorm:"not null;type:text"`
	CommentCount   int64
	AnswerCount    int64
	ViewCount      int64
	FollowersCount int64
	FavoriteCount  int64
}

type answer struct {
	gorm.Model
	AnswerID       string `gorm:"index;unique"`
	QuestionID     string `gorm:"not null"`
	UserID         string `gorm:"not null"`
	User           user   `gorm:"ForeignKey:UserID"`
	Content        string `gorm:"not null;type:text"`
	CommentCount   int64
	UpVoteCount    int64
	DownVoteCount  int64
	FollowersCount int64
}

type article struct {
	gorm.Model
	ArticleID      string `gorm:"index;unique"`
	UserID         string `gorm:"not null"`
	User           user   `gorm:"ForeignKey:UserID"`
	Title          string `gorm:"not null"`
	Content        string `gorm:"not null;type:text"`
	UpVoteCount    int64
	DownVoteCount  int64
	FollowersCount int64
}

type comment struct {
	gorm.Model
	CommentID      string `gorm:"index;unique"`
	TargetID       string `gorm:"not null"`
	UserID         string `gorm:"not null"`
	User           user   `gorm:"ForeignKey:UserID"`
	ParentID       string `gorm:"not null"`
	Content        string `gorm:"not null"`
	Count          int64
	UpVoteCount    int64
	DownVoteCount  int64
	FollowersCount int64
}

type topic struct {
	gorm.Model
	TopicID        string `gorm:"index;unique"`
	TopicName      string `gorm:"not null"`
	Description    string `gorm:"not null"`
	FollowersCount int64
}

//动作，描述谁对谁干了什么
type message struct {
	gorm.Model
	UserID   string `gorm:"not null"`
	User     user   `gorm:"ForeignKey:UserID"`
	Action   string `gorm:"not null"`
	TargetID string `gorm:"not null"`
}

type topic_to_question struct {
	gorm.Model
	QuestionID string
	Question   question `gorm:"ForeignKey:QuestionID"`
	TopicID    string
	Topic      topic `gorm:"ForeignKey:TopicID"`
}

type topic_to_article struct {
	gorm.Model
	ArticleID string
	Article   article `gorm:"ForeignKey:ArticleID"`
	TopicID   string
	Topic     topic `gorm:"ForeignKey:TopicID"`
}

type token struct {
	gorm.Model
	AccessToken string `gorm:"unique;index"`
	UserID      string `gorm:"unique"`
	User        user   `gorm:"ForeignKey:UserID"`
}
