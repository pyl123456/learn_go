package v1

import (
    "github.com/gin-gonic/gin"
    "week04/internal/account/db"
    "week04/internal/account/service"
)

func RegisterUser(c *gin.Context) {
    // 邮件配置
    dbConfig := &db.Config{}
    // 邮件配置
    mailConfig := &service.MailConfig{}
    s, err := NewService(dbConfig, mailConfig)
    if err != nil {
        c.JSON(400, gin.H{
            "code":-1,
            "message": "registerd user failed",
        })
        return
    }
    s.UserSignUp()

    c.JSON(200, gin.H{
        "code":-1,
        "message": "success",
    })
}