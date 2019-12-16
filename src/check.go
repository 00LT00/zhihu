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