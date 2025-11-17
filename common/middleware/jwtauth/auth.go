package jwtauth

import (
	"errors"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"

	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/common/api"
	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/common/global"
	"github.com/kingwel-xie/k2/common/service"
	"github.com/kingwel-xie/k2/core/captcha"
	"github.com/kingwel-xie/k2/core/utils"
)

var ErrInvalidVerificationCode = errors.New("invalid verification code")

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(SysUser)
		r, _ := v["role"].(SysRole)
		identity := map[string]interface{}{
			"UserId":    u.UserId,
			"Username":  u.Username,
			"DeptId":    u.DeptId,
			"RoleId":    r.RoleId,
			"RoleKey":   r.RoleKey,
			"RoleName":  r.RoleName,
			"DataScope": r.DataScope,
		}
		return jwt.MapClaims{
			jwt.IdentityKey: identity,
		}
	}
	return jwt.MapClaims{}
}

// Authenticator 获取token
// @Summary 登陆
// @Description 获取token
// @Description LoginHandler can be used by clients to get a jwt token.
// @Description Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// @Description Reply will be of the form {"token": "TOKEN"}.
// @Description dev mode：It should be noted that all fields cannot be empty, and a value of 0 can be passed in addition to the account password
// @Description 注意：开发模式：需要注意全部字段不能为空，账号密码外可以传入0值
// @Tags 登陆
// @Accept  application/json
// @Product application/json
// @Param account body Login  true "account"
// @Success 200 {string} string "{"code": 200, "expire": "2019-08-07T12:45:48+08:00", "token": ".eyJleHAiOjE1NjUxNTMxNDgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU2NTE0OTU0OH0.-zvzHvbg0A" }"
// @Router /api/v1/login [post]
func Authenticator(c *gin.Context) (interface{}, error) {
	log := api.MustGetLogger(c)
	db := api.MustGetOrm(c)

	var loginVals Login
	var status = "2"
	var msg = "登录成功"
	var username = ""
	defer func() {
		LoginLogToDB(c, status, msg, username)
	}()

	if err := c.ShouldBind(&loginVals); err != nil {
		username = loginVals.Username
		msg = "参数解析失败: " + err.Error()
		status = "1"
		log.Warnw(msg, "user", username)

		return nil, jwt.ErrMissingLoginValues
	}

	if loginVals.WechatId != "" {
		var user SysUser
		// 尝试通过微信ID查找用户
		err := db.Table("sys_user").Where("wechat_id = ? and status = 2", loginVals.WechatId).First(&user).Error
		if err == nil {
			// 已绑定微信ID的用户，直接登录
			var role SysRole
			db.Table("sys_role").Where("role_id = ?", user.RoleId).First(&role)
			username = user.Username
			msg = "微信登录成功"
			log.Infow(msg, "user", username)
			return map[string]interface{}{"user": user, "role": role}, nil
		}

		// 校验用户名密码
		user, role, e := loginVals.GetUser(db)
		if e != nil {
			msg = "用户名或密码错误"
			status = "1"
			log.Warnw(msg, "user", loginVals.Username, "error", e)
			return nil, jwt.ErrFailedAuthentication
		}

		// 绑定微信ID
		err = db.Model(&user).Update("wechat_id", loginVals.WechatId).Error
		if err != nil {
			msg = "绑定微信ID失败"
			status = "1"
			log.Warnw(msg, "user", user.Username, "error", err)
			return nil, errors.New("绑定微信ID失败")
		}

		username = user.Username
		msg = "微信绑定并登录成功"
		log.Infow(msg, "user", username)
		return map[string]interface{}{"user": user, "role": role}, nil
	}

	if config.ApplicationConfig.Mode != utils.ModeDev.String() {
		if !captcha.Verify(loginVals.UUID, loginVals.Code, true) {
			username = loginVals.Username
			msg = "验证码错误"
			status = "1"
			log.Warnw(msg, "user", username)

			return nil, ErrInvalidVerificationCode
		}
	}
	user, role, e := loginVals.GetUser(db)
	if e == nil {
		username = loginVals.Username
		msg = "登录成功"
		log.Infow(msg, "user", username)
		return map[string]interface{}{"user": user, "role": role}, nil
	} else {
		username = loginVals.Username
		msg = "登录失败"
		status = "1"
		log.Warnw(msg, "user", username, "error", e)
	}
	return nil, jwt.ErrFailedAuthentication
}

func LogoutResponse(c *gin.Context, code int) {
	LoginLogToDB(c, "2", "退出成功", service.GetIdentity(c).Username)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

// Authorizator See AuthCheckRole handler
func Authorizator(data interface{}, c *gin.Context) bool {
	//identity := data.(map[string]interface{})
	//_, ok := identity["UserId"]
	return true
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

// LoginLogToDB Write log to database
func LoginLogToDB(c *gin.Context, status string, msg string, username string) {
	if !config.LoggerConfig.EnabledDB {
		return
	}

	log := api.MustGetLogger(c)
	l := make(map[string]interface{})

	ua := user_agent.New(c.Request.UserAgent())
	l["ipaddr"] = utils.GetClientIP(c)
	l["loginTime"] = utils.GetCurrentTime()
	l["status"] = status
	l["remark"] = c.Request.UserAgent()
	browserName, browserVersion := ua.Browser()
	l["browser"] = browserName + " " + browserVersion
	l["os"] = ua.OS()
	l["platform"] = ua.Platform()
	l["username"] = username
	l["msg"] = msg

	message, err := common.Runtime.GetStreamMessage("", global.LoginLog, l)
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
