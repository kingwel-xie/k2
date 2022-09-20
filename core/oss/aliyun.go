package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"hash"
	"io"
	"time"
)

type AliyunOSS struct{
	client          *oss.Client
	bucket			*oss.Bucket
	BucketName      string
	BucketUrl       string
}

func (e *AliyunOSS) Name() string {
	return "aliyun"
}

func NewAliyun(endpoint, accessKeyId, accessKeySecret, bucketName string, bucketUrl string) Oss {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Warnf("failed to setup Aliyun: %v", err)
		panic(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		panic(err)
	}
	return &AliyunOSS{ client: client, bucket: bucket, BucketName: bucketName, BucketUrl: bucketUrl}
}

func (e *AliyunOSS) GetFileMeta(filename string) (map[string][]string, error) {
	return e.bucket.GetObjectDetailedMeta(filename)
}

// UpLoadLocalFile 文件上传
func (e *AliyunOSS) UpLoadLocalFile(yourObjectName string, localFile string) error {
	// 设置分片大小为100 KB，指定分片上传并发数为3，并开启断点续传上传。
	// 其中<yourObjectName>与objectKey是同一概念，表示断点续传上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// "LocalFile"为filePath，100*1024为partSize。
	err := e.bucket.UploadFile(yourObjectName, localFile, 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	if err != nil {
		log.Warnf("failed to upload to Aliyun: %v", err)
		return err
	}
	return nil
}

func (e *AliyunOSS) UploadFile(file io.Reader, filename string) (string, error) {
	// 上传文件流。
	err := e.bucket.PutObject(filename, file)
	if err != nil {
		log.Errorf("PutObject Failed, %v", err)
		return "", err
	}
	return e.BucketUrl + "/" + filename, nil
}

func (e *AliyunOSS) DownloadFile(filename string) (io.ReadCloser, error) {
	// 上传文件流。
	reader, err := e.bucket.GetObject(filename)
	if err != nil {
		log.Errorf("GetObject Failed, %v", err)
		return nil, err
	}
	return reader, nil
}

func (e *AliyunOSS) DeleteFile(filename string) error {
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err := e.bucket.DeleteObject(filename)
	if err != nil {
		log.Errorf("function bucketManager.Delete() Filed, %v", err)
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// Post Policy
func get_gmt_iso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

type ConfigStruct struct{
	Expiration string `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct{
	AccessKeyId string `json:"accessid"`
	Host string `json:"host"`
	Expire int64 `json:"expire"`
	Signature string `json:"signature"`
	Policy string `json:"policy"`
	Directory string `json:"dir"`
	Callback string `json:"callback"`
}

func (e *AliyunOSS) GeneratePresignedToken(uploadDir, filename string, expireSeconds int64) (interface{}, error) {
	now := time.Now().Unix()
	if expireSeconds == 0 {
		expireSeconds = 60
	}
	expireEnd := now + expireSeconds
	var tokenExpire = get_gmt_iso8601(expireEnd)

	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, uploadDir)
	config.Conditions = append(config.Conditions, condition)

	// calculate signature
	result, _ := json.Marshal(config)
	bytes := base64.StdEncoding.EncodeToString(result)

	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(e.client.Config.AccessKeySecret))
	io.WriteString(h, bytes)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var policyToken PolicyToken
	policyToken.AccessKeyId = e.client.Config.AccessKeyID
	policyToken.Host = fmt.Sprintf("https://%s.%s", e.BucketName, e.client.Config.Endpoint)
	policyToken.Expire = expireEnd
	policyToken.Signature = signedStr
	policyToken.Directory = uploadDir
	policyToken.Policy = bytes
	policyToken.Callback = ""

	return &policyToken, nil
}

