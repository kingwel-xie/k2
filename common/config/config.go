package config

import (
	"fmt"
	"github.com/spf13/viper"
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
	Oss   *Oss   `yaml:"oss"`
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
		},
		Extend:    ExtendConfig,
		callbacks: fs,
	}
	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(&_cfg); err != nil {
		fmt.Println(err)
	}

	_cfg.Init()
}
