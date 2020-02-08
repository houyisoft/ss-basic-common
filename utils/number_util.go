package utils

import (
	"fmt"
	"ss-basic-common/utils/cast"
	"strconv"
)

//算百分比
func GetRate(a int, b int) string {
	if a == 0 || b == 0 {
		return "0"
	}
	n := (cast.ToFloat64(a) / cast.ToFloat64(b)) * 100
	n2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", n), 64)
	return cast.ToString(n2)
}

func GetAvageRate(a string, b string) string {
	n := (cast.ToFloat64(a) + cast.ToFloat64(b)) / 2
	n2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", n), 64)
	return cast.ToString(n2)
}
