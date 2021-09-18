package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/kingwel-xie/k2/common/config"
	"github.com/kingwel-xie/k2/common/middleware/jwtauth"
)

// AuthInit jwt验证new
func AuthInit(prefix string) (*jwt.GinJWTMiddleware, error) {
	// set the preix to allow using different user+role tables
	jwtauth.SetTablePrefix(prefix)

	timeout := time.Hour
	sendCookie := false
	if config.ApplicationConfig.Mode == "dev" {
		timeout = time.Duration(876010) * time.Hour
		sendCookie = true
	} else {
		if config.JwtConfig.Timeout != 0 {
			timeout = time.Duration(config.JwtConfig.Timeout) * time.Second
		}
	}
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "github.com/kingwel-xie/k2 zone",
		Key:             []byte(config.ApplicationConfig.JwtSecret),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     jwtauth.PayloadFunc,
		Authenticator:   jwtauth.Authenticator,
		Authorizator:    jwtauth.Authorizator,
		Unauthorized:    jwtauth.Unauthorized,
		LogoutResponse:  jwtauth.LogoutResponse,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		SendCookie:      sendCookie,
	})
}
