package config

import (
	"fmt"
	"time"

	"github.com/donetkit/contrib/server"
	"github.com/donetkit/contrib/utils/config"
)

type serverConfig struct {
	App        appConfig        `yaml:"app"`
	Redis      redisConfig      `yaml:"redis"`
	OpenWechat openWechatConfig `yaml:"open_wechat"`
}

type appConfig struct {
	ReadTimeOut   time.Duration `yaml:"read_time_out"`
	WriterTimeOut time.Duration `yaml:"writer_time_out"`
	Host          string        `yaml:"host"`
	Port          int           `yaml:"port"`
	RunMode       string        `yaml:"run_mode"`
	LoggerLevel   int           `yaml:"logger_level"` // 日志等级
	ApiVersion    string        `yaml:"api_version"`  // 版本号
	Url           string        `yaml:"url"`
}

type redisConfig struct {
	Addr     string `yaml:"addr"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type openWechatConfig struct {
	AppId          string `yaml:"app_id"`
	AppSecret      string `yaml:"app_secret"`
	Url            string `yaml:"url"`
	Token          string `yaml:"token"`
	EncodingAESKey string `yaml:"encoding_aes_key"`
	CheckPublish   bool   `yaml:"check_publish"`
}

var ymlConfig = new(config.YMLConfig)

var ServerConfig = new(serverConfig)

func init() {
	ymlConfig.LoadFullPath(fmt.Sprintf("config/config_%s.yml", server.EnvName), ServerConfig)
}
