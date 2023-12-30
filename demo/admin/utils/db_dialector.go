package utils

import (
	"fmt"
	k2Error "github.com/kingwel-xie/k2/common/error"
	"gorm.io/gorm"
)

func AssertUpdateAffected(db *gorm.DB, expected int) error {
	if db.Error != nil {
		return db.Error
	}
	if db.Dialector.Name() == "sqlite" {
		if db.RowsAffected != int64(expected) {
			return k2Error.ErrPermissionDenied
		}
	} else {
		// https://stackoverflow.com/questions/3747314/why-are-2-rows-affected-in-my-insert-on-duplicate-key-update
		// https://dev.mysql.com/doc/refman/8.0/en/insert-on-duplicate.html
		// mysql
		// With ON DUPLICATE KEY UPDATE, the affected-rows value per row is 1 if the row is inserted as a new row and 2 if an existing row is updated.
		if db.RowsAffected != int64(expected*2) {
			return k2Error.ErrPermissionDenied
		}
	}
	return nil
}

// UpdateConcatSyntax per SQL dialect
func UpdateConcatSyntax(db *gorm.DB) string {
	if db.Dialector.Name() == "sqlite" {
		return "log||?"
	} else {
		return "concat(log, ?)"
	}
}

func SqlFormatMonth(db *gorm.DB, column string) string {
	return SqlFormatDate(db, "%m", column)
}

func SqlFormatYear(db *gorm.DB, column string) string {
	return SqlFormatDate(db, "%Y", column)
}

func SqlFormatYearMonth(db *gorm.DB, column string) string {
	return SqlFormatDate(db, "%Y-%m", column)
}

func SqlFormatISOWeek(db *gorm.DB, column string) string {
	if IsSqlite(db) {
		return fmt.Sprintf("strftime('%%W', %s)", column)
	} else {
		return fmt.Sprintf("WEEK(%s, 3)", column)
	}
}

func SqlFormatDate(db *gorm.DB, format string, column string) string {
	if IsSqlite(db) {
		return fmt.Sprintf("strftime('%s', %s)", format, column)
	} else {
		return fmt.Sprintf("DATE_FORMAT(%s, '%s')", column, format)
	}
}

func IsSqlite(db *gorm.DB) bool {
	return db.Dialector.Name() == "sqlite"
}

func MigrateView(db *gorm.DB, viewName string, sql string) error {
	err := db.Exec(fmt.Sprintf(`DROP VIEW IF EXISTS %s`, viewName)).Error
	if err != nil {
		return err
	}
	return db.Exec(fmt.Sprintf("CREATE VIEW %s AS %s", viewName, sql)).Error
}
