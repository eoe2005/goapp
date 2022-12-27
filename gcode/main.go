package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eoe2005/goapp/gcode/cmds"
)

var cmdMap = map[string]cmds.CmdInterFace{
	"kafka": &cmds.KafkaCmd{},
	"redis": &cmds.RedisCmd{},
	"mysql": &cmds.MysqlCmd{},
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println(os.Args)
		help()
	}
	cmdName := strings.ToLower(os.Args[1])
	c, ok := cmdMap[cmdName]
	if ok {
		c.Run()
	} else {
		help()
	}
}
func help() {
	fmt.Println(" " + os.Args[0] + "  命令  参数列表\n\n 命令:\n")
	for name, s := range cmdMap {
		fmt.Printf("    %s    %s\n\n", name, s.Help())
	}
	os.Exit(-1)
}
