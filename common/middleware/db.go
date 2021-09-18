package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/common"
)

func WithContextDb(c *gin.Context) {
	c.Set("db", common.Runtime.GetDb().WithContext(c))
	c.Next()
}
