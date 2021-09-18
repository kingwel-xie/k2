package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"sync/atomic"

	"github.com/kingwel-xie/k2/core/logger"
)

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func Setup(db *gorm.DB) *casbin.SyncedEnforcer {
	Apter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		panic(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	e.SetLogger(&Logger{})
	e.EnableLog(true)
	return e
}


// Logger is the implementation for a Logger using golang log.
type Logger struct {
	enable int32
}

var log2 = logger.Logger("casbin")

// EnableLog controls whether print the message.
func (l *Logger) EnableLog(enable bool) {
	i := 0
	if enable {
		i = 1
	}
	atomic.StoreInt32(&(l.enable), int32(i))
}

// IsEnabled returns if logger is enabled.
func (l *Logger) IsEnabled() bool {
	return atomic.LoadInt32(&(l.enable)) != 0
}

// LogModel log info related to model.
func (l *Logger) LogModel(model [][]string) {
	if l.IsEnabled() {
		var str string
		for i := range model {
			for j := range model[i] {
				str += " " + model[i][j]
			}
			str += "\n"
		}
		log2.Info(str)
	}
}

// LogEnforce log info related to enforce.
func (l *Logger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	//logger.DefaultLogger.Fields(map[string]interface{}{
	//	"matcher":  matcher,
	//	"request":  request,
	//	"result":   result,
	//	"explains": explains,
	//}).Log(logger.InfoLevel, nil)
	if l.IsEnabled() {
		log2.Infow("LogEnforce", "matcher", matcher, "request", request, "result", result, "explains", explains)
	}
}

// LogRole log info related to role.
func (l *Logger) LogRole(roles []string) {
	//logger.DefaultLogger.Fields(map[string]interface{}{
	//	"roles": roles,
	//})

	if l.IsEnabled() {
		log2.Infow("LogRole", "role", roles)
	}
}

// LogPolicy log info related to policy.
func (l *Logger) LogPolicy(policy map[string][][]string) {
	if l.IsEnabled() {
		log2.Infow("LogPolicy", "policy", policy)
	}
}
