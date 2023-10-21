package middleware

import (
	"fmt"
	"github.com/kingwel-xie/k2/core/utils"
	"github.com/shirou/gopsutil/host"
	"net/http"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// 获取相差时间
func GetHourDiffer(startTime, endTime string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

// ServerInfo 获取系统信息
// @Summary 系统信息
// @Description 获取JSON
// @Tags 系统信息
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /server-monitor [get]
// @Security Bearer
func ServerInfo(custom func(ret map[string]interface{}) map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		sysInfo, err := host.Info()
		osDic := make(map[string]interface{}, 0)
		osDic["goOs"] = runtime.GOOS
		osDic["arch"] = runtime.GOARCH
		osDic["mem"] = runtime.MemProfileRate
		osDic["compiler"] = runtime.Compiler
		osDic["version"] = runtime.Version()
		osDic["numGoroutine"] = runtime.NumGoroutine()
		osDic["ip"] = utils.GetLocaHonst()
		osDic["projectDir"] = utils.GetCurrentPath()
		osDic["hostName"] = sysInfo.Hostname
		osDic["time"] = time.Now().Format("2006-01-02 15:04:05")

		dis, _ := disk.Usage("/")
		diskTotalGB := int(dis.Total) / GB
		diskFreeGB := int(dis.Free) / GB
		diskDic := make(map[string]interface{}, 0)
		diskDic["total"] = diskTotalGB
		diskDic["free"] = diskFreeGB

		mem, _ := mem.VirtualMemory()
		memUsedMB := int(mem.Used) / GB
		memTotalMB := int(mem.Total) / GB
		memFreeMB := int(mem.Free) / GB
		memUsedPercent := int(mem.UsedPercent)
		memDic := make(map[string]interface{}, 0)
		memDic["total"] = memTotalMB
		memDic["used"] = memUsedMB
		memDic["free"] = memFreeMB
		memDic["usage"] = memUsedPercent

		cpuDic := make(map[string]interface{}, 0)
		cpuDic["cpuInfo"], _ = cpu.Info()
		percent, _ := cpu.Percent(0, false)
		cpuDic["Percent"] = utils.Round(percent[0], 2)
		cpuDic["cpuNum"], _ = cpu.Counts(false)

		//服务器磁盘信息
		disklist := make([]disk.UsageStat, 0)
		//所有分区
		diskInfo, err := disk.Partitions(true)
		if err == nil {
			for _, p := range diskInfo {
				diskDetail, err := disk.Usage(p.Mountpoint)
				if err == nil {
					diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
					diskDetail.Total = diskDetail.Total / 1024 / 1024
					diskDetail.Used = diskDetail.Used / 1024 / 1024
					diskDetail.Free = diskDetail.Free / 1024 / 1024
					disklist = append(disklist, *diskDetail)
				}
			}
		}
		data := map[string]interface{}{
			"requestId": utils.GenerateMsgIDFromContext(c),
			"code":      200,
			"os":        osDic,
			"mem":       memDic,
			"cpu":       cpuDic,
			"disk":      diskDic,
			"diskList":  disklist,
		}
		if custom != nil {
			data = custom(data)
		}
		c.Set("result", data)
		c.AbortWithStatusJSON(http.StatusOK, data)
	}
}
