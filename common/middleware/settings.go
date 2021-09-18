package middleware

type UrlInfo struct {
	Url    string
	Method string
}

// CasbinExclude casbin 排除的路由列表
var CasbinExclude = []UrlInfo{
	{Url: "/api/v1/dict/type-option-select", Method: "GET"},
	{Url: "/api/v1/dict-data/option-select", Method: "GET"},

	{Url: "/api/v1/deptTree", Method: "GET"},

	{Url: "/api/v1/menu-role", Method: "GET"},
	//{Url: "/api/v1/menuids", Method: "GET"},

	{Url: "/api/v1/configKey/:configKey", Method: "GET"},
	{Url: "/api/v1/app-config", Method: "GET"},

	{Url: "/api/v1/getinfo", Method: "GET"},
	{Url: "/api/v1/user/profile", Method: "GET"},
	{Url: "/api/v1/user/avatar", Method: "POST"},
	{Url: "/api/v1/user/pwd", Method: "PUT"},
	{Url: "/api/v1/user/status", Method: "PUT"},

	{Url: "/api/v1/roleMenuTreeSelect/:roleId", Method: "GET"},
	{Url: "/api/v1/roleDeptTreeSelect/:roleId", Method: "GET"},
	{Url: "/api/v1/public/uploadFile", Method: "POST"},

	// probably we should comment below, which are actually under no-auth
	{Url: "/api/v1/captcha", Method: "GET"},
	{Url: "/api/v1/login", Method: "POST"},
	{Url: "/api/v1/logout", Method: "POST"},
	{Url: "/api/v1/refresh_token", Method: "GET"},
	{Url: "/metrics", Method: "GET"},
	{Url: "/health", Method: "GET"},
	{Url: "/", Method: "GET"},
	{Url: "/info", Method: "GET"},
	{Url: "/server-monitor", Method: "GET"},
}
