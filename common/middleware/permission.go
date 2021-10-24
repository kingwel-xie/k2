package middleware

import (
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/core/utils"
	"net/http"

	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/api"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/common/response"
	"github.com/kingwel-xie/k2/common/service"
)

// AuthCheckRole 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.MustGetLogger(c)

		// get jwt claims
		v := service.GetIdentity(c)
		// casbin
		e := common.Runtime.GetCasbin()

		var res, casbinExclude bool
		var err error
		//检查权限
		if v.RoleKey == "admin" {
			res = true
			c.Next()
			return
		}
		// DEV mode, check CasbinExclude
		if config.ApplicationConfig.Mode == utils.ModeDev.String() {
			for _, i := range CasbinExclude {
				if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
					casbinExclude = true
					break
				}
			}
			if casbinExclude {
				log.Errorf("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
				c.Next()
				return
			}
		}
		res, err = e.Enforce(v.RoleKey, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			log.Errorf("AuthCheckRole error: %s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
			response.Error(c, 500, err, "内部错误")
			return
		}

		if res {
			log.Infof("isTrue: %v, role: %s method: %s path: %s", res, v.RoleKey, c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else {
			log.Warnf("isTrue: %v, role: %s method: %s path: %s message: %s", res, v.RoleKey, c.Request.Method, c.Request.URL.Path, "当前request无权限，请管理员确认！")
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
			})
			c.Abort()
			return
		}

	}
}
