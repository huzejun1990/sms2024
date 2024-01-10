package main

import (
	"fmt"
	"log"
	"sms2024/sms"
)

func main() {
	s := sms.NewSms()
	req := &sms.SendSmsReq{
		PhoneNumberSet:   []string{"13800138000", "13800138001"},
		TemplateId:       "1244",
		TemplateParamSet: []string{"123456", "5"},
	}
	res, err := s.Send(req)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res.ToJsonString())
	/*
		// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
		// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
		// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
		credential := common.NewCredential(
			"AKIDJSN5asxobJySrPM3hAPl4C8GLG3jDUaU",
			"oaSak0SoEwdLIBhLyX1iWBqWDqT0TuX6",
		)
		// 实例化一个client选项，可选的，没有特殊需求可以跳过
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
		// 实例化要请求产品的client对象,clientProfile是可选的
		client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

		// 实例化一个请求对象,每个接口都会对应一个request对象
		request := sms.NewSendSmsRequest()

		request.PhoneNumberSet = common.StringPtrs([]string{ "13800138000", "13800138001" })
		request.SmsSdkAppId = common.StringPtr("1400881659")
		request.SignName = common.StringPtr("签名")
		request.TemplateId = common.StringPtr("1244")
		request.TemplateParamSet = common.StringPtrs([]string{ "123456", "5" })

		// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
		response, err := client.SendSms(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return
		}
		if err != nil {
			panic(err)
		}
		// 输出json格式的字符串回包
		fmt.Printf("%s", response.ToJsonString())
	*/
}
