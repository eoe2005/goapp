package cmds

import (
	"fmt"
	"os"
	"strings"
)

type CmdInterFace interface {
	Help() string
	Run()
}

func read(msg string) string {

	fmt.Printf("%s", msg)
	red := make([]byte, 1024)
	l, e := os.Stdin.Read(red)
	if e != nil {
		return ""
	}
	ret := string(red[:l])
	return strings.TrimSpace(ret)
}
func readDef(msg, def string) string {
	data := read(msg)
	if data == "" {
		return def
	}
	return data
}
func mustread(msg string) string {
	ret := ""
	for ret == "" {
		ret = read(msg)
	}
	return ret
}
