package migration

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

var Migrate = &Migration{
	version: make(map[int]func(db *gorm.DB, version string) error),
}

type Migration struct {
	db      *gorm.DB
	version map[int]func(db *gorm.DB, version string) error
	mutex   sync.Mutex
}

func (e *Migration) GetDb() *gorm.DB {
	return e.db
}

func (e *Migration) SetDb(db *gorm.DB) {
	e.db = db
}

func (e *Migration) SetVersion(k int, f func(db *gorm.DB, version string) error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.version[k] = f
}

func (e *Migration) Migrate() {
	versions := make([]int, 0)
	for k := range e.version {
		versions = append(versions, k)
	}
	if !sort.IntsAreSorted(versions) {
		sort.Ints(versions)
	}
	var err error
	var count int64
	for _, v := range versions {
		err = e.db.Table("sys_migration").Where("version = ?", v).Count(&count).Error
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if count > 0 {
			fmt.Printf("Version %d has been migrated, count=%d\n", v, count)
			count = 0
			continue
		}
		err = (e.version[v])(e.db.Debug(), strconv.Itoa(v))
		if err != nil {
			fmt.Println("Migration failed:", err)
			os.Exit(1)
		}
	}
}

func GetFilename(s string) int {
	s = filepath.Base(s)
	return cast.ToInt(s[:13])
}
