package backlog

import (
	//"admin/common/defs"
	"admin/models"
	"admin/utils"
	"github.com/kingwel-xie/k2/common"
	"github.com/kingwel-xie/k2/core/storage"
	"sync"
)

const (
	RefreshCache = "CacheRefresh"
)

var (
	lock sync.Mutex

	systemMap map[string]map[string]models.SysDictData

	userMap     map[int]models.SysUser
	usernameMap map[string]models.SysUser
	apiTokenMap map[string]models.SysUser

	countryMap map[string]models.TbxCountry
	regionMap  map[string]models.TbxCountry
	// TODO: more dicts
)

func InitCache() {
	queue := common.Runtime.Queue()
	queue.Register(RefreshCache, postRefreshCache)

	// force loading cache for the first time
	_ = postRefreshCache(nil)

	// start a cron job to reload cache periodically, every hour
	common.Runtime.GetCrontab().AddJob("cacheRefresher", "0 0 * * * ?", func() {
		_ = postRefreshCache(nil)
	})
}

func GetSysDict(category, code string) *models.SysDictData {
	lock.Lock()
	defer lock.Unlock()
	ch, ok := systemMap[category]
	if !ok {
		return nil
	} else {
		data, ok := ch[code]
		if !ok {
			return nil
		} else {
			return &data
		}
	}
}

func GetSysDicts(category string) map[string]models.SysDictData {
	lock.Lock()
	defer lock.Unlock()
	return systemMap[category]
}

func GetCountry(code string) *models.TbxCountry {
	lock.Lock()
	defer lock.Unlock()
	c, ok := countryMap[code]
	if !ok {
		return nil
	} else {
		return &c
	}
}

func GetCountryOrRegion(code string) *models.TbxCountry {
	lock.Lock()
	defer lock.Unlock()
	c, ok := countryMap[code]
	if ok {
		return &c
	}
	c, ok = regionMap[code]
	if !ok {
		return nil
	} else {
		return &c
	}
}

func Username(userId int) string {
	user := GetSysUser(userId)
	if user == nil {
		return ""
	}
	return user.Username
}

func GetSysUser(userId int) *models.SysUser {
	lock.Lock()
	defer lock.Unlock()
	c, ok := userMap[userId]
	if !ok {
		return nil
	} else {
		return &c
	}
}

func GetSysUserByToken(token string) *models.SysUser {
	if !utils.CompareToken(token) {
		return nil
	}

	lock.Lock()
	defer lock.Unlock()
	user, ok := apiTokenMap[token]
	if !ok || user.Status != "2" {
		return nil
	} else {
		return &user
	}
}

func loadSystemDict() []models.SysDictData {
	var list = make([]models.SysDictData, 0)
	err := common.Runtime.GetDb().
		//Where("status = '2'").
		Find(&list).Error

	log.Infof("loading - system-dict, err=%v", err)

	return list
}

func loadSysUser() []models.SysUser {
	var list = make([]models.SysUser, 0)
	err := common.Runtime.GetDb().Preload("Dept").Preload("Role").
		Find(&list).Error

	log.Infof("loading - sysuser, err=%v", err)
	return list
}

func loadCountry() []models.TbxCountry {
	var list = make([]models.TbxCountry, 0)
	err := common.Runtime.GetDb().
		Order("display_sort").
		Find(&list).Error

	// here we format country and its regions
	// first, get country list
	var list2 = make([]models.TbxCountry, 0)
	for _, c := range list {
		// if it is a country
		if c.BelongTo == "" {
			// try to build its children regions
			for _, r := range list {
				if r.BelongTo == c.Code {
					c.Children = append(c.Children, r)
				}
			}
			list2 = append(list2, c)
		}
	}

	log.Infof("loading - country, err=%v", err)
	return list2
}

func LoadAllDict() map[string]interface{} {
	var data = make(map[string]interface{})

	data["systemList"] = loadSystemDict()
	data["userList"] = loadSysUser()
	data["countryList"] = loadCountry()

	return data
}

func postRefreshCache(_ storage.Messager) error {
	// load from DB
	data := LoadAllDict()

	lock.Lock()

	systemMap = make(map[string]map[string]models.SysDictData)
	for _, c := range data["systemList"].([]models.SysDictData) {
		_, ok := systemMap[c.DictType]
		if !ok {
			systemMap[c.DictType] = make(map[string]models.SysDictData)
		}
		systemMap[c.DictType][c.DictValue] = c
	}

	// user map & token map
	userMap = make(map[int]models.SysUser)
	usernameMap = make(map[string]models.SysUser)
	apiTokenMap = make(map[string]models.SysUser)
	for _, c := range data["userList"].([]models.SysUser) {
		userMap[c.UserId] = c
		usernameMap[c.Username] = c
		if len(c.Token) > 0 {
			apiTokenMap[c.Token] = c
		}
	}

	// country map
	countryMap = make(map[string]models.TbxCountry)
	regionMap = make(map[string]models.TbxCountry)
	for _, c := range data["countryList"].([]models.TbxCountry) {
		countryMap[c.Code] = c
		for _, r := range c.Children {
			regionMap[r.Code] = r
		}
	}

	lock.Unlock()
	return nil
}

func ReloadCacheAsync() {
	message, err := common.Runtime.GetStreamMessage("", RefreshCache, nil)
	if err != nil {
		// quite unlikely
		log.Errorf("GetStreamMessage error, %s", err.Error())
	} else {
		err = common.Runtime.Queue().Append(message)
		if err != nil {
			log.Errorf("Append message error, %s", err.Error())
		}
	}
}

func SetupCacheForTest() {
	_ = postRefreshCache(nil)
}
