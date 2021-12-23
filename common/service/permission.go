package service

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/core/utils"
	"gorm.io/gorm"
)

type AuthIdentity struct {
	UserId    int
	Username  string
	DeptId    int
	RoleId    int
	RoleKey   string
	RoleName  string
	DataScope string
}

func GetIdentity(c *gin.Context) *AuthIdentity {
	raw, exists := c.Get(jwt.IdentityKey)
	if exists {
		data := raw.(map[string]interface{})
		return &AuthIdentity{
			UserId:    int(data["UserId"].(float64)),
			Username:  data["Username"].(string),
			DeptId:    int(data["DeptId"].(float64)),
			RoleId:    int(data["RoleId"].(float64)),
			RoleName:  data["RoleName"].(string),
			RoleKey:   data["RoleKey"].(string),
			DataScope: data["DataScope"].(string),
		}
	} else {
		return &NoAuthIdentity
	}
}

// Permission check permission of data scope
func Permission(tableName string, p *AuthIdentity) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		//if !config.ApplicationConfig.EnableDP {
		//	return db
		//}
		switch p.DataScope {
		case "2":
			return db.Where(tableName+".create_by in (SELECT sys_user.username from sys_role_dept left join sys_user on sys_user.dept_id=sys_role_dept.dept_id where sys_role_dept.role_id = ?)", p.RoleId)
		case "3":
			return db.Where(tableName+".create_by in (SELECT username from sys_user where dept_id = ? )", p.DeptId)
		case "4":
			return db.Where(tableName+".create_by in (SELECT username from sys_user where sys_user.dept_id in(select dept_id from sys_dept where dept_path like ? ))", "%/"+utils.IntToString(p.DeptId)+"/%")
		case "5":
			return db.Where(tableName+".create_by = ?", p.Username)
		default:
			return db
		}
	}
}
