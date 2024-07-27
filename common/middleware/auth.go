package middleware

import (
	"github.com/kingwel-xie/k2/core/utils"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/common/middleware/jwtauth"
)

// AuthInit jwt验证new
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	sendCookie := false
	if config.ApplicationConfig.Mode == utils.ModeDev.String() {
		timeout = time.Duration(876010) * time.Hour
		sendCookie = true
	} else {
		if config.JwtConfig.Timeout != 0 {
			timeout = time.Duration(config.JwtConfig.Timeout) * time.Second
		}
	}
	// extra initialization for using EntraId
	if config.EntraConfig.Enable && config.EntraConfig.Mgmt.SecretKey != "" {
		err := jwtauth.SyncUpEntra()
		if err != nil {
			log.Warnf("failed to sync-up SysUser with EntraId, %v", err)
		}
	}
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:          "kobh zone",
		Key:            []byte(config.JwtConfig.Secret),
		Timeout:        timeout,
		MaxRefresh:     time.Hour,
		PayloadFunc:    jwtauth.PayloadFunc,
		Authenticator:  jwtauth.Authenticator,
		Authorizator:   jwtauth.Authorizator,
		Unauthorized:   jwtauth.Unauthorized,
		LogoutResponse: jwtauth.LogoutResponse,
		TokenLookup:    "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:  "Bearer",
		TimeFunc:       time.Now,
		SendCookie:     sendCookie,
	})
}
