package user

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
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
	//签名
	sign := Sign(fmt.Sprintf("appId%stoken%s", AppId, token))

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
	fmt.Println("----", reply, err)

	//if reply.Code != "200000" {
	//	return "", errors.New(fmt.Sprintf("接口响应错误:%s", reply.Message))
	//}

	mobileName := "1F881288CC68352FC410E8D4A36FC6E0"
	mobile := Decrypt(mobileName)
	//mobile := Decrypt(reply.Data.MobileName)
	if err != nil {
		return "", errors.New("解密手机号码失败")
	}

	return mobile, nil
}

// 签名
func Sign(text string) string {
	hMac := hmac.New(sha256.New, []byte(AppKey))
	hMac.Write([]byte(text))
	signature := hex.EncodeToString(hMac.Sum(nil))
	return string(signature)
}

func Decrypt(mobileName string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover error:", r)
			return
		}
	}()

	hash := md5.Sum([]byte(AppKey))
	hashString := hex.EncodeToString(hash[:])

	block, _ := aes.NewCipher([]byte(hashString[:16]))
	ecb := cipher.NewCBCDecrypter(block, []byte(hashString[16:]))
	source, _ := hex.DecodeString(mobileName)
	decrypted := make([]byte, len(source))
	ecb.CryptBlocks(decrypted, source)
	return string(unpad(decrypted))
}

func decrptPhone(data string, key string) string {
	hash := md5.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])
	block, _ := aes.NewCipher([]byte(hashString[:16]))
	ecb := cipher.NewCBCDecrypter(block, []byte(hashString[16:]))
	source, _ := hex.DecodeString(data)
	decrypted := make([]byte, len(source))
	ecb.CryptBlocks(decrypted, source)
	return string(unpad(decrypted))
}

func PKCS5Unpadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])

	fmt.Println("--encrypt--", ciphertext, length, unpadding, ciphertext[length-1])

	return ciphertext[:(length - unpadding)]
}
