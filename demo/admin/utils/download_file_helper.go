package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/middleware"
	"io"
	"net/http"
)

func ConcatDownloadPrefix(c *gin.Context) string {
	scheme := c.GetHeader("X-Forwarded-Proto")
	if len(scheme) == 0 {
		scheme = "http"
	}
	return fmt.Sprintf("%s://%s%s", scheme, c.Request.Host, middleware.DownloadUrlPrefix)
}

type ReaderOpener func() (io.ReadCloser, error)

func Base64Opener(content string) ReaderOpener {
	return func() (io.ReadCloser, error) {
		data, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			return nil, err
		}
		return io.NopCloser(bytes.NewReader(data)), nil
	}
}

func UrlOpener(url string) ReaderOpener {
	return func() (io.ReadCloser, error) {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		return res.Body, nil
	}
}

func UploadFile(prefix string, filename string, readerOpener ReaderOpener) (string, error) {

	reader, err := readerOpener()
	if err != nil {
		return "", err
	}
	defer func() {
		_ = reader.Close()
	}()

	oss := common.Runtime.GetOss()
	_, err = oss.UploadFile(reader, filename)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", prefix, filename), nil
}
