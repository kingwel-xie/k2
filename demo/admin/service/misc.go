package service

import (
	k2Error "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/service"
	"io/ioutil"
	"admin/service/dto"
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
	Data []byte			// file data if Data is not nil
	Uuid string			// or an uuid pointed to the FileStore
}

// LimitedDownload 受限下载
func (e *TbxMisc) LimitedDownload(c *dto.TbxLimitedDownloadReq) error {
	var err error

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
		c.Data = entry.Data
	} else {
		// TODO: loading data from FileStore
		c.Data, err = ioutil.ReadFile(entry.Uuid)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddLimitedDownload(url, contentType string, data []byte, duration time.Duration) {
	lock.Lock()
	_, ok := uuidMap[url]
	if ok {
		delete(uuidMap, url)
	}
	uuidMap[url] = dataDescriptor{ ContentType: contentType, Data: data }
	lock.Unlock()

	go func() {
		// start a timer to delete the url after a while, 10s
		timer := time.NewTimer(10 * time.Second)
		<-timer.C
		lock.Lock()
		delete(uuidMap, url)
		lock.Unlock()
	}()
}