package middleware

import (
	"context"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"software_experiment/pkg/comm/model"
	"software_experiment/pkg/web/controller"
	"software_experiment/pkg/web/plugin"
	"time"
)

type login struct {
	Username      string `form:"username" json:"username" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
	CaptchaId     string `json:"captcha_id" binding:"required"`
	CaptchaResult string `json:"captcha_result"`
}

var identityKey = "username"

var AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
	Realm:       "test zone",
	Key:         []byte("secret key"),
	Timeout:     300 * time.Hour,
	MaxRefresh:  time.Hour,
	IdentityKey: identityKey,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*model.UserModel); ok {
			return jwt.MapClaims{
				identityKey: v.Username,
			}
		}
		return jwt.MapClaims{}
	},
	IdentityHandler: func(c *gin.Context) interface{} {
		ctx := context.Background()
		claims := jwt.ExtractClaims(c)
		user, _ := controller.GetUserByUsername(ctx, claims["username"].(string))
		return user
	},
	Authenticator: func(c *gin.Context) (interface{}, error) {
		ctx := context.Background()
		var loginJSON login
		if err := c.ShouldBind(&loginJSON); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		user, err := controller.GetUserByUsername(ctx, loginJSON.Username)
		if err != nil {
			return nil, err
		}
		ifSuccess, err := controller.VerifyCaptcha(ctx, model.CaptchaModel{CaptchaId: loginJSON.CaptchaId, Result: loginJSON.CaptchaResult})
		if err != nil {
			return nil, err
		}
		if !ifSuccess {
			return nil, plugin.CustomErr{
				Code:        500,
				StatusCode:  200,
				Information: "验证错误",
			}
		}
		if user.PasswordHash == controller.GenPasswordHash(ctx, loginJSON.Password) {
			return user, nil
		}
		return nil, jwt.ErrFailedAuthentication
	},
	Authorizator: func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*model.UserModel); ok {
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
