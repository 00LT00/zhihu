package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() {
	r := gin.Default()
	r.Use(cors.Default())

	/* 登陆获取token*/
	r.GET("/auth/login",
		func(c *gin.Context) {
			s.check(c)
		},
		func(c *gin.Context) {
			c.JSON(s.Auth(c))
		})

	/* 用户模块*/
	users := r.Group("/users",
		func(c *gin.Context) {
			s.check(c)
		})
	//	注册
	users.PUT("/",
		func(c *gin.Context) {
			c.JSON(s.Register(c))
		})
	//	更改
	users.PATCH("/:id",
		func(c *gin.Context) {
			c.JSON(s.Change(c))
		})
	//	获取用户详情
	users.GET("/:id",
		func(c *gin.Context) {
			c.JSON(s.GetUser(c))
		})

	/*	文章模块*/
	articles := r.Group("/articles",
		func(c *gin.Context) {
			s.check(c)
		})
	//	发布文章
	articles.POST("/", func(c *gin.Context) {
		c.JSON(s.AddArticle(c))
	})
	//	查找某人的文章
	articles.GET("/", func(c *gin.Context) {
		c.JSON(s.GetArticle(c))
	})
	//	通过文章id查找
	articles.GET("/id/:id", func(c *gin.Context) {
		c.JSON(s.GetArticleByID(c))
	})
	//	查找全部文章，可选多种排序
	articles.GET("/all/:flag", func(c *gin.Context) {
		c.JSON(s.GetArticles(c))
	})

	/*	问题模块*/
	questions := r.Group("/questions",
		func(c *gin.Context) {
			s.check(c)
		})
	//	发布问题
	questions.POST("/", func(c *gin.Context) {
		c.JSON(s.AddQuestion(c))
	})
	//	查找某人的问题
	questions.GET("/", func(c *gin.Context) {
		c.JSON(s.GetQuestion(c))
	})
	//	通过问题id查找
	questions.GET("/id/:id", func(c *gin.Context) {
		c.JSON(s.GetQuestionByID(c))
	})
	//	查找全部问题，可选多种排序
	questions.GET("/all/:flag", func(c *gin.Context) {
		c.JSON(s.GetQuestions(c))
	})

	/* 	回答模块*/
	answers := r.Group("/answers",
		func(c *gin.Context) {
			s.check(c)
		})
	//	发布回答
	answers.POST("/", func(c *gin.Context) {
		c.JSON(s.AddAnswer(c))
	})
	//	根据回答id获取回答
	answers.GET("/", func(c *gin.Context) {
		c.JSON(s.GetAnswerByID(c))
	})
	//	获得问题的答案
	answers.GET("/question/:flag", func(c *gin.Context) {
		c.JSON(s.GetAnswersByQuestionID(c))
	})
	//	获得用户的所有回答
	answers.GET("/user/:flag", func(c *gin.Context) {
		c.JSON(s.GetAnswersByToken(c))
	})

	/*	评论模块*/
	comments := r.Group("/comments",
		func(c *gin.Context) {
			s.check(c)
		})
	//	发评论
	comments.POST("/", func(c *gin.Context) {
		c.JSON(s.AddComment(c))
	})
	//	根据文章id或者问题回答id获取评论
	comments.GET("/id/:flag", func(c *gin.Context) {
		c.JSON(s.GetAllComments(c))
	})
	//	根据评论id获取评论详情
	comments.GET("/target/", func(c *gin.Context) {
		c.JSON(s.GetCommentsByTargetID(c))
	})

	s.Router = r
	err := s.Router.Run(s.Conf.Server.Port)
	panic(err)
}
