package longpoll

import "github.com/gin-gonic/gin"

func (c *Component) AdaptHandler(ctx *gin.Context) {
	c.lpManager.SubscriptionHandler(ctx.Writer, ctx.Request)
}
