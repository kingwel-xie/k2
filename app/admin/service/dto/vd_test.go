package dto

import (
	"fmt"
	"testing"

	vd "github.com/bytedance/go-tagexpr/v2/validator"
)

func TestResolve(t *testing.T) {
	var update = TbxCountryUpdateReq{
		CName: "1",
		EName: "ename",
		EName2: "2",
		Code2: "2",
	}

	//vd.SetErrorFactory(func(failPath, msg string) error {
	//	return fmt.Errorf(`"validation failed: %s %s"`, failPath, msg)
	//})
	if err := vd.Validate(update); err != nil {
		fmt.Println(err)
	}
}

