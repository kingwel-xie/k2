module admin

go 1.15

require (
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/bytedance/go-tagexpr/v2 v2.7.12
	github.com/gin-gonic/gin v1.7.4
	github.com/google/uuid v1.3.0 // indirect
	github.com/kingwel-xie/k2 v1.1.25
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.8.0
	github.com/swaggo/gin-swagger v1.3.1
	github.com/swaggo/swag v1.7.1
	github.com/xuri/excelize/v2 v2.8.0
	golang.org/x/crypto v0.12.0
	gorm.io/driver/mysql v1.1.2 // indirect
	gorm.io/gorm v1.21.15
)

// replace github.com/kingwel-xie/k2 latest => ../k2
