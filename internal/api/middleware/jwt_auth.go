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
		ok, userId, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), operationId)
		if !ok || len(userId) == 0 {
			log.NewError("", "GetUserIDFromToken false ", c.Request.Header.Get("token"))
			c.Abort()
			c.JSON(http.StatusOK, gin.H{"errCode": 403, "errMsg": errInfo})
			return
		} else {
			// 用户id
			c.Set("userId", userId)

			//operationID
			c.Set("operationId", operationId)
		}
	}
}
