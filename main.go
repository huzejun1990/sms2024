package main

import (
	"fmt"
	"log"
	"sms2024/sms"
)

func main() {
	s := sms.NewSms()
	req := &sms.SendSmsReq{
		PhoneNumberSet: []string{"13800138000", "13800138001"},
		//TemplateId:       "1244",
		TemplateName:     "零声教育-短信模板1",
		TemplateParamSet: []string{"123456", "5"},
	}
	res, err := s.Send(req)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(res.ToJsonString())

	/*	fmt.Println(config.Conf.GetString("TencentCloudSms.Endpoint"))
		fmt.Println(config.Secret.GetString("TencentCloudApiKey.SecretId"))*/

}
