package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	cerr "github.com/kingwel-xie/k2/common/error"
	"github.com/kingwel-xie/k2/common/response"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthCheckRole(t *testing.T) {
	signature := ""
	router := gin.New()
	router.Use(authCheckRole())

	router.GET("/", func(c *gin.Context) {
		signature += "D"
	})

	adminHeader := header{
		Key:   "isAdmin",
		Value: "admin",
	}
	w := performRequest(router, "GET", "/",adminHeader)

	// TEST 有权限
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "D", signature)

	visitorHeader := header{
		Key:   "isAdmin",
		Value: "visitor",
	}

	// TEST 无权限
	w2 := performRequest(router, "GET", "/",visitorHeader)
	resp := response.Response{}
	err := json.Unmarshal(w2.Body.Bytes(), &resp)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, int32(http.StatusForbidden), resp.Code)
	assert.Equal(t, "D", signature)

}

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func authCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := c.Request.Header.Get("isAdmin")
		if isAdmin == "admin" {
			log.Debug("角色通过")
			c.Next()
		} else {
			log.Warnf( "当前request无权限，请管理员确认！")
			response.Error(c,cerr.ErrNoPermission)
			c.Abort()
			return
		}
	}
}