package entry

import (
	"github.com/kingwel-xie/k2/common"
	"gorm.io/gorm"
	"kobh/models"
	"time"
)

// startDatabaseKeeper clean up some tables periodically
func startDatabaseKeeper(db *gorm.DB) {
	// start a cron job to clean DB, every 2nd, at 0:01am
	_ = common.Runtime.GetCrontab().AddJob("dbKeeper", "0 1 0 1 * ?", func() {
		clean(db, &models.SysOperaLog{}, 6)
		clean(db, &models.SysLoginLog{}, 6)
	})
}

func clean(db *gorm.DB, which interface{}, months int)  {
	someMonthsAgo := time.Now().AddDate(0, -months, 0)
	db.Delete(which, "created_at < ?", someMonthsAgo)
}