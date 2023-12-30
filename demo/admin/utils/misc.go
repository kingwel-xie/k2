package utils

import (
	"encoding/json"
	"time"
)

func JsonMustMarshal(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / 1e6
}
