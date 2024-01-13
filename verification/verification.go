// @Author huzejun 2024/1/13 16:11:00
package verification

import (
	"fmt"
	"log"
	"sms2024/sms"
	"time"
)

var verificationPrefix = "verification_"
var verification *Verification

type Verification struct {
	Storage Storage
}

type Storage interface {
	Get(key string) (string, error)
	Set(key string, val string, duration time.Duration) error
	Del(key string) error
}

func NewVerification() *Verification {
	if verification == nil {
		verification = &Verification{
			Storage: NewRedisStorage(),
		}
	}
	return verification
}

func getKey(phoneNumber string) string {
	return verificationPrefix + phoneNumber
}

func (v *Verification) Send(phoneNumber, signName, tempName string, tempParamsFunc func(code string) (params []string), expireTs time.Duration) (*sms.SendSmsRes, error) {
	//生成随机6位验证码
	code := GeneratorRandNo(6)
	//将验证码存入缓存
	key := getKey(phoneNumber)
	//err := v.storage.Set(key, code, expireTs)
	err := v.Storage.Set(key, code, expireTs)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//发送短信
	s := sms.NewSms()
	request := &sms.SendSmsReq{
		PhoneNumberSet:   []string{phoneNumber},
		TemplateName:     tempName,
		TemplateParamSet: tempParamsFunc(code),
	}
	res, err := s.Send(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, err
	//返回
}

func (v *Verification) Check(phoneNumber, code string) bool {
	key := getKey(phoneNumber)
	val, err := v.Storage.Get(key)
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println(val, code)
	if val == code {
		v.Storage.Del(key)
		return true
	}
	return false
}
