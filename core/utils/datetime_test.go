package utils

import "testing"

func TestJsonMarshal(t *testing.T) {
	var time = `"10:47:01"`

	var jst JSONTime
	err := jst.UnmarshalJSON([]byte(time))
	if err != nil {
		t.Fatal(err)
	}

}
