package code

import (
	"strconv"
	"time"
)

func TimeUnixToString(src string) string {
	d, _ := strconv.ParseInt(src, 10, 64)
	return time.Unix(d, 0).Format("2006-01-2 15:04:05")
}
func TimeUnixMilliString(src string) string {
	d, _ := strconv.ParseInt(src, 10, 64)
	return time.UnixMilli(d).Format("2006-01-2 15:04:05.000")
}
