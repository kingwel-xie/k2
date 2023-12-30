package utils

import (
	"fmt"
	"github.com/kingwel-xie/k2/common/config"
	"testing"
)

func TestOssUpload(t *testing.T) {

	config := &config.Oss {
		Which: "aliyun",
		Aliyun: config.AliyunOSS {
			Endpoint: "oss-cn-hongkong.aliyuncs.com",
			AccessKeyId: "LTAI5tPeVhT5ABqWRzAJg6EQ",
			AccessKeySecret: "S32H9EOm1V2HuLLEFWSo9eqxnGkY0d",
			BucketName: "kobh-test",
		},
	}
	oss := config.Setup()

	err := oss.UpLoadLocalFile("testttt", "d:/tt.dat")

	fmt.Println(err)
}