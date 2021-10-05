package version

import (
	"fmt"
	"runtime"

	common "github.com/kingwel-xie/k2/common/models"
	"github.com/kingwel-xie/k2/core/migration"
	"github.com/kingwel-xie/k2/core/utils"
	"gorm.io/gorm"

	"admin/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1599190683659Tables)
}

func _1599190683659Tables(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Migrator().AutoMigrate(
			new(models.SysDept),
			new(models.SysConfig),
			new(models.SysMenu),
			new(models.SysLoginLog),
			new(models.SysOperaLog),
			new(models.SysUser),
			new(models.SysRole),
			new(models.SysPost),
			new(models.SysDictData),
			new(models.SysDictType),
			new(models.SysConfig),
			new(models.SysApi),
			new(models.TbxCountry),
		)
		if err != nil {
			return err
		}
		sqlFileName := fmt.Sprintf("config/%s.sql", tx.Dialector.Name())
		if err := utils.ExecSql(tx, sqlFileName); err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
