package gen

type SysTables struct {
	PackageName  string `json:"packageName"`  // 包名 admin
	ModuleName   string `json:"moduleName"`   // 模塊名 country
	ClassName    string `json:"className"`    // 类名 TbxCountry
	TableComment string `json:"tableComment"` // 表备注

	TBName          string `json:"tableName"`       // 表名称 tbx_country，文件名, 缺省由 ClassName 生成
	ModuleFrontName string `json:"moduleFrontName"` // tbx-country, 缺省由 ClassName 生成
	BusinessName    string `json:"businessName"`    // tbxCountry, 缺省由 ClassName 生成

	DataScope bool         `json:"dataScope"` // 数据权限 是否支持数据权限
	IsAuth    bool         `json:"isAuth"`    // 路由 是否带认证
	Columns   []SysColumns `json:"columns"`   // 列

	HardDelete bool `json:"hardDelete"` // model, ModelTimeHardDelete
	HasExport  bool `json:"hasExport"`  // vue, Export 按钮

	// 以下字段提取自Columns
	PkColumn    string `json:"-"` // 主键数据库模型列名
	PkComment   string `json:"-"` // 主键 注释
	PkGoField   string `json:"-"` // 主键 go 字段名
	PkGoType    string `json:"-"` // 主键 go 类型
	PkJsonField string `json:"-"` // 主键 json 字段名
}

type SysColumns struct {
	GoField       string `json:"goField"`       // go 字段名，由此生成 ColumnName+JsonField
	GoType        string `json:"goType"`        // go 数据类型, string/int/...
	ColumnComment string `json:"columnComment"` // 列字段注释，简短描述，一般不超过5个汉字
	Comment 	  string `json:"comment"` 		// 注释，不影响生成代码
	ColumnName    string `json:"columnName"`    // gorm 数据库字段名, 缺省由GoField 生成
	JsonField     string `json:"jsonField"`     // json 字段名, 缺省由GoField 生成
	GormTag       string `json:"gormTag"`       // gorm tag, 类型+约束, f.g., size:64;unique;index;uniqueIndex;not null;default;
	Validator     string `json:"validator"`     // validator
	IsPk          bool   `json:"isPk"`          // 主鍵?
	Required      bool   `json:"required"`      // dto, 是否必选字段, gin validator binding:required
	AutoInc       bool   `json:"autoInc"`       // 自增？
	Queryable     bool   `json:"queryable"`     // dto, 可查询的 vue GetPageReq 参数
	QueryType     string `json:"queryType"`     // dto, EQ/NE/LIKE/GT/GTE/LT/LTE  => SQL exact/iexact/contains/gt/gte/lt/lte, 缺省 EQ
	Sortable      bool   `json:"sortable"`      // dto, 可以排序 vue GetPageReq 参数
	NotOnUpdate   bool   `json:"notOnUpdate"`   // dto, 可以更新 UpdateReq 参数, vue Update dialog
	NotOnInsert   bool   `json:"notOnInsert"`   // dto, 不出现于 InsertReq, vue 不出现于 insert dialog
	NotOnList     bool   `json:"notOnList"`     // vue, 列表不显示
	HtmlType      string `json:"htmlType"`      // vue, datetime/file/radio/select/input/textarea/switch, 缺省 input
	DictType      string `json:"dictType"`      // vue, 字典類型
	//Sort               int          `gorm:"column:sort;" json:"sort"`								// ?
	FkClassName       string `json:"fkClassName"`       // vue, 外键类名 TbxCountry, 由此生成 FkTableName+FkModuleFrontName
	FkTableName       string `json:"fkTableName"`       // vue, 外键表名 tbx_country, 缺省由FkClassName生成
	FkModuleFrontName string `json:"fkModuleFrontName"` // vue, 外键前端模块名 tbx-country, 缺省由FkClassName生成
	FkLabelId         string `json:"fkLabelId"`         // vue, 外键id json 字段名, f.g., id
	FkLabelName       string `json:"fkLabelName"`       // vue, 外键name json 字段名, f.g., name
}
