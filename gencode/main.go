package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eoe2005/goapp/gencode/code"
)

var cmds = map[string]code.Cmd{
	"gzipencode": {
		Help:   "gzip 压缩，base64编码",
		Handle: code.GzipEncodeBase64,
	},
	"gzipdecode": {
		Help:   "gzip 解压，base64编码",
		Handle: code.GzipDecodeBase64,
	},
}

func main() {
	if len(os.Args) < 3 {
		showHelp("参数错误")
	}
	cmd, ok := cmds[strings.ToLower(os.Args[1])]
	if ok {
		fmt.Printf("\033[1;36m输出 :\033[0m \t \033[7;35m%s\033[0m\n\n", cmd.Handle(os.Args[2]))

	} else {
		showHelp("参数不存在")
	}
}
func showHelp(msg string) {
	fmt.Println(msg)
	fmt.Printf(" \033[1;31m%s\033[0m \033[1;32m[cmd]\033[0m \033[1;33mdata \033[0m\n", os.Args[0])
	for c, val := range cmds {
		fmt.Printf("   \033[3;32m%s\033[0m  \033[1;30m%s\033[0m\n", c, val.Help)
	}
	fmt.Println("")
	os.Exit(0)
}
