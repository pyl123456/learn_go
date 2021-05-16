package router

import (
	"github.com/gin-gonic/gin"
	"week04/cmd/account/router/v1"
)

func InitRouter(c *gin.Engine)  {
	GroupV1 := c.Group("/week04/api/").Use(gin.Recovery())
	{
		GroupV1.POST("/user_register", v1.RegisterUser)
	}
	
}
