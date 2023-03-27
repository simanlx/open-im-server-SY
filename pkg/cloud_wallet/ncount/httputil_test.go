package ncount

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestEncrpt(t *testing.T) {
	testdata := &NewAccountMsgCipher{
		MerUserId: "steven_test",
		Mobile:    "15282603386",
		UserName:  "沈晨曦",
		CertNo:    "511623199808185554",
	}
	// 1. 将报文信息转换为 JSON 格式
	data, err := json.Marshal(testdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2. 将 JSON 格式的报文信息用平台公钥 RSA 加密后 base64 的编码值
	cipher, err := Encrpt(data, PUBLIC_KEY)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 3.version=[]tranCode=[]merId=[]merOrderId=[]submitTime=[]msgCiphertext=[]signType=[]
	// signValue= version
	// 2. 使用RSA进行私钥签名
	// 3. 签名后的二进制转Base64编码

	body := &NewAccountBaseParam{
		Version:    "1.0",
		TranCode:   "R010",
		MerId:      MER_USER_ID,
		MerOrderId: "10086",
		//格 式 ： YYYYMMDDHHMMSS
		SubmitTime: fmt.Sprintf("%d", time.Now().Unix()),
		MsgCiphert: string(cipher),
		SignType:   "1",
		Charset:    "1",
	}
	err, str := body.flushSignValue()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 4. 使用RSA进行私钥签名
	sign, err := Sign([]byte(str), PRIVATE_KEY)
	if err != nil {
		fmt.Println(err)
		return
	}
	body.SignValue = sign
	content, err := json.Marshal(body)

	result, err := httpPost(NewAccountURL, content)
	if err != nil {
		panic(err)
	}
	reply := &NewAccountReturnFromPlatform{}
	err = json.Unmarshal(result, reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}

func TestMain1(t *testing.T) {
	// 加密数据
	// 签名
	plainData := "Hello, world!"
	c, err := Sign([]byte(plainData), PRIVATE_KEY)
	fmt.Println("sign:", c, err)
}

func TestMain2(t *testing.T) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(PRIVATE_KEY)
	if err != nil {
		panic(err)
	}

	// 将私钥解析为 *rsa.PrivateKey 对象
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
	if err != nil {
		panic(err)
	}

	hashed := sha1.Sum([]byte("Hello, world!"))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hashed[:])
	if err != nil {
		panic(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(signature))
}
