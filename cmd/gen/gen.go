package gen

import (
	"bytes"
	"fmt"
	k2template "github.com/kingwel-xie/k2/template"
	"strconv"
	"strings"
	"text/template"
	"time"

	"gorm.io/gorm/schema"

	"github.com/kingwel-xie/k2/core/utils"
)

var (
	FrontPath string = "front"
)

type Gen struct {
}

func toLowerFirstChar(s string) string {
	if len(s) == 0 {
		return s
	}
	var f1 = s[:1]
	var f2 = s[1:]

	return strings.ToLower(f1) + f2
}

func preprocessTable(tab *SysTables) error {

	var ns = schema.NamingStrategy{
		TablePrefix:   "",
		SingularTable: true,
		NameReplacer:  strings.NewReplacer("CID", "Cid"),
	}

	// these fields must be specified
	if len(tab.PackageName) == 0 || len(tab.ModuleName) == 0 || len(tab.ClassName) == 0 {
		return fmt.Errorf("PackageName/ModuleName/ClassName not specified")
	}

	// set TableName if not specified
	if len(tab.TBName) == 0 {
		tab.TBName = ns.TableName(tab.ClassName)
	}
	// set ModuleFrontName if not specified
	if len(tab.ModuleFrontName) == 0 {
		s := ns.TableName(tab.ClassName)
		tab.ModuleFrontName = strings.Replace(s, "_", "-", -1)
	}
	if len(tab.BusinessName) == 0 {
		tab.BusinessName = toLowerFirstChar(tab.ClassName)
	}

	for index, v := range tab.Columns {
		if len(v.GoField) == 0 {
			return fmt.Errorf("column %d no GoField specified", index)
		}
		if len(v.GoType) == 0 {
			return fmt.Errorf("column %d no GoType specified", index)
		}

		// set ColumnName if not specified
		if len(v.ColumnName) == 0 {
			tab.Columns[index].ColumnName = ns.ColumnName("", v.GoField)
		}
		// set JsonField if not specified
		if len(v.JsonField) == 0 {
			tab.Columns[index].JsonField = toLowerFirstChar(v.GoField)
		}
		// set HtmlType if not specified
		if len(v.HtmlType) == 0 {
			if strings.Contains(v.GoType, "time.Time") {
				tab.Columns[index].HtmlType = "datetime"
			} else {
				tab.Columns[index].HtmlType = "input"
			}
		}
		// set QueryType to EQ if not specified
		if len(v.QueryType) == 0 {
			tab.Columns[index].QueryType = "EQ"
		}

		// set FkTableName if not specified
		if len(v.FkTableName) == 0 {
			tab.Columns[index].FkTableName = ns.TableName(v.FkClassName)
		}
		// set FkModuleFrontName if not specified
		if len(v.FkModuleFrontName) == 0 {
			s := ns.TableName(v.FkClassName)
			tab.Columns[index].FkModuleFrontName = strings.Replace(s, "_", "-", -1)
		}

		// in the end, extract the PK information
		if v.IsPk {
			tab.PkGoField = tab.Columns[index].GoField
			tab.PkComment = tab.Columns[index].ColumnComment
			tab.PkGoType = tab.Columns[index].GoType
			tab.PkColumn = tab.Columns[index].ColumnName
			tab.PkJsonField = tab.Columns[index].ColumnName
		}
	}

	return nil
}

func (e Gen) GenCode(tab *SysTables) {

	if err := preprocessTable(tab); err != nil {
		fmt.Println(err)
		return
	}

	basePath := "v4/"

	routerFile := basePath
	if tab.IsAuth {
		routerFile += "router_check_role.go.template"
	} else {
		routerFile += "router_no_check_role.go.template"
	}

	t1, err := parseByName(basePath + "model.go.template")
	if err != nil {
		fmt.Println(err)
		return
	}
	t2, err := parseByName(basePath + "apis.go.template")
	if err != nil {
		fmt.Println(err)
		return
	}
	t3, err := parseByName(routerFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	t4, err := parseByName(basePath + "js.go.template")
	if err != nil {
		fmt.Println(err)
		return
	}
	t5, err := parseByName(basePath + "vue.go.template")
	if err != nil {
		fmt.Println(err)
		return
	}
	t6, err := parseByName(basePath + "dto.go.template")
	if err != nil {
		fmt.Println(err)
		return
	}
	t7, err := parseByName(basePath + "service.go.template")
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = utils.PathCreate(tab.PackageName + "/apis/")
	_ = utils.PathCreate(tab.PackageName + "/models/")
	_ = utils.PathCreate(tab.PackageName + "/router/")
	_ = utils.PathCreate(tab.PackageName + "/service/dto/")
	_ = utils.PathCreate(FrontPath + "/api/")
	_ = utils.PathCreate(FrontPath + "/views/" + tab.ModuleFrontName)

	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	err = t4.Execute(&b4, tab)
	var b5 bytes.Buffer
	err = t5.Execute(&b5, tab)
	var b6 bytes.Buffer
	err = t6.Execute(&b6, tab)
	var b7 bytes.Buffer
	err = t7.Execute(&b7, tab)
	utils.FileCreate(b1, tab.PackageName+"/models/"+tab.TBName+".go")
	utils.FileCreate(b2, tab.PackageName+"/apis/"+tab.TBName+".go")
	utils.FileCreate(b3, tab.PackageName+"/router/"+tab.TBName+".go")
	utils.FileCreate(b4, FrontPath+"/api/"+tab.ModuleFrontName+".js")
	utils.FileCreate(b5, FrontPath+"/views/"+tab.ModuleFrontName+"/index.vue")
	utils.FileCreate(b6, tab.PackageName+"/service/dto/"+tab.TBName+".go")
	utils.FileCreate(b7, tab.PackageName+"/service/"+tab.TBName+".go")

	fmt.Println("Code generated successfully！")
}

func (e Gen) GenApiToFile(tab *SysTables) {
	t1, err := parseByName("api_migrate.template")
	if err != nil {
		fmt.Println(err)
		return
	}
	i := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	var b1 bytes.Buffer
	err = t1.Execute(&b1, struct {
		SysTables
		GenerateTime string
	}{*tab, i})

	utils.FileCreate(b1, "migrate/version-local/"+i+"_migrate.go")

	fmt.Println("Code generated successfully！")
}

//
//func (e Gen) ActionsGen(c *gin.Context, tab tools.SysTables) {
//
//	basePath := "template/v4/"
//	routerFile := basePath + "actions/router_check_role.go.template"
//
//	if tab.IsAuth == 2 {
//		routerFile = basePath + "actions/router_no_check_role.go.template"
//	}
//
//	t1, err := template.ParseFiles(basePath + "model.go.template")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	t3, err := template.ParseFiles(routerFile)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	t4, err := template.ParseFiles(basePath + "js.go.template")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	t5, err := template.ParseFiles(basePath + "vue.go.template")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	t6, err := template.ParseFiles(basePath + "dto.go.template")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	_ = utils.PathCreate("./app/" + tab.PackageName + "/models/")
//	_ = utils.PathCreate("./app/" + tab.PackageName + "/router/")
//	_ = utils.PathCreate("./app/" + tab.PackageName + "/service/dto/")
//	_ = utils.PathCreate(config.GenConfig.FrontPath + "/api/")
//	_ = utils.PathCreate(config.GenConfig.FrontPath + "/views/" + tab.ModuleFrontName)
//
//	var b1 bytes.Buffer
//	err = t1.Execute(&b1, tab)
//	var b3 bytes.Buffer
//	err = t3.Execute(&b3, tab)
//	var b4 bytes.Buffer
//	err = t4.Execute(&b4, tab)
//	var b5 bytes.Buffer
//	err = t5.Execute(&b5, tab)
//	var b6 bytes.Buffer
//	err = t6.Execute(&b6, tab)
//
//	utils.FileCreate(b1, "./app/"+tab.PackageName+"/models/"+tab.TBName+".go")
//	utils.FileCreate(b3, "./app/"+tab.PackageName+"/router/"+tab.TBName+".go")
//	utils.FileCreate(b4, config.GenConfig.FrontPath+"/api/"+tab.ModuleFrontName+".js")
//	utils.FileCreate(b5, config.GenConfig.FrontPath+"/views/"+tab.ModuleFrontName+"/index.vue")
//	utils.FileCreate(b6, "./app/"+tab.PackageName+"/service/dto/"+tab.TBName+".go")
//}
//

//func (e Gen) GenMenuAndApi(c *gin.Context) {
//	s:=service.SysMenu{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		fmt.Println(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	table := tools.SysTables{}
//	id, err := utils.StringToInt(c.Param("tableId"))
//	utils.HasError(err, "", -1)
//
//
//	table.TableId = id
//	tab, _ := table.Get(e.Orm)
//
//	Mmenu := dto.SysMenuControl{}
//	Mmenu.Title = tab.TableComment
//	Mmenu.Icon = "pass"
//	Mmenu.Path = "/" + strings.Replace(tab.TBName, "_", "-", -1)
//	Mmenu.MenuType = "M"
//	Mmenu.Action = "无"
//	Mmenu.ParentId = 0
//	Mmenu.NoCache = false
//	Mmenu.Component = "Layout"
//	Mmenu.Sort = 0
//	Mmenu.Visible = "0"
//	Mmenu.IsFrame = "0"
//	Mmenu.CreateBy = 1
//	s.Insert(&Mmenu)
//
//	Cmenu := dto.SysMenuControl{}
//	Cmenu.MenuName = tab.ClassName + "Manage"
//	Cmenu.Title = tab.TableComment
//	Cmenu.Icon = "pass"
//	Cmenu.Path = tab.TBName
//	Cmenu.MenuType = "C"
//	Cmenu.Action = "无"
//	Cmenu.Permission = tab.PackageName + ":" + tab.ModuleFrontName + ":list"
//	Cmenu.ParentId = Mmenu.MenuId
//	Cmenu.NoCache = false
//	Cmenu.Component = "/" + tab.ModuleFrontName + "/index"
//	Cmenu.Sort = 0
//	Cmenu.Visible = "0"
//	Cmenu.IsFrame = "0"
//	Cmenu.CreateBy = 1
//	Cmenu.UpdateBy = 1
//	s.Insert(&Cmenu)
//
//	MList := dto.SysMenuControl{}
//	MList.MenuName = ""
//	MList.Title = "分页获取" + tab.TableComment
//	MList.Icon = ""
//	MList.Path = tab.TBName
//	MList.MenuType = "F"
//	MList.Action = "无"
//	MList.Permission = tab.PackageName + ":" + tab.ModuleFrontName + ":query"
//	MList.ParentId = Cmenu.MenuId
//	MList.NoCache = false
//	MList.Sort = 0
//	MList.Visible = "0"
//	MList.IsFrame = "0"
//	MList.CreateBy = 1
//	MList.UpdateBy = 1
//	s.Insert(&MList)
//
//	MCreate := dto.SysMenuControl{}
//	MCreate.MenuName = ""
//	MCreate.Title = "创建" + tab.TableComment
//	MCreate.Icon = ""
//	MCreate.Path = tab.TBName
//	MCreate.MenuType = "F"
//	MCreate.Action = "无"
//	MCreate.Permission = tab.PackageName + ":" + tab.ModuleFrontName + ":add"
//	MCreate.ParentId = Cmenu.MenuId
//	MCreate.NoCache = false
//	MCreate.Sort = 0
//	MCreate.Visible = "0"
//	MCreate.IsFrame = "0"
//	MCreate.CreateBy = 1
//	MCreate.UpdateBy = 1
//	s.Insert(&MCreate)
//
//	MUpdate := dto.SysMenuControl{}
//	MUpdate.MenuName = ""
//	MUpdate.Title = "修改" + tab.TableComment
//	MUpdate.Icon = ""
//	MUpdate.Path = tab.TBName
//	MUpdate.MenuType = "F"
//	MUpdate.Action = "无"
//	MUpdate.Permission = tab.PackageName + ":" + tab.ModuleFrontName + ":edit"
//	MUpdate.ParentId = Cmenu.MenuId
//	MUpdate.NoCache = false
//	MUpdate.Sort = 0
//	MUpdate.Visible = "0"
//	MUpdate.IsFrame = "0"
//	MUpdate.CreateBy = 1
//	MUpdate.UpdateBy = 1
//	s.Insert(&MUpdate)
//
//	MDelete := dto.SysMenuControl{}
//	MDelete.MenuName = ""
//	MDelete.Title = "删除" + tab.TableComment
//	MDelete.Icon = ""
//	MDelete.Path = tab.TBName
//	MDelete.MenuType = "F"
//	MDelete.Action = "无"
//	MDelete.Permission = tab.PackageName + ":" + tab.ModuleFrontName + ":remove"
//	MDelete.ParentId = Cmenu.MenuId
//	MDelete.NoCache = false
//	MDelete.Sort = 0
//	MDelete.Visible = "0"
//	MDelete.IsFrame = "0"
//	MDelete.CreateBy = 1
//	MDelete.UpdateBy = 1
//	s.Insert(&MDelete)
//
//
//	fmt.Println("数据生成成功！")
//}

func parseByName(name string) (*template.Template, error) {
	p, err := k2template.Asset(name)
	if err != nil {
		return nil, err
	}
	return template.New(name).Parse(string(p))
}
