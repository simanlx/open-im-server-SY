package cloud_wallet

import (
	"Open_IM/pkg/cloud_wallet/ncount"
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
	"github.com/spf13/cast"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type WxPayParams struct {
	AppID          string `json:"appid"`            //应用id
	MchID          string `json:"mch_id"`           //商户号
	NonceStr       string `json:"nonce_str"`        //随机字符串
	SignType       string `json:"sign_type"`        //签名类型
	Sign           string `json:"sign"`             //签名
	Body           string `json:"body"`             //商品描述
	OutTradeNo     string `json:"out_trade_no"`     //商户订单号
	TotalFee       int64  `json:"total_fee"`        //总金额、单位分
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

// 微信统一下单支付
func (rpc *CloudWalletServer) WxPayUnifiedOrder(_ context.Context, req *cloud_wallet.WxPayUnifiedOrderReq) (*cloud_wallet.WxPayUnifiedOrderResp, error) {
	resp := &cloud_wallet.WxPayUnifiedOrderResp{CommonResp: &cloud_wallet.CommonResp{ErrCode: 0, ErrMsg: ""}}

	//创建订单

	// 订单号
	orderNo := ncount.GetMerOrderID()

	//微信统一下单接口
	wxPay := NewWxPay()
	wxPayParams := WxPayParams{
		Body:           req.Body,
		OutTradeNo:     orderNo,
		TotalFee:       req.Amount,
		SpbillCreateIp: req.Ip,
		NotifyUrl:      config.Config.WxPay.WxPayNotifyUrl,
	}
	data, err := wxPay.UnifiedOrder(wxPayParams)
	fmt.Println(data, err)

	//map[ prepay_id:wx06171532227961d90acc350936ce910000 result_code:SUCCESS return_code:SUCCESS return_msg:OK sign:A756059360BC9F501E0FA83FB435557A trade_type:APP] <nil>

	return resp, nil
}

func NewWxPay() *WxPay {
	srv := &WxPay{
		SignType: "MD5",
		AppID:    config.Config.WxPay.AppId,
		MchID:    config.Config.WxPay.MchId,
		MchKey:   config.Config.WxPay.MchKey,
		CertData: nil,
	}
	return srv
}

// https need no cert post
func (p *WxPay) postWithoutCert(url string, params WxPayParams) (string, error) {
	// 填充微信支付参数数据
	params = p.fillWxPayParams(params)

	h := &http.Client{}
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

// https need cert post
func (p *WxPay) postWithCert(url string, params WxPayParams) (string, error) {
	// 填充微信支付参数数据
	params = p.fillWxPayParams(params)

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

// wxpayparams转map
func WxPayParamsToMap(wxPayParams WxPayParams) map[string]string {
	return map[string]string{
		"appid":            wxPayParams.AppID,
		"mch_id":           wxPayParams.MchID,
		"nonce_str":        wxPayParams.NonceStr,
		"sign_type":        wxPayParams.SignType,
		"sign":             wxPayParams.Sign,
		"body":             wxPayParams.Body,
		"out_trade_no":     wxPayParams.OutTradeNo,
		"total_fee":        cast.ToString(wxPayParams.TotalFee),
		"spbill_create_ip": wxPayParams.SpbillCreateIp,
		"notify_url":       wxPayParams.NotifyUrl,
		"trade_type":       wxPayParams.TradeType,
	}
}

// 签名
func (p *WxPay) Sign(wxPayParams WxPayParams) string {
	params := WxPayParamsToMap(wxPayParams)

	var keys = make([]string, 0, len(params))
	// 遍历签名参数
	for k := range params {
		if k != "sign" { // 排除sign字段
			keys = append(keys, k)
		}
	}

	// 排序
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
	buf.WriteString(p.MchKey)
	dataMd5 := md5.Sum(buf.Bytes())
	str := hex.EncodeToString(dataMd5[:]) //需转换成切片

	return strings.ToUpper(str)
}

func (p *WxPay) XmlToMap(xmlStr string) map[string]string {
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

// 填充微信支付参数数据
func (p *WxPay) fillWxPayParams(wxPayParams WxPayParams) WxPayParams {
	wxPayParams.AppID = p.AppID
	wxPayParams.MchID = p.MchID
	wxPayParams.SignType = p.SignType
	wxPayParams.TradeType = "APP"
	wxPayParams.NonceStr = nonceStr()

	//签名
	wxPayParams.Sign = p.Sign(wxPayParams)

	return wxPayParams
}

// 统一下单、返回预支付交易会话标识
func (p *WxPay) UnifiedOrder(wxPayParams WxPayParams) (string, error) {
	//统一下单接口
	url := config.Config.WxPay.UnifiedOrder
	xmlStr, err := p.postWithoutCert(url, wxPayParams)
	if err != nil {
		return "", err
	}

	// 解析xml为map数据
	wxPayResp := p.XmlToMap(xmlStr)
	prepayId, ok := wxPayResp["prepay_id"]
	if !ok {
		return "", errors.New("交易失败,no prepay_id ")
	}

	return prepayId, nil
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
