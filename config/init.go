package config

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/niubir/utils"
	yaml "gopkg.in/yaml.v2"
)

var Config config

// Init 初始化配置文件
// configFilePath->env->./config/config.yml
func Init() error {
	configFilePath := "./config/config.yml"
	if envConfigFilePath := os.Getenv(`CONFIG_FILE`); envConfigFilePath != "" {
		configFilePath = envConfigFilePath
	}

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return errors.New("invalid load config file: " + err.Error())
	}
	if err := yaml.Unmarshal(data, &Config); err != nil {
		return errors.New("unmarshal config file: " + err.Error())
	}
	if err := Config.init(); err != nil {
		return errors.New("failed to init config: " + err.Error())
	}

	configBody, _ := yaml.Marshal(Config)
	utils.BoldPrint("***Config Information***\n\n%s", string(configBody))

	return nil
}
