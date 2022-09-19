package oss

import (
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"

	"github.com/kingwel-xie/k2/core/utils"
)

type Local struct{
	Path string
}

func (l *Local) PresignToken(c *gin.Context) {
	c.AbortWithStatus(400)
}

func NewLocal(path string) Oss {
	_ = utils.IsNotExistMkDir(path)
	return &Local{ Path: path }
}

func (l *Local) UpLoadLocalFile(objectName string, localFile string) error {
	if l.Path != "" {
		source, err := os.Open(localFile)
		if err != nil {
			return err
		}
		defer source.Close()

		destination, err := os.Create(path.Join(l.Path, objectName))
		if err != nil {
			return err
		}

		defer destination.Close()
		_, err = io.Copy(destination, source)
		return err
	}
	return nil
}

func (l *Local) UploadFile(file io.Reader, filename string) (string, error) {
	// 拼接路径和文件名
	p := path.Join(l.Path, filename)

	err := os.MkdirAll(path.Dir(p), os.ModePerm)
	if err != nil {
		log.Errorf("function os.MkDirAll() failed for %s, %v", path.Dir(p), err)
		return "", err
	}

	out, createErr := os.Create(p)
	if createErr != nil {
		log.Errorf("function os.Create() Filed, %v", createErr)
		return "", createErr
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, file) // 传输（拷贝）文件
	if copyErr != nil {
		log.Errorf("function io.Copy() Filed, %v", copyErr)
		return "", copyErr
	}
	return p, nil
}

func (l *Local) DownloadFile(filename string) (io.ReadCloser, error) {
	p := path.Join(l.Path, filename)

	f, openError := os.Open(p) // 读取文件
	if openError != nil {
		log.Errorf("function file.Open() Filed, %v", openError)
		return nil, openError
	}
	return f, nil
}

func (l *Local) DeleteFile(filename string) error {
	p := path.Join(l.Path, filename)
	if err := os.Remove(p); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}
	return nil
}

func (l *Local) GetFileMeta(filename string) (map[string][]string, error) {
	var header = make(map[string][]string)
	return header, nil
}
