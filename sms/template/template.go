// @Author huzejun 2024/1/10 20:02:00
package template

import (
	"errors"
	"log"
	"sync"
)

var data *tempData

type tempData struct {
	sync.RWMutex
	templates map[string]*Template
}

type Template struct {
	ID          string
	TempName    string
	ParamsCount int
}

func init() {
	data = &tempData{
		templates: make(map[string]*Template),
	}

	temp1 := &Template{
		ID:          "13333168",
		TempName:    "零声教育-短信模板1",
		ParamsCount: 2,
	}
	temp2 := &Template{
		ID:          "683784",
		TempName:    "TCP/IP",
		ParamsCount: 0,
	}
	Register(temp1.TempName, temp1)
	Register(temp2.TempName, temp2)
}

func Register(key string, temp *Template) error {
	if temp == nil {
		err := errors.New("短信模板对象不能为空")
		log.Println(err)
		return err
	}
	data.Lock()
	defer data.Unlock()
	if _, ok := data.templates[key]; ok {
		err := errors.New("同一个短信模板不能添加多次")
		log.Println(err)
		return err
	}
	data.templates[key] = temp
	return nil
}
func GetTemplate(key string) (*Template, error) {
	data.RLock()
	defer data.RUnlock()
	t, ok := data.templates[key]
	if ok {
		return t, nil
	} else {
		err := errors.New("模板不存在")
		return nil, err
	}
}
