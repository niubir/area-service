package config

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/niubir/utils"
	yaml "gopkg.in/yaml.v2"
)

var (
	configFilePath = "/config/config.yml"
	Config         config
)

func init() {
	if envConfigFilePath := os.Getenv(`CONFIG_FILE_PATH`); envConfigFilePath != "" {
		configFilePath = envConfigFilePath
	}
}

// Init 初始化配置文件
func Init() error {
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
