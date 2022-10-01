package router

import (
	"github.com/gin-gonic/gin"
	"github.com/youichiro/go-slack-my-unipos/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	slackHander := handler.SlackHandler{}

	r.POST("/", func(c *gin.Context) { slackHander.Receive(c) })

	return r
}
