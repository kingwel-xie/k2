package utils

import (
	"fmt"
	"math"
	"strconv"
)

func IsEqual(f1, f2 float64, drift float64) bool {
	if f1 > f2 {
		return math.Dim(f1, f2) < drift
	}else{
		return math.Dim(f2, f1) < drift
	}
}

func FormatFloat(num float64, argv...int) float64 {
	decimal := 2
	if len(argv) > 0 {
		decimal = argv[0]
	}
	format := "%." + strconv.Itoa(decimal) + "f"

	v, _ := strconv.ParseFloat(fmt.Sprintf(format, num), 64)
	return v
}
