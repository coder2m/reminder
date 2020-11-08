/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/8 16:40
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/myxy99/reminder/internal/reminder-apisvr/jwt"
	R "github.com/myxy99/reminder/pkg/response"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) < 7 {
			c.Abort()
			R.Response(c, http.StatusUnauthorized, "未登录", nil, http.StatusUnauthorized)
			return
		}
		jwtUserInfo := jwt.Info{}
		err := jwtUserInfo.ParseToken(token[7:])
		if err != nil {
			R.Response(c, http.StatusUnauthorized, "未登录", nil, http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("openid", jwtUserInfo.Openid)
		c.Next()
		return
	}
}
