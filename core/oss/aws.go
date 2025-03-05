package oss

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"os"
	"path"
	"strconv"
	"time"
)

type S3 struct {
	client     *s3.Client
	uploader   *manager.Uploader
	downloader *manager.Downloader
	//Region        	string
	//AccessKeyId     string
	//AccessKeySecret string
	BucketName string
	BucketUrl  string
}

func (e *S3) Name() string {
	return "aws"
}

func (e *S3) UpLoadLocalFile(objectName string, localFile string) error {
	file, err := os.Open(localFile)
	if err != nil {
		return err
	}
	_, err = e.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &e.BucketName,
		Key:    &objectName,
		Body:   file,
	})
	return err
}

func (e *S3) UploadFile(file io.Reader, filename string) (string, error) {
	_, err := e.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &e.BucketName,
		Key:    &filename,
		Body:   file,
	})
	return "", err
}

func (e *S3) IsFileExists(filename string) (bool, error) {
	_, err := e.client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(e.BucketName),
		Key:    aws.String(filename),
	})
	if err != nil {
		return false, err // 错误
	}
	return true, nil // 文件存在
}

func (e *S3) SignTemporaryExternalUrl(filename string, expiredInSec int64) (string, error) {
	presignClient := s3.NewPresignClient(e.client)
	presignResult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(e.BucketName),
		Key:    aws.String(filename),
	}, s3.WithPresignExpires(time.Duration(expiredInSec)*time.Second))
	if err != nil {
		return "", fmt.Errorf("couldn't get presigned URL for GetObject")
	}
	return presignResult.URL, nil
}

func (e *S3) DownloadFile(filename string) (io.ReadCloser, error) {
	w, err := os.CreateTemp("", "s3-*.tmp")
	if err != nil {
		return nil, err
	}
	_, err = e.downloader.Download(context.TODO(), w, &s3.GetObjectInput{
		Bucket: &e.BucketName,
		Key:    &filename,
	})
	return w, err
}

func (e *S3) GetFileMeta(filename string) (map[string][]string, error) {
	resp, err := e.client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: &e.BucketName,
		Key:    &filename,
	})
	if err != nil {
		return nil, err
	}
	metadata := map[string][]string{
		"Content-Type":   {aws.ToString(resp.ContentType)},
		"Content-Length": {strconv.FormatInt(*resp.ContentLength, 10)},
		"Last-Modified":  {resp.LastModified.Format(time.RFC3339)},
		"ETag":           {aws.ToString(resp.ETag)},
	}

	// 合并用户自定义元数据
	for k, v := range resp.Metadata {
		metadata[k] = []string{v}
	}
	return metadata, nil
}

func (e *S3) DeleteFile(filename string) error {
	_, err := e.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(e.BucketName),
		Key:    aws.String(filename),
	})
	return err
}

func NewS3(region, accessKeyId, accessKeySecret, bucketName string, bucketUrl string) *S3 {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: accessKeyId, SecretAccessKey: accessKeySecret, SessionToken: "",
				Source: "example hard coded credentials",
			},
		}))
	if err != nil {
		panic("Failed to load configuration")
	}
	cfg.Region = region

	ss := s3.NewFromConfig(cfg)
	return &S3{
		client:     ss,
		uploader:   manager.NewUploader(ss),
		downloader: manager.NewDownloader(ss),
		//Region: region,
		BucketName: bucketName,
		BucketUrl:  bucketUrl,
	}
}

func (e *S3) GeneratePresignedToken(directory string, filename string, exp int64) (interface{}, error) {
	presignClient := s3.NewPresignClient(e.client)

	presignResult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(e.BucketName),
		Key:    aws.String(path.Join(directory, filename)),
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't get presigned URL for GetObject")
	}
	return presignResult, nil
}
