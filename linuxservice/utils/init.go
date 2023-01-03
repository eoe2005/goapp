package utils

import (
	"os/exec"
	"runtime"
	"strings"
	"unsafe"
)

const (
	UNKNOW = 0
	WINDOW = 1
	MAC    = 2
	UBUNTU = 3
	ALPINE = 4
)

func Str2Byte(in string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{in, len(in)},
	))
}
func Bytes2String(in []byte) string {
	return *(*string)(unsafe.Pointer(&in))
}
func GetOs() int {
	switch runtime.GOOS {
	case "Window":
		return WINDOW
	case "Darwin":
		return MAC
	case "linux":
		c := exec.Command("uname", "-a")
		dd, _ := c.Output()
		outStr := Bytes2String(dd)
		if strings.Index(outStr, "Alpine") >= 0 {
			return ALPINE
		} else if strings.Index(outStr, "Ubuntu") >= 0 {
			return UBUNTU
		} else if strings.Index(outStr, "Darwin") >= 0 {
			return MAC
		}
	}

	return UNKNOW
}
