package utils

import (
	"errors"
	"kobh/common/defs"
	"kobh/models"
	"time"

	"github.com/kingwel-xie/k2/common"
	"gorm.io/gorm"
)

const (
	defaultLockDuration = 600 // second
	cleanupInterval     = 300 // second
)

func init() {
	go func() {
		for range time.Tick(cleanupInterval * time.Second) {
			cleanup()
		}
	}()
}

func cleanup() {
	db := common.Runtime.GetDb()
	if db == nil {
		return
	}

	db.Where("Expire < ?", time.Now()).Delete(&models.TbiKlock{})
}

// Lock locks a key, expiry specified by params[0]
func Lock(db *gorm.DB, key string, params ...interface{}) error {
	if len(key) == 0 {
		panic(defs.ErrInvalidParameter)
	}
	expiry := defaultLockDuration
	if len(params) > 0 {
		value, ok := params[0].(int)
		if !ok {
			panic(defs.ErrInvalidParameter)
		}
		expiry = value
	}
	return lock(db, key, expiry)
}

// TryLock try to lock a key with max retries
func TryLock(db *gorm.DB, key string, maxRetries int) bool {
	if len(key) == 0 {
		return false
	}

	retries := 0
	for {
		if retries > maxRetries {
			return false
		}
		if err := Lock(db, key); err == nil {
			return true
		}
		time.Sleep(time.Duration(10) * time.Second)
		retries += 1
	}
}

// lock locks a key with specific expiry
func lock(db *gorm.DB, key string, expiry int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var keyList []models.TbiKlock
		err := tx.Find(&keyList, "code = ?", key).Error
		if err != nil {
			// quite unlikely, db error
			return err
		}

		// if key exists, it means the key was locked
		if len(keyList) > 0 {
			return errors.New("key already locked")
		}

		model := &models.TbiKlock{
			Code:   key,
			Expire: time.Now().Add(time.Duration(expiry) * time.Second),
		}
		// this will lock the specific key
		err = tx.Create(model).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func Unlock(db *gorm.DB, key string) {
	model := &models.TbiKlock{Code: key}
	db.Delete(model)
}
