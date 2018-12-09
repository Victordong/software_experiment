package plugin

import (
	"context"
	"github.com/gin-gonic/gin"
)

func SetContext(c *gin.Context, ctx context.Context) context.Context {
	for key, value := range c.Keys {
		ctx = context.WithValue(ctx, key, value)
	}
	return ctx
}
