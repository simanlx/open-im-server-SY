package cloud_wallet

import (
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/proto/cloud_wallet"
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	Fail            = "FAIL"
	Success         = "SUCCESS"
	UnifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
)

type WxPayParams struct {
	AppID          string `json:"appid"`            //应用id
	MchID          string `json:"mch_id"`           //商户号
	NonceStr       string `json:"nonce_str"`        //随机字符串
	SignType       string `json:"sign_type"`        //签名类型
	Sign           string `json:"sign"`             //签名
	Body           string `json:"body"`             //商品描述
	OutTradeNo     string `json:"out_trade_no"`     //商户订单号
	TotalFee       int    `json:"total_fee"`        //总金额、单位分
	SpbillCreateIp string `json:"spbill_create_ip"` //终端ip
	NotifyUrl      string `json:"notify_url"`       //通知地址
	TradeType      string `json:"trade_type"`       //交易类型
}

type WxPay struct {
	SignType string // 签名类型
	AppID    string //appId
	MchID    string //mchId
	MchKey   string //mchKey
	CertData []byte //证书
}

// 获取云账户信息
func (rpc *CloudWalletServer) UserNcountAccount22(_ context.Context, req *cloud_wallet.UserNcountAccountReq) (*cloud_wallet.UserNcountAccountResp, error) {
	resp := &cloud_wallet.UserNcountAccountResp{Step: 0, BalAmount: "0", CommonResp: &cloud_wallet.CommonResp{ErrCode: 0, ErrMsg: ""}}

	wxPay := NewWxPay()
	wxPayParams := WxPayParams{
		Body:           "测试订单",
		OutTradeNo:     "13568323788",
		TotalFee:       1,
		SpbillCreateIp: "127.0.0.1",
		NotifyUrl:      "http://server.jiadengni.com:10002/cloudWallet/charge_account_callback",
		TradeType:      "APP",
	}
	data, err := wxPay.UnifiedOrder(wxPayParams)
	fmt.Println(data, err)

	//map[ prepay_id:wx06171532227961d90acc350936ce910000 result_code:SUCCESS return_code:SUCCESS return_msg:OK sign:A756059360BC9F501E0FA83FB435557A trade_type:APP] <nil>

	return resp, nil
}

func NewWxPay() *WxPay {
	srv := &WxPay{
		SignType: "MD5",
		AppID:    "wx93c8d6511c240c66",
		MchID:    "1616233949",
		MchKey:   "VCCaXCPTcaeiadxyXioCYyAqo7xlARFO",
		CertData: nil,
	}

	return srv
}

// 读取证书
func (p *WxPay) WriteCertData() error {
	certPath := config.Config.WxPay.CertPath
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return errors.New("读取证书失败")
	}
	p.CertData = certData
	return nil
}

// https need no cert post
func (p *WxPay) postWithoutCert(url string, params WxPayParams) (string, error) {
	h := &http.Client{}
	response, err := h.Post(url, "application/xml; charset=utf-8", strings.NewReader(ToXml(params)))
	if err != nil {
		return "", err
	}

	fmt.Println("--postWithoutCert--", ToXml(params))

	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// https need cert post
func (p *WxPay) postWithCert(url string, params WxPayParams) (string, error) {
	if p.CertData == nil {
		return "", errors.New("证书数据为空")
	}

	// 将pkcs12证书转成pem
	cert := pkcs12ToPem(p.CertData, p.MchID)

	conf := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	transport := &http.Transport{
		TLSClientConfig:    conf,
		DisableCompression: true,
	}

	h := &http.Client{Transport: transport}
	response, err := h.Post(url, "application/xml; charset=utf-8", strings.NewReader(ToXml(params)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// wxpay参数转map
func WxPayParamsToMap(wxPayParams WxPayParams) map[string]string {
	return map[string]string{
		"appid":            wxPayParams.AppID,
		"mch_id":           wxPayParams.MchID,
		"nonce_str":        wxPayParams.NonceStr,
		"sign_type":        wxPayParams.SignType,
		"sign":             wxPayParams.Sign,
		"body":             wxPayParams.Body,
		"out_trade_no":     wxPayParams.OutTradeNo,
		"total_fee":        strconv.Itoa(wxPayParams.TotalFee),
		"spbill_create_ip": wxPayParams.SpbillCreateIp,
		"notify_url":       wxPayParams.NotifyUrl,
		"trade_type":       wxPayParams.TradeType,
	}
}

func ToXml(wxPayParams WxPayParams) string {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range WxPayParamsToMap(wxPayParams) {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(v)
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</xml>`)

	return buf.String()
}

// 用时间戳生成随机字符串
func nonceStr() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}

// 将Pkcs12转成Pem
func pkcs12ToPem(p12 []byte, password string) tls.Certificate {
	blocks, err := pkcs12.ToPEM(p12, password)

	defer func() {
		if x := recover(); x != nil {
			log.Print(x)
		}
	}()

	if err != nil {
		panic(err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}
	return cert
}

// 签名
func (c *WxPay) Sign(wxPayParams WxPayParams) string {
	params := WxPayParamsToMap(wxPayParams)

	fmt.Println("--Sign---", params)

	// 创建切片
	var keys = make([]string, 0, len(params))
	// 遍历签名参数
	for k := range params {
		if k != "sign" { // 排除sign字段
			keys = append(keys, k)
		}
	}

	// 由于切片的元素顺序是不固定，所以这里强制给切片元素加个顺序
	sort.Strings(keys)

	//创建字符缓冲
	var buf bytes.Buffer
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(`=`)
		buf.WriteString(params[k])
		buf.WriteString(`&`)
	}
	// 加入apiKey作加密密钥
	buf.WriteString(`key=`)
	buf.WriteString(c.MchKey)

	fmt.Println("--Sign-string--", string(buf.Bytes()))

	dataMd5 := md5.Sum(buf.Bytes())
	str := hex.EncodeToString(dataMd5[:]) //需转换成切片

	return strings.ToUpper(str)
}

func XmlToMap(xmlStr string) map[string]string {

	params := make(map[string]string)
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))

	var (
		key   string
		value string
	)

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement: // 开始标签
			key = token.Name.Local
		case xml.CharData: // 标签内容
			content := string([]byte(token))
			value = content
		}
		if key != "xml" {
			if value != "\n" {
				params[key] = value
			}
		}
	}

	return params
}

// 处理 HTTPS API返回数据，转换成Map对象。return_code为SUCCESS时，验证签名。
func (c *WxPay) processResponseXml(xmlStr string) (map[string]string, error) {
	params := XmlToMap(xmlStr)
	returnCode, ok := params["return_code"]
	if !ok {
		return nil, errors.New("no return_code in XML")
	}

	if returnCode == Fail {
		return params, nil
	} else if returnCode == Success {
		if c.ValidSign(params) {
			return params, nil
		} else {
			return nil, errors.New("invalid sign value in XML")
		}
	} else {
		return nil, errors.New("return_code value is invalid in XML")
	}
}

// 验证签名
func (c *WxPay) ValidSign(params map[string]string) bool {
	return true
	//if !params.ContainsKey(Sign) {
	//	return false
	//}
	//return params.GetString(Sign) == c.Sign(params)
}

// 统一下单
func (c *WxPay) UnifiedOrder(wxPayParams WxPayParams) (map[string]string, error) {

	//随机字符串
	wxPayParams.NonceStr = nonceStr()

	wxPayParams.AppID = c.AppID

	wxPayParams.MchID = c.MchID

	wxPayParams.SignType = c.SignType

	fmt.Println("--wxPayParams--", wxPayParams)

	//参数签名
	wxPayParams.Sign = c.Sign(wxPayParams)

	//统一下单接口
	url := UnifiedOrderUrl
	xmlStr, err := c.postWithoutCert(url, wxPayParams)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(xmlStr)
}
