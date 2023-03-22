package cloud_wallet

import (
	api "Open_IM/pkg/base_info"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/token_verify"
	rpc "Open_IM/pkg/proto/friend"
	"Open_IM/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 检查用户是否有账号
// @Description 检查用户是否有账号
// @Tags 云钱包
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func CheckUserHaveAccount(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func CreateUserAccount(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func CheckUserAccountBalance(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func CheckUserAccountBalanc(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func GetAccountChangeList(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

} // @Summary 用户注册

// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func GetUserBankCardList(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

} //

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func AddUserBankCard(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

} //

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func DelUserBankCard(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func ChargeAccount(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func DrawAccount(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func SendRedPacket(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func ClickRedPacket(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func GetRedPacketInfo(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func RedPacketClickDetail(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func ListRedPacketRecord(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Summary 用户注册
// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func SetPaymentSecret(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func SendRedPacketNotify(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func DrawNotify(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func DelRedPacketRecord(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func DelAccountChangeRecord(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}

// @Description 用户注册
// @Tags 鉴权认证
// @ID UserRegister
// @Accept json
// @Param req body api.UserRegisterReq true "secret为openIM密钥, 详细见服务端config.yaml secret字段 <br> platform为平台ID <br> ex为拓展字段 <br> gender为性别, 0为女, 1为男"
// @Produce json
// @Success 0 {object} api.UserRegisterResp
// @Failure 500 {object} api.Swagger500Resp "errCode为500 一般为服务器内部错误"
// @Failure 400 {object} api.Swagger400Resp "errCode为400 一般为参数输入错误, token未带上等"
// @Router /auth/user_register [post]
func ConfirmSendRedPacketCode(c *gin.Context) {
	params := api.ImportFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.NewError("0", "BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 400, "errMsg": err.Error()})
		return
	}
	req := &rpc.ImportFriendReq{}
	utils.CopyStructFields(req, &params)
	var ok bool
	var errInfo string
	ok, req.OpUserID, errInfo = token_verify.GetUserIDFromToken(c.Request.Header.Get("token"), req.OperationID)
	if !ok {
		errMsg := req.OperationID + " " + "GetUserIDFromToken failed " + errInfo + " token:" + c.Request.Header.Get("token")
		log.NewError(req.OperationID, errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"errCode": 500, "errMsg": errMsg})
		return
	}

}