package account

import (
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/token_verify"
	"Open_IM/pkg/proto/cloud_wallet"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取账户信息
func Account(c *gin.Context) {

	info := cloud_wallet.UserNcountAccountResp{
		Step:          2,
		RealAuth:      0,
		IdCard:        "234",
		RealName:      "23424",
		AccountStatus: 0,
		CommonResp:    nil,
	}
	c.JSON(http.StatusOK, info)
	return
}

// 身份证实名认证
func IdCardRealNameAuth(c *gin.Context) {
	params := IdCardRealNameAuth{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}

	if len(params.RealName) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "实名姓名错误"})
		return
	}

	//验证身份证
	if !verifyByIDCard(params.IdCard) {
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": "身份证号码错误"})
		return
	}

	//获取token用户id
	ok, userId, errInfo := token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), params.OperationID)
	if !ok {
		log.NewError(params.OperationID, fmt.Sprintf(params.OperationID+" "+"GetUserIDFromToken failed "+errInfo+" token:"+c.Request.Header.Get("token")))
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": errInfo})
		return
	}

	//数据入库
	params.UserId = userId

	c.JSON(http.StatusOK, params)
	return

}

func verifyByIDCard(idCard string) bool {
	if len(idCard) != 18 {
		return false
	}
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	validate := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	sum := 0
	for i := 0; i < len(weight); i++ {
		sum += weight[i] * int(byte(idCard[i])-'0')
	}
	m := sum % 11
	return validate[m] == idCard[17]
}
