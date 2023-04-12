package tencent_cloud

import (
	"Open_IM/pkg/common/log"
	"fmt"

	asr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func main() {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		"AKIDkKfk5Fx0bzB8ug8esuiKgnlU3OFrtPIa",
		"EG2Y8TaV0FF0jDSSCnAH9l9KGwTmsiJf",
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "asr.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := asr.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := asr.NewSentenceRecognitionRequest()

	request.ProjectId = common.Uint64Ptr(0)
	request.SubServiceType = common.Uint64Ptr(2)
	request.EngSerViceType = common.StringPtr("16k_zh")
	request.SourceType = common.Uint64Ptr(0)
	request.Url = common.StringPtr("http://office.najieguo.com:10005/openim/1681271806059167000-4548151638488372469.aac")
	request.VoiceFormat = common.StringPtr("aac")
	request.UsrAudioKey = common.StringPtr("dfadf")

	// 返回的resp是一个SentenceRecognitionResponse的实例，与请求对象对应
	response, err := client.SentenceRecognition(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}

const (
	SecretID  = "AKIDkKfk5Fx0bzB8ug8esuiKgnlU3OFrtPIa"
	SecretKey = "EG2Y8TaV0FF0jDSSCnAH9l9KGwTmsiJf"
)

func GetTencentCloudTranslate(url, OperationID string) (*asr.SentenceRecognitionResponse, error) {
	credential := common.NewCredential(
		SecretID,
		SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "asr.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := asr.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := asr.NewSentenceRecognitionRequest()

	request.ProjectId = common.Uint64Ptr(0)
	request.SubServiceType = common.Uint64Ptr(2)
	request.EngSerViceType = common.StringPtr("16k_zh")
	request.SourceType = common.Uint64Ptr(0)
	// http://office.najieguo.com:10005/openim/1681271806059167000-4548151638488372469.aac
	request.Url = common.StringPtr(url)
	request.VoiceFormat = common.StringPtr("aac")
	request.UsrAudioKey = common.StringPtr("dfadf")

	// 返回的resp是一个SentenceRecognitionResponse的实例，与请求对象对应
	response, err := client.SentenceRecognition(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Error(OperationID, "An API error has returned: %s", err)
		return nil, err
	}
	if err != nil {
		log.Error(OperationID, "An API error has returned: %s", err)
		return nil, err
	}
	// 输出json格式的字符串回包
	return response, nil
}
