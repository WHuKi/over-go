package main

import (
	"net/http"
	"over-go/sign"

	"github.com/gin-gonic/gin"
)

/*
post 请求参数验签 && 重放攻击防范
*/

func main() {
	router := gin.Default()
	// 签名

	newGroup := router.Group("/")
	newGroup.Use(gin_use.SetUp())

	// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
	newGroup.POST("/user", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.Run(":8088")
}
