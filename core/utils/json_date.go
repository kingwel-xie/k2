package utils

import (
	"fmt"
	"time"
)

type JSONDate struct {
	time.Time
}

func (d JSONDate) MarshalJSON() ([]byte, error) {
	if (d == JSONDate{}) {
		formatted := fmt.Sprintf("\"%s\"", "")
		return []byte(formatted), nil
	} else {
		return []byte(fmt.Sprintf(`"%s"`, d.Format("2006-01-02"))), nil
	}
}

func (d *JSONDate) UnmarshalJSON(b []byte) error {
	var err error
	// 指定时区
	d.Time, err = time.ParseInLocation(`"2006-01-02"`, string(b), time.Local)
	if err != nil {
		return err
	}
	return nil
}
