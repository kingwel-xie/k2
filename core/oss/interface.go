package oss

import (
	"github.com/kingwel-xie/k2/core/logger"
	"mime/multipart"
)

var log = logger.Logger("oss")


type Oss interface {
	UpLoadLocalFile(objectName string, localFile string) error
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}
