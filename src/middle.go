package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) check(c *gin.Context) {
	key := c.GetHeader("sign")
	if key != s.Conf.Server.Key {
		c.JSON(s.makeErrJSON(403, 40300, "forbidden"))
		c.Abort()
		return
	}
	return
}

func(s *Service) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
