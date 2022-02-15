package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type AliyunOSS struct{
	client          *oss.Client
	bucket			*oss.Bucket
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	BucketUrl       string
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

