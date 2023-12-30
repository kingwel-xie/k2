package utils

import (
	"github.com/kingwel-xie/k2/common"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func init() {
	common.SetupRuntimeForTest("../kobh.test.db")

	db := common.Runtime.GetDb()

	db.Exec("delete from tbi_klock")
}

func TestKeyLock(t *testing.T) {
	db := common.Runtime.GetDb()
	if db == nil {
		t.Error("Test not run, db is nil")
		return
	}
	go func() {
		for range time.Tick(2 * time.Second) {
			cleanup()
		}
	}()
	assert := require.New(t)

	keyLock := "keyLock"
	assert.NoError(Lock(db, keyLock))
	assert.NotEmpty(Lock(db, keyLock))
	Unlock(db, keyLock)
	assert.NoError(Lock(db, keyLock))
	Unlock(db, keyLock)

	assert.Equal(TryLock(db, keyLock, 1), true)
	assert.Equal(TryLock(db, keyLock, 2), false)
	Unlock(db, keyLock)
	assert.Equal(TryLock(db, keyLock, 1), true)
	Unlock(db, keyLock)

	assert.NoError(Lock(db, keyLock, 1))
	assert.NotEmpty(Lock(db, keyLock, 1), false)
	assert.Equal(TryLock(db, keyLock, 2), true)
	Unlock(db, keyLock)
}

//
//func TestKeyLockWithMultiGo(t *testing.T) {
//	db := common.Runtime.GetDb()
//	if db == nil {
//		t.Error("Test not run, db is nil")
//		return
//	}
//	mutex := sync.Mutex{}
//
//	wg := sync.WaitGroup{}
//	for c := 0; c < 10; c++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for i := 0; i < 100; i++ {
//				key := uuid.New().String()
//				mutex.Lock()
//				if Lock(db, key) != nil {
//					t.Error("Lock db failed")
//				}
//				Unlock(db, key)
//				mutex.Unlock()
//			}
//		}()
//	}
//	wg.Wait()
//}
