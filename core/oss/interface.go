package oss

import (
	"io"
	"mime/multipart"

	"github.com/kingwel-xie/k2/core/logger"
)

var log = logger.Logger("oss")


type Oss interface {
	UpLoadLocalFile(objectName string, localFile string) error
	UploadFile(file *multipart.FileHeader, filename string) error
	DownloadFile(filename string) (io.ReadCloser, error)
	DeleteFile(filename string) error
}
