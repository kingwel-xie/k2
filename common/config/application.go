package config

type Application struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          int64
	Name          string
	JwtSecret     string
	Mode          string
	DemoMsg       string
}

var ApplicationConfig = new(Application)
