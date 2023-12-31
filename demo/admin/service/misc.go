package service

import (
	"admin/service/dto"
	"bytes"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/service"
	"os"
	"sync"
	"time"
)

var lock sync.RWMutex = sync.RWMutex{}
var uuidMap = map[string]dataDescriptor{}

type TbxMisc struct {
	service.Service
}

type dataDescriptor struct {
	ContentType string	// 'application/pdf' ...
	Data *bytes.Buffer	// buffer to the data
	Uuid string			// or an uuid pointed to the FileStore
}

// LimitedDownload 受限下载
func (e *TbxMisc) LimitedDownload(c *dto.TbxLimitedDownloadReq) error {
	//config.Extend.LabelCache
	lock.RLock()
	entry, ok := uuidMap[c.Uuid]
	// here, it is safer if we 'copy' entry.Data, but harmless even we refer to the data
	// lucky that GC handles the complex for us
	lock.RUnlock()
	if !ok {
		return k2Error.ErrBadRequest
	}
	c.ContentType = entry.ContentType
	if entry.Data != nil {
		// construct the reader, based on the underlay buffer
		c.Reader = bytes.NewReader(entry.Data.Bytes())
		c.ContentLength = int64(entry.Data.Len())
	} else {
		// TODO: loading data from FileStore
		reader, err := os.Open(entry.Uuid)
		if err != nil {
			return err
		}
		c.ContentLength, err = reader.Seek(0, 2)
		if err != nil {
			return err
		}
		c.Reader = reader
	}
	return nil
}

func AddLimitedDownload(url, contentType string, data *bytes.Buffer, duration time.Duration) {
	lock.Lock()
	_, ok := uuidMap[url]
	if ok {
		delete(uuidMap, url)
	}
	uuidMap[url] = dataDescriptor{ ContentType: contentType, Data: data }
	lock.Unlock()

	go func() {
		// start a timer to delete the url after a while, duration
		timer := time.NewTimer(duration)
		<-timer.C
		lock.Lock()
		delete(uuidMap, url)
		lock.Unlock()
	}()
}