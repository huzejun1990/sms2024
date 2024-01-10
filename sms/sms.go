// @Author huzejun 2024/1/9 14:42:00
package sms

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"log"
	"sms2024/config"
)

type Sms struct {
	SignName    string
	SmsSdkAppId string
}

var credential *common.Credential

func init() {
	getCredential()
}

func NewSms() *Sms {
	return &Sms{
		SignName:    config.Conf.GetString("TencentCloudSms.DefaultSignName"),
		SmsSdkAppId: config.Conf.GetString("TencentCloudSms.SdkAppId"),
	}
}

func getCredential() {
	credential = common.NewCredential(
		config.Secret.GetString("TencentCloudApiKey.SecretId"),
		config.Secret.GetString("TencentCloudApiKey.SecretKey"),
	)
}

func getClient(region ...string) (*sms.Client, error) {
	defaultRegion := config.Conf.GetString("TencentCloudSms.DefaultRegion")
	if len(region) > 0 {
		defaultRegion = region[0]
	}
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	//cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.HttpProfile.Endpoint = config.Conf.GetString("TencentCloudSms.Endpoint")
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, err := sms.NewClient(credential, defaultRegion, cpf)
	return client, err
}

type SendSmsReq struct {
	PhoneNumberSet   []string
	TemplateId       string
	TemplateParamSet []string
}

type SendSmsRes struct {
	*sms.SendSmsResponse
}

func (s *Sms) Send(req *SendSmsReq) (res *SendSmsRes, err error) {
	client, err := getClient()
	if err != nil {
		log.Println(err)
		return
	}
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(req.PhoneNumberSet)
	request.SmsSdkAppId = common.StringPtr(s.SmsSdkAppId)
	request.SignName = common.StringPtr(s.SignName)
	request.TemplateId = common.StringPtr(req.TemplateId)
	request.TemplateParamSet = common.StringPtrs(req.TemplateParamSet)

	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Println(err)
		return
	}
	res = &SendSmsRes{
		SendSmsResponse: response,
	}
	return res, err
}
