package oss

import (
	"io"

	"github.com/kingwel-xie/k2/core/logger"
)

var log = logger.Logger("oss")


type Oss interface {
	UpLoadLocalFile(objectName string, localFile string) error
	UploadFile(file io.Reader, filename string) (string, error)
	DownloadFile(filename string) (io.ReadCloser, error)
	GetFileMeta(filename string) (map[string][]string, error)
	DeleteFile(filename string) error
}
