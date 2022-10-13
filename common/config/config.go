package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

var (
	ExtendConfig interface{}
	_cfg         *Settings
)

// Settings 兼容原先的配置结构
type Settings struct {
	Settings  Config      `yaml:"settings"`
	Extend    interface{} `yaml:"extend"`
	callbacks []func()
}

func (e *Settings) runCallback() {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

func (e *Settings) OnChange() {
	e.init()
	fmt.Println("!!! config change and reload")
}

func (e *Settings) Init() {
	e.init()
	fmt.Println("!!! config init")
}

func (e *Settings) init() {
	e.Settings.Logger.Setup()
	e.runCallback()
}

// Config 配置集合
type Config struct {
	Application *Application `yaml:"application"`
	Ssl         *Ssl         `yaml:"ssl"`
	Logger      *Logger      `yaml:"logger"`
	Jwt         *Jwt         `yaml:"jwt"`
	Database    *Database    `yaml:"database"`
	Cache       *Cache       `yaml:"cache"`
	Queue       *Queue       `yaml:"queue"`
	Locker      *Locker      `yaml:"locker"`
	Oss   		*Oss   		 `yaml:"oss"`
	Sms   		*Sms   		 `yaml:"sms"`
	File   		*File   	 `yaml:"file"`
}

// Setup 载入配置文件
func Setup(configFile string,
	fs ...func()) {
	_cfg = &Settings{
		Settings: Config{
			Application: ApplicationConfig,
			Ssl:         SslConfig,
			Logger:      LoggerConfig,
			Jwt:         JwtConfig,
			Database:    DatabaseConfig,
			Cache:       CacheConfig,
			Queue:       QueueConfig,
			Locker:      LockerConfig,
			Oss:   		 OssConfig,
			Sms:   		 SmsConfig,
			File:        FileConfig,
		},
		Extend:    ExtendConfig,
		callbacks: fs,
	}

	v := viper.New()
	//自动获取全部的env加入到viper中。默认别名和环境变量名一致。（如果环境变量多就全部加进来）
	v.AutomaticEnv()

	//替换读取格式。默认a.b.c.d格式读取env，改为a_b_c_d格式读取
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 本地配置文件位置
	v.SetConfigFile(configFile)

	//支持 "yaml", "yml", "json", "toml", "hcl", "tfvars", "ini", "properties", "props", "prop", "dotenv", "env":
	v.SetConfigType("yaml")

	//读配置文件到viper配置中
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 系列化成config对象
	if err = v.Unmarshal(&_cfg); err != nil {
		panic(err)
	}

	_cfg.Init()
}
