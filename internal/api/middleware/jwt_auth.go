package middleware

import (
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/token_verify"
	"Open_IM/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		operationId := utils.OperationIDGenerator()
		ok, userId, _ := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), operationId)
		if !ok || len(userId) == 0 {
			log.NewError("", "GetUserIDFromToken false ", c.Request.Header.Get("token"))
			c.Abort()
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "msg": "token授权认证失败"})
			return
		} else {
			// 用户id
			c.Set("userId", userId)

			//operationID
			c.Set("operationId", operationId)
		}
	}
}
