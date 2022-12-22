package console

import (
	"fmt"
	"os"
	"strings"
)

//读取数据
func Read(msg string) string {

	fmt.Printf("%s", msg)
	red := make([]byte, 1024)
	l, e := os.Stdin.Read(red)
	if e != nil {
		return ""
	}
	ret := string(red[:l])
	return strings.TrimSpace(ret)
}
func ReadDefault(msg, def string) string {
	data := Read(msg)
	if data == "" {
		return def
	}
	return data
}
