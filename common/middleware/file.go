package middleware

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/api"
	"github.com/kingwel-xie/k2/common/config"
	cerr "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/core/utils"
)

const DownloadUrlPrefix = "public/downloadFile/"

var (
	failedInitFilePath = cerr.New(4500, "初始化路径失败", "failed to initialize upload path")
	failedSavingFile = cerr.New(4501, "未能保存文件", "failed to save the uploaded file")
	failedUploadToOss = cerr.New(4502, "上传OSS失败", "failed to upload to oss")
	requiredFileErr = cerr.New(4503, "文件不能为空", "file cannot be empty")
	failedDownloadFromOss = cerr.New(4504, "OSS下载失败", "failed to download from oss")
	failedUploadingfile = cerr.New(4505, "未能上传文件", "failed to upload")
)

type FileResponse struct {
	Path     string `json:"path"`
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
}

type File struct {
	api.Api
}

// UploadFile 上传图片
// @Summary 上传图片
// @Description 获取JSON
// @Tags 公共接口
// @Accept multipart/form-data
// @Param type query string true "type" (1：单图，2：多图, 3：base64图片)
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/public/uploadFile [post]
// @Security Bearer
func (e File) UploadFile(c *gin.Context) {
	e.MakeContext(c)

	// local path as file storage
	if config.FileConfig.Path != "" {
		err := utils.IsNotExistMkDir(config.FileConfig.Path)
		if err != nil {
			e.Error(failedInitFilePath.Wrap(err))
			return
		}
	}

	tag, _ := c.GetPostForm("type")

	urlPrefix := fmt.Sprintf("http://%s/%s", c.Request.Host, DownloadUrlPrefix)
	var response interface{}
	var err error
	switch tag {
	case "1": // 单图
		response, err = e.singleFile(c, urlPrefix)
	case "2": // 多图
		response, err = e.multipleFile(c, urlPrefix)
	case "3": // base64
		response, err = e.baseImg(c, urlPrefix)
	default:
		response, err = e.singleFile(c, urlPrefix)
	}
	if err != nil {
		e.Error(err)
	} else {
		e.OK(response, "上传成功")
	}
}

// DownloadFile 下载文件
// @Summary 下载文件
// @Description 下载文件
// @Tags 公共接口
// @Param id filename string true "filename"
// @Success 200
// @Failure 503
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /public/downloadFile/{pathname}/{filename} [get]
// @Security Bearer
func (e File) DownloadFile(c *gin.Context) {
	var req struct {
		Pathname string `uri:"pathname"`
		Filename string `uri:"filename"`
	}
	err := e.MakeContext(c).
		Bind(&req, nil).
		Errors
	if err != nil {
		_ = e.Context.AbortWithError(400, err)
		return
	}

	// local path as file storage
	if config.FileConfig.Path != "" {
		fullname := path.Join(config.FileConfig.Path, req.Pathname, req.Filename)
		e.Context.File(fullname)
	} else {
		oss := common.Runtime.GetOss()
		filename := req.Pathname + "/" + req.Filename
		reader, err := oss.DownloadFile(filename)
		if err != nil {
			_ = e.Context.AbortWithError(400, failedDownloadFromOss.Wrap(err))
			return
		}
		defer reader.Close()

		// try to get the heaers
		headers, err := oss.GetFileMeta(filename)
		if err != nil {
			_ = e.Context.AbortWithError(400, failedDownloadFromOss.Wrap(err))
			return
		}

		types, ok1 := headers["Content-Type"]
		lengths, ok2 := headers["Content-Length"]
		if ok1 && ok2 && len(types) > 0 && len(lengths) > 0 {
			contentType := types[0]
			contentLength, _ := strconv.ParseInt(lengths[0], 10, 0)
			e.Context.DataFromReader(200, contentLength, contentType, reader, nil)
		} else {
			data, err := ioutil.ReadAll(reader)
			if err != nil {
				_ = e.Context.AbortWithError(400, failedDownloadFromOss.Wrap(err))
				return
			}
			contentType := http.DetectContentType(data)
			e.Context.Data(200, contentType, data)
		}
	}
}


func (e File) baseImg(c *gin.Context, urlPerfix string) (*FileResponse, error) {
	files, ok := c.GetPostForm("file")
	if !ok {
		return nil, requiredFileErr
	}
	file2list := strings.Split(files, ",")
	if len(file2list) < 2 {
		return nil, requiredFileErr.Wrapf("wrong base64 image format")
	}
	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])

	// get ext name from file2list[0]
	filename := e.filename("*.jpg")
	err := e.saveFile(bytes.NewReader(ddd), filename)
	if err != nil {
		return nil, err
	}

	fileResponse := &FileResponse{
		Path:     filename,
		FullPath: urlPerfix + filename,
		Name:     "",
	}
	return fileResponse, nil
}

func (e File) multipleFile(c *gin.Context, urlPerfix string) ([]FileResponse, error) {
	files, ok := c.Request.MultipartForm.File["file"]
	if !ok {
		return nil, requiredFileErr
	}

	var multipartFile []FileResponse
	for _, f := range files {
		filename := e.filename(f.Filename)
		reader, err := f.Open()
		if err != nil {
			continue
		}

		err = e.saveFile(reader, filename)
		reader.Close()
		if err != nil {
			continue
		}

		fileResponse := FileResponse{
			Path:     filename,
			FullPath: urlPerfix + filename,
			Name:     f.Filename,
		}
		multipartFile = append(multipartFile, fileResponse)
	}
	return multipartFile, nil
}

func (e File) saveFile(file io.Reader, filename string) error {
	var err error
	if config.FileConfig.Path != "" {
		dir := path.Join(config.FileConfig.Path, e.GetIdentity().Username)
		err = utils.IsNotExistMkDir(dir)
		if err != nil {
			return failedInitFilePath.Wrap(err)
		}
		fullname := path.Join(config.FileConfig.Path, filename)
		out, err := os.Create(fullname)
		if err != nil {
			return failedSavingFile.Wrap(err)
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			return failedSavingFile.Wrap(err)
		}
	} else {
		oss := common.Runtime.GetOss()
		_, err = oss.UploadFile(file, filename)
		if err != nil {
			return failedUploadToOss.Wrap(err)
		}
	}
	return nil
}

func (e File) singleFile(c *gin.Context, urlPerfix string) (*FileResponse, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, requiredFileErr.Wrap(err)
	}
	filename := e.filename(file.Filename)
	reader, err := file.Open()
	if err != nil {
		return nil, failedUploadingfile.Wrap(err)
	}
	defer reader.Close()

	err = e.saveFile(reader, filename)
	if err != nil {
		return nil, err
	}

	fileResponse := &FileResponse{
		Path:     filename,
		FullPath: urlPerfix + filename,
		Name:     file.Filename,
	}
	return fileResponse, nil
}

func (e File) filename(oldname string) string {
	return fmt.Sprintf("%s/%s%s", e.GetIdentity().Username, uuid.New().String(), utils.GetExt(oldname))
}
