package oss

import (
	"errors"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct{
	Path string
}

func NewLocal(path string) Oss {
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

func (l *Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	//name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(l.Path, os.ModePerm)
	if mkdirErr != nil {
		log.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := l.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		log.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		log.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		log.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

func (l *Local) DeleteFile(key string) error {
	p := path.Join(l.Path, key)
	if err := os.Remove(p); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}
	return nil
}
