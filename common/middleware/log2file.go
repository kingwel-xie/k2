package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/kingwel-xie/k2/core/utils"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/api"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/common/global"
	"github.com/kingwel-xie/k2/common/service"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// don't care http.MethodOptions
		if c.Request.Method == http.MethodOptions {
			return
		}

		log := api.MustGetLogger(c)
		// 开始时间
		startTime := time.Now()
		// 处理请求
		var body string
		// make body only when EnabledDB is on
		if config.LoggerConfig.EnabledDB {
			switch c.Request.Method {
			case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
				bf := bytes.NewBuffer(nil)
				wt := bufio.NewWriter(bf)
				_, err := io.Copy(wt, c.Request.Body)
				if err != nil {
					log.Warnf("copy body error, %s", err.Error())
					err = nil
				}
				rb, _ := ioutil.ReadAll(bf)
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
				body = string(rb)
			}
		}

		c.Next()

		url := c.Request.RequestURI
		if strings.Index(url, "logout") > -1 ||
			strings.Index(url, "login") > -1 {
			return
		}

		// 执行时间
		latencyTime := time.Now().Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := utils.GetClientIP(c)

		// API code
		st, bl := c.Get("status")
		var statusBus = 0
		if bl {
			statusBus = st.(int)
		}

		log.Infow("", "status", statusCode, "api-code", statusBus, "method", reqMethod, "uri", reqUri, "latency", latencyTime, "client", clientIP)

		if config.LoggerConfig.EnabledDB && statusCode != 404 {
			rt, bl := c.Get("result")
			var result = ""
			if bl {
				rb, err := json.Marshal(rt)
				if err != nil {
					log.Warnf("json Marshal result error, %s", err.Error())
				} else {
					result = string(rb)
				}
			}
			SetDBOperLog(c, clientIP, statusCode, reqUri, reqMethod, latencyTime, body, result, statusBus)
		}
	}
}

// SetDBOperLog 写入操作日志表 fixme 该方法后续即将弃用
func SetDBOperLog(c *gin.Context, clientIP string, statusCode int, reqUri string, reqMethod string, latencyTime time.Duration, body string, result string, status int) {
	log := api.MustGetLogger(c)
	l := make(map[string]interface{})
	l["_fullPath"] = c.FullPath()
	l["operUrl"] = reqUri
	l["operIp"] = clientIP
	l["operName"] = service.GetIdentity(c).Username
	l["requestMethod"] = c.Request.Method
	l["operParam"] = body
	l["operTime"] = time.Now()
	l["jsonResult"] = result
	l["latencyTime"] = latencyTime.String()
	l["statusCode"] = statusCode
	if status == http.StatusOK {
		l["status"] = "2"
	} else {
		l["status"] = "1"
	}
	message, err := common.Runtime.GetStreamMessage("", global.OperateLog, l)
	if err != nil {
		log.Errorf("GetStreamMessage error, %s", err.Error())
		//日志报错错误，不中断请求
	} else {
		err = common.Runtime.Queue().Append(message)
		if err != nil {
			log.Errorf("Append message error, %s", err.Error())
		}
	}
}
