package user

import (
	"Open_IM/pkg/utils"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	AppId  = "xnFiQUGP"
	AppKey = "xnFiQUGP"
	Uri    = "https://api.253.com/open/flashsdk/mobile-query"
)

type OneClickLoginReq struct {
	AppId       string `json:"appId"`
	Token       string `json:"token"`
	EncryptType string `json:"encryptType"` //手机号码加密方式 0（AES加密）、1（RSA加密）
	Sign        string `json:"sign"`
}

type OneClickLoginResp struct {
	Code         string      `json:"code"`         //200000表示成功，其他代码都为失败，详情参考附录
	ChargeStatus int         `json:"chargeStatus"` //是否收费，枚举值：1:收费/0:不收费
	Message      string      `json:"message"`
	Data         *UserMobile `json:"data"`
}

type UserMobile struct {
	TradeNo    string `json:"tradeNo"`    //闪验的交易流水号
	MobileName string `json:"mobileName"` //交易流水号
}

// 一键登录授权token换取用户手机号码
func TokenExchangeMobile(token string) (string, error) {
	signStr := fmt.Sprintf("appId%stoken%s", AppId, token)
	sign := Sign(signStr)

	//封装参数、请求创蓝api
	resp, err := http.PostForm(Uri, url.Values{"appId": {AppId}, "token": {token}, "encryptType": {"0"}, "sign": {sign}})
	if err != nil {
		return "", errors.New("请求api接口失败")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("请求api接口失败")
	}

	reply := &OneClickLoginResp{}
	err = json.Unmarshal(body, reply)
	if err != nil {
		return "", errors.New("请求api接口失败")
	}

	mobile, err := Decrypt(reply.Data.MobileName)
	if err != nil {
		return "", errors.New("解密手机号码失败")
	}

	return mobile, nil
}

// 签名
func Sign(text string) string {
	hMac := hmac.New(sha256.New, []byte(AppKey))
	hMac.Write([]byte(text))
	return hex.EncodeToString(hMac.Sum(nil))
}

func Decrypt(text string) (string, error) {
	key := utils.Md5(AppKey)
	//如未填写则只能使用AES CBC算法，以md5(appKey)前16位字符串为秘钥，后16位字符为初始化向量解密
	decodeData, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", nil
	}
	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher([]byte(key[0:16]))
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(key[16:]))
	//输出到[]byte数组
	origin_data := make([]byte, len(decodeData))
	blockMode.CryptBlocks(origin_data, decodeData)
	//去除填充,并返回
	return string(unpad(origin_data)), nil
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
