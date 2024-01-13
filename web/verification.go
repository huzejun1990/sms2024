// @Author huzejun 2024/1/13 20:00:00
package web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"sms2024/config"
	"sms2024/verification"
	"strconv"
	"time"
)

func SendVerification(c *gin.Context) {
	//phoneNumber := c.GetString("phone_number")
	phoneNumber := c.PostForm("phone_number")

	v := validator.New()
	//err := v.Var(phoneNumber, "4164")
	err := v.Var(phoneNumber, "e164")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	sign := config.Conf.GetString("Verification.SignName")
	tempName := config.Conf.GetString("Verification.TempName")
	expireTs := config.Conf.GetInt("Verification.ExpireTs")
	vc := verification.NewVerification()
	res, err := vc.Send(phoneNumber, sign, tempName, func(code string) (params []string) {
		return []string{code, strconv.Itoa(expireTs / 60)}
	}, time.Second*time.Duration(expireTs))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	/*	c.JSON(http.StatusOK,res.ToJsonString())
		fmt.Println(phoneNumber, sign, tempName, expireTs)
	*/
	c.JSON(http.StatusOK, res.Response)
}

func CheckVerification(c *gin.Context) {
	phoneNumber := c.PostForm("phone_number")
	code := c.PostForm("code")
	v := validator.New()
	err := v.Var(phoneNumber, "e164")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = v.Var(code, "number")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	vc := verification.NewVerification()
	res := vc.Check(phoneNumber, code)
	if !res {
		err := errors.New("短信验证不通过！")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "验证通过")
}
