package middleware

// DownloadFile 下载文件
// @Summary 下载文件
// @Description 下载文件
// @Tags 公共接口
// @Param pathname path string true "pathname"
// @Param filename path string true "filename"
// @Param as query string true "as"
// @Success 200
// @Failure 503
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /public/downloadFile/{pathname}/{filename} [get]
// @Security Bearer
//func (e File) DownloadFile(c *gin.Context) {
//
//}