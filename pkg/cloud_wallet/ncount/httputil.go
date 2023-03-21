package ncount

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpPost(url string, body []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 1. 用新账通平台公钥对字符串进行非对称加密；
// 2. 对加密后的二进制转 Base64 编码
func Encrpt(message []byte, publicKeyString string) ([]byte, error) {
	if publicKeyString == "" {
		return nil, fmt.Errorf("publicKeyString is empty")
	}
	publicKeyBlock, _ := pem.Decode([]byte(publicKeyString))
	if publicKeyBlock == nil {
		return nil, fmt.Errorf("publicKeyString is invalid")
	}
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey.(*rsa.PublicKey),
		message,
		nil,
	)
	if err != nil {
		return nil, err
	}
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return []byte(encoded), nil
}
