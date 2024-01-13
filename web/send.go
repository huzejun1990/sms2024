// @Author huzejun 2024/1/13 13:28:00
package web

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	"log"
	"net/http"
	"sms2024/sms"
)

type SendReq struct {
	Template       string `form:"template" binding:"required"`
	TemplateParams string `form:"template_params" binding:"json"`
	PhoneList      string `form:"phone_list" binding:"required,json"`
}

func Send(c *gin.Context) {
	req := &SendReq{}
	err := c.ShouldBind(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var templateParams []string
	err = json.Unmarshal([]byte(req.TemplateParams), &templateParams)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var phoneList []string
	err = json.Unmarshal([]byte(req.PhoneList), &phoneList)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	v := validator.New()
	err = v.Var(phoneList, "gte=1,dive,required,e164")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	//fmt.Println(req.Template,phoneList,templateParams)

	s := sms.NewSms()
	request := &sms.SendSmsReq{
		PhoneNumberSet:   phoneList,
		TemplateName:     req.Template,
		TemplateParamSet: templateParams,
	}
	res, err := s.Send(request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res.Response)
}
