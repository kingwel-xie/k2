package service


const (
	RefreshCache = "CacheRefresh"
)

var (
	lock sync.Mutex

	systemMap map[string]map[string]models.SysDictData

	countryMap map[string]models.TbxCountry
	// TODO: more dicts
)

func InitCache() {
	queue := common.Runtime.Queue()
	queue.Register(RefreshCache, postRefreshCache)

	// force loading cache for the first time
	_ = postRefreshCache(nil)
}

func loadSystemDict() []models.SysDictData {
	var list = make([]models.SysDictData, 0)
	err := common.Runtime.GetDb().
		Where("status == '2'").
		Find(&list).Error

	log.Infof("loading - system-dict, err=%v", err)

	return list
}

func loadCountry() []models.TbxCountry {
	var list = make([]models.TbxCountry, 0)
	err := common.Runtime.GetDb().
		Find(&list).Error

	log.Infof("loading - country, err=%v", err)
	return list
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
	postRefreshCache(nil)
}

// Validation functions
func ValidateCountryCode(code string) error {
	if GetCountry(code) == nil {
		return defs.ErrInvalidCountryCode.Wrapf(code)
	}
	return nil
}

func ValidateSysDict(category, code string) error {
	if GetSysDict(category, code) == nil {
		return defs.ErrInvalidSystemDict.Wrapf("type=%s code=%s", category, code)
	}
	return nil
}

