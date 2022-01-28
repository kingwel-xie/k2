package common

import (
	"github.com/kingwel-xie/k2/core/oss"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/kingwel-xie/k2/core/cronjob"
	"github.com/kingwel-xie/k2/core/storage"
	"github.com/kingwel-xie/k2/core/storage/queue"
)

var Runtime = NewApplication()

type Application struct {
	mux     sync.RWMutex
	db      *gorm.DB
	casbin  *casbin.SyncedEnforcer
	crontab *cronjob.CronJob
	engine  http.Handler
	cache   storage.AdapterCache
	queue   storage.AdapterQueue
	locker  storage.AdapterLocker
	oss     oss.Oss
	routers []Router
}

type Router struct {
	HttpMethod, RelativePath, Handler string
}



type Routers struct {
	List []Router
}

// SetDb 设置对应key的db
func (e *Application) SetDb(db *gorm.DB) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.db = db
}

// GetDb 获取所有map里的db数据
func (e *Application) GetDb() *gorm.DB {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.db
}

func (e *Application) SetCasbin(enforcer *casbin.SyncedEnforcer) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.casbin = enforcer
}

func (e *Application) GetCasbin() *casbin.SyncedEnforcer {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.casbin
}

// SetEngine 设置路由引擎
func (e *Application) SetEngine(engine http.Handler) {
	e.engine = engine
}

// GetEngine 获取路由引擎
func (e *Application) GetEngine() http.Handler {
	return e.engine
}

type routerSlice []Router

func (x routerSlice) Len() int           { return len(x) }
func (x routerSlice) Less(i, j int) bool { return strings.Compare(x[i].Handler, x[j].Handler) < 0 }
func (x routerSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// GetRouter 获取路由表
func (e *Application) GetRouter() []Router {
	switch e.engine.(type) {
	case *gin.Engine:
		routers := e.engine.(*gin.Engine).Routes()
		for _, router := range routers {
			e.routers = append(e.routers, Router{RelativePath: router.Path, Handler: router.Handler, HttpMethod: router.Method})
		}
	}
	s := routerSlice(e.routers)
	sort.Sort(s)
	return s
}

// NewApplication 默认值
func NewApplication() Application {
	return Application{
		routers: make([]Router, 0),
	}
}

// SetCrontab 设置对应key的crontab
func (e *Application) SetCrontab(crontab *cronjob.CronJob) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.crontab = crontab
}

// GetCrontab 获取所有map里的crontab数据
func (e *Application) GetCrontab() *cronjob.CronJob {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.crontab
}

// SetCacheAdapter 设置缓存
func (e *Application) SetCacheAdapter(c storage.AdapterCache) {
	e.cache = c
}

// WxCache 获取cache
func (e *Application) Cache() storage.AdapterCache {
	return e.cache
}

// SetQueueAdapter 设置队列适配器
func (e *Application) SetQueueAdapter(c storage.AdapterQueue) {
	e.queue = c
}

// Queue 获取queue
func (e *Application) Queue() storage.AdapterQueue {
	return e.queue
}

// SetLockerAdapter 设置分布式锁
func (e *Application) SetLockerAdapter(c storage.AdapterLocker) {
	e.locker = c
}

// Locker 获取分布式锁
func (e *Application) Locker() storage.AdapterLocker {
	return e.locker
}

// SetOss 设置Oss service
func (e *Application) SetOss(oss oss.Oss) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.oss = oss
}

// SetOss 设置Oss service
func (e *Application) GetOss() oss.Oss {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.oss
}

// GetStreamMessage 获取队列需要用的message
func (e *Application) GetStreamMessage(id, stream string, value map[string]interface{}) (storage.Messager, error) {
	message := &queue.Message{}
	message.SetID(id)
	message.SetStream(stream)
	message.SetValues(value)
	return message, nil
}
