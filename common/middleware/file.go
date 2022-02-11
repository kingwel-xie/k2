package middleware

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/api"
	"github.com/kingwel-xie/k2/common/config"
	cerr "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/core/utils"
)

const UploadFilePath = "static/upload/"
const DownloadUrlPrefix = "/public/downloadFile/"

var (
	failedInitFilePath = cerr.New(4500, "初始化路径失败", "failed to initialize upload path")
	failedSavingFile = cerr.New(4501, "未能保存文件", "failed to save the uploaded file")
	failedUploadToOss = cerr.New(500, "上传第三方失败", "failed to upload third party")
	requiredPicFileErr = cerr.New(500, "图片不能为空", "picture file cannot be empty")
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
	urlPrefix := DownloadUrlPrefix

	switch tag {
	case "1": // 单图
		fileResponse, err := e.singleFile(c, urlPrefix)
		if err != nil {
			e.Error(err)
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	case "2": // 多图
		multipartFile := e.multipleFile(c, urlPrefix)
		e.OK(multipartFile, "上传成功")
		return
	case "3": // base64
		fileResponse := e.baseImg(c, urlPrefix)
		e.OK(fileResponse, "上传成功")
	default:
		fileResponse, err := e.singleFile(c, urlPrefix)
		if err != nil {
			e.Error(err)
			return
		}
		e.OK(fileResponse, "上传成功")
		return
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
// @Router /public/downloadFile/{filename} [get]
// @Security Bearer
func (e File) DownloadFile(c *gin.Context) {
	var req struct {
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
		fullname := path.Join(config.FileConfig.Path, req.Filename)
		e.Context.File(fullname)
	} else {
		oss := common.Runtime.GetOss()
		reader, err := oss.DownloadFile(req.Filename)
		if err != nil {
			_ = e.Context.AbortWithError(400, failedUploadToOss.Wrap(err))
			return
		}
		defer reader.Close()

		// TODO: remove ReadAll, bad efficency
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			_ = e.Context.AbortWithError(400, failedUploadToOss.Wrap(err))
			return
		}

		contentType := http.DetectContentType(data)
		e.Context.Data(200, contentType, data)
	}
}


func (e File) baseImg(c *gin.Context, urlPerfix string) *FileResponse {
	files, _ := c.GetPostForm("file")
	file2list := strings.Split(files, ",")
	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
	guid := uuid.New().String()
	fileName := guid + ".jpg"

	base64File := UploadFilePath + fileName
	_ = ioutil.WriteFile(base64File, ddd, 0666)
	//typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
	fileResponse := &FileResponse{
		Path:     base64File,
		FullPath: urlPerfix + base64File,
		Name:     "",
	}
	err := upload(fileName, base64File)
	if err != nil {
		e.Error(failedUploadToOss.Wrap(err))
		return fileResponse
	}
	return fileResponse
}

func (e File) multipleFile(c *gin.Context, urlPerfix string) []FileResponse {
	files := c.Request.MultipartForm.File["file"]
	var multipartFile []FileResponse
	for _, f := range files {
		guid := uuid.New().String()
		fileName := guid + utils.GetExt(f.Filename)

		multipartFileName := UploadFilePath + fileName
		err1 := c.SaveUploadedFile(f, multipartFileName)
		if err1 == nil {
			err := upload(fileName, multipartFileName)
			if err != nil {
				e.Error(failedUploadToOss.Wrap(err))
			} else {
				fileResponse := FileResponse{
					Path:     multipartFileName,
					FullPath: urlPerfix + multipartFileName,
					Name:     f.Filename,
				}
				multipartFile = append(multipartFile, fileResponse)
			}
		}
	}
	return multipartFile
}

func (e File) saveFile(file *multipart.FileHeader, filename string) error {
	var err error
	if config.FileConfig.Path != "" {
		dir := path.Join(config.FileConfig.Path, e.GetIdentity().Username)
		err = utils.IsNotExistMkDir(dir)
		if err != nil {
			return failedInitFilePath.Wrap(err)
		}
		fullname := path.Join(config.FileConfig.Path, filename)
		err = e.Context.SaveUploadedFile(file, fullname)
		if err != nil {
			return failedSavingFile.Wrap(err)
		}
	} else {
		oss := common.Runtime.GetOss()
		err = oss.UploadFile(file, filename)
		if err != nil {
			return failedUploadToOss.Wrap(err)
		}
	}
	return nil
}

func (e File) singleFile(c *gin.Context, urlPerfix string) (*FileResponse, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, requiredPicFileErr.Wrap(err)
	}
	filename := e.filename(file.Filename)
	err = e.saveFile(file, filename)
	if err != nil {
		return nil, err
	}

	fileResponse := &FileResponse{
		Path:     filename,
		FullPath: urlPerfix + filename,
		Name:     file.Filename,
	}
	//fileType, _ := utils.GetType(singleFile)
	//fileResponse = FileResponse{
	//	Size:     utils.GetFileSize(singleFile),
	//	Path:     filename,
	//	FullPath: urlPerfix + singleFile,
	//	Name:     files.Filename,
	//	Type:     fileType,
	//}
	//err = upload(fileName, singleFile)
	//if err != nil {
	//	e.Error(err)
	//	return FileResponse{}, true
	//}
	//fileResponse.Path = "/static/uploadfile/" + fileName
	//fileResponse.FullPath = "/static/uploadfile/" + fileName
	return fileResponse, nil
}

func (e File) filename(oldname string) string {
	return fmt.Sprintf("[%s]%s%s", e.GetIdentity().Username, uuid.New().String(), utils.GetExt(oldname))
}
//
//func (e File) save(file *multipart.FileHeader, filename string) {
//	singleFile := UploadFilePath + fileName
//	_ = c.SaveUploadedFile(files, singleFile)
//
//}

func upload(name string, path string) error {
	log.Infof("uploading name=%s, path=%s", name, path)

	_ = common.Runtime.GetOss().UpLoadLocalFile(name, path)
	return nil
}

