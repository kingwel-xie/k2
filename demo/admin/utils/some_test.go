package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)


func TestFmt(t *testing.T) {
	t.Log(insertDictData("test", "fee", 5))
	t.Log(insertDictData("", "fee2", "*"))
}

func TestGenFeeTypesSQL(t *testing.T) {
	//filepath := "C:\\Users\\hello\\Desktop\\kobh\\费用说明.xlsx"
	filepath := "/Users/dian/Desktop/tpls/费用说明20220927.xlsx"
	file, err := excelize.OpenFile(filepath)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// Close the spreadsheet.
		file.Close()
	}()

	//cols, err := file.Cols("Sheet1")
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		t.Fatal(err)
	}
	group := ""
	value := 0
	t.Log("DELETE FROM sys_dict_data WHERE dict_type = 'tbx_fee_type';")

	t.Log(insertDictData("", "计重价", "*"))
	for _, row := range rows {
		if len(row[1]) > 0 {
			group = row[1]
		}
		t.Log(insertDictData(group, row[2], value))
		value += 1
	}
}

func insertDictData(group, feeName string, value interface{}) string {
	hyphen := "-"
	if len(group) == 0 || len(feeName) == 0 {
		hyphen = ""
	}
	return fmt.Sprintf("INSERT INTO sys_dict_data (dict_sort, dict_label, dict_value, dict_type, status, create_by, update_by) VALUES (0, '%s%s%s', '%v', 'tbx_fee_type', 2, 'admin', 'admin');", group, hyphen, feeName, value)
}
