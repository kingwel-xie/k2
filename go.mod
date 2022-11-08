module github.com/kingwel-xie/k2

go 1.15

require (
	github.com/alibaba/sentinel-golang v1.0.4
	github.com/alibaba/sentinel-golang/pkg/adapters/gin v0.0.0-20220815123005-3926bcac42e0
	github.com/alibabacloud-go/darabonba-openapi v0.1.14
	github.com/alibabacloud-go/dysmsapi-20170525/v2 v2.0.9
	github.com/alibabacloud-go/tea v1.1.17
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190307165228-86c17b95fcd5
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/aws/aws-sdk-go-v2 v1.16.15
	github.com/aws/aws-sdk-go-v2/config v1.17.6
	github.com/aws/aws-sdk-go-v2/credentials v1.12.19
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.32
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.10
	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
	github.com/bsm/redislock v0.5.0
	github.com/bytedance/go-tagexpr/v2 v2.7.12
	github.com/casbin/casbin/v2 v2.28.3
	github.com/casbin/gorm-adapter/v3 v3.3.2
	github.com/chanxuehong/wechat v0.0.0-20201110083048-0180211b69fd
	github.com/gin-gonic/gin v1.8.1
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.10.0
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/google/uuid v1.2.0
	github.com/gorilla/websocket v1.4.1
	github.com/json-iterator/go v1.1.12
	github.com/mattn/go-isatty v0.0.14
	github.com/mojocn/base64Captcha v1.3.4
	github.com/mssola/user_agent v0.5.2
	github.com/nsqio/go-nsq v1.0.8
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.11.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/robinjoseph08/redisqueue/v2 v2.1.0
	github.com/shamsher31/goimgext v1.0.0
	github.com/shirou/gopsutil v3.21.5+incompatible
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/slok/go-http-metrics v0.9.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.7.1
	github.com/unrolled/secure v1.0.8
	go.uber.org/multierr v1.5.0
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	gorm.io/driver/mysql v1.0.4-0.20201206014609-ae5fd10184f6
	gorm.io/driver/postgres v1.0.8
	gorm.io/driver/sqlite v1.1.5-0.20201206014648-c84401fbe3ba
	gorm.io/gorm v1.21.11
	gorm.io/plugin/dbresolver v1.1.0
)

//replace (
//)
