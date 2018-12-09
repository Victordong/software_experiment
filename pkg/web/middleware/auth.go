package middleware

import (
	"auto_fertilizer_back/pkg/comm/model"
	"auto_fertilizer_back/pkg/web/controller"
	"context"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "username"

var AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
	Realm:       "test zone",
	Key:         []byte("secret key"),
	Timeout:     300 * time.Hour,
	MaxRefresh:  time.Hour,
	IdentityKey: identityKey,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*model.OperatorModel); ok {
			return jwt.MapClaims{
				identityKey: v.Username,
			}
		}
		return jwt.MapClaims{}
	},
	IdentityHandler: func(c *gin.Context) interface{} {
		ctx := context.Background()
		claims := jwt.ExtractClaims(c)
		user, _ := controller.GetOperatorByUsername(ctx, claims["username"].(string))
		return user
	},
	Authenticator: func(c *gin.Context) (interface{}, error) {
		ctx := context.Background()
		var loginJSON login
		if err := c.ShouldBind(&loginJSON); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		user, err := controller.GetOperatorByUsername(ctx, loginJSON.Username)
		if err != nil {
			return nil, err
		}
		if user.PasswordHash == controller.GenPasswordHash(ctx, loginJSON.Password) {
			return user, nil
		}
		return nil, jwt.ErrFailedAuthentication
	},
	Authorizator: func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*model.OperatorModel); ok {
			c.Set("currentUser", v)
			return true
		}
		return false
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "账号或密码错误",
		})
		return
	},
	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	TokenLookup: "header:Authorization",
	// TokenLookup: "query:token",
	// TokenLookup: "cookie:token",

	// TokenHeadName is a string in the header. Default value is "Bearer"
	TokenHeadName: "Bearer",

	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	TimeFunc: time.Now,
})
