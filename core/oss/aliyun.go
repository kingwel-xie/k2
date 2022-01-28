package oss

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"mime/multipart"
	"path/filepath"
	"time"
)

type AliyunOSS struct{
	client          *oss.Client
	bucket			*oss.Bucket
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	BucketUrl		string
}

func NewAliyun(endpoint, accessKeyId, accessKeySecret, bucketName, bucketUrl string) Oss {
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
	return &AliyunOSS{ client: client, bucket: bucket, BucketName: bucketName, BucketUrl: bucketUrl }
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

func (e *AliyunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		log.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename

	// 上传文件流。
	err := e.bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		log.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return e.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (e *AliyunOSS) DeleteFile(key string) error {
	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err := e.bucket.DeleteObject(key)
	if err != nil {
		log.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}

	return nil
}

