package config

import (
	"github.com/kingwel-xie/k2/core/oss"
)

type Oss struct {
	Which string 			`yaml:"which"`
	Local Local  			`yaml:"local"`
	Qiniu      Qiniu      `yaml:"qiniu"`
	Aliyun  AliyunOSS  `yaml:"aliyun"`
	Tencent TencentCOS `yaml:"tencent"`
}

var OssConfig = new(Oss)


type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
}

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                // 存储区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                          // 空间名称
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"`                     // CDN加速域名
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`                  // 是否使用https
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`               // 秘钥AK
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`               // 秘钥SK
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"` // 上传是否使用CDN上传加速
}

type AliyunOSS struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
	BucketName      string `mapstructure:"bucket-name" json:"bucketName" yaml:"bucket-name"`
	BucketUrl       string `mapstructure:"bucket-url" json:"bucketUrl" yaml:"bucket-url"`
}

type TencentCOS struct {
	Bucket     string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	SecretID   string `mapstructure:"secret-id" json:"secretID" yaml:"secret-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	BaseURL    string `mapstructure:"base-url" json:"baseURL" yaml:"base-url"`
	PathPrefix string `mapstructure:"path-prefix" json:"pathPrefix" yaml:"path-prefix"`
}

func (e Oss) Setup() oss.Oss {
	switch e.Which {
	case "local":
		return oss.NewLocal(e.Local.Path)
	//case "qiniu":
	//	oss = &Qiniu{}
	//case "tencent":
	//	return oss.NewTencent(e.Tencent.Region, e.Tencent.Bucket, e.Tencent.BaseURL, e.Tencent.PathPrefix...)
	case "aliyun":
		return oss.NewAliyun(e.Aliyun.Endpoint, e.Aliyun.AccessKeyId, e.Aliyun.AccessKeySecret, e.Aliyun.BucketName, e.Aliyun.BucketUrl)
	default:
		return oss.NewLocal(e.Local.Path)
	}
}