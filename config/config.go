// @Author huzejun 2024/1/10 17:34:00
package config

import (
	"github.com/spf13/viper"
)

type config struct {
	viper *viper.Viper
}

var Conf *config
var Secret *config

func init() {
	Conf = &config{
		viper: getConf("conf", "config/conf"),
	}

	Secret = &config{
		viper: getConf("secret", "config/secret"),
	}
}

func getConf(configName, configPath string) *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(configName)
	v.AddConfigPath(configPath)
	v.ReadInConfig()
	/*	err := v.ReadInConfig()
		if err != nil {
			fmt.Println(err)
		}*/
	return v
}

func (c *config) GetString(key string) string {
	//return c.GetString(key) //循环调用
	return c.viper.GetString(key)

}
