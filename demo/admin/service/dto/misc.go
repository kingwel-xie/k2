package dto

import "io"

// TbxLimitedDownloadReq 受限下载请求参数
type TbxLimitedDownloadReq struct {
	Uuid     string `uri:"uuid"`
	Filename string `form:"filename"`
	// response
	ContentType   string    `json:"-"`
	ContentLength int64     `json:"-"`
	Reader        io.Reader `json:"-"`
}

