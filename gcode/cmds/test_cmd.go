package cmds

import "fmt"

type TestCmd struct {
}

func (c *TestCmd) Help() string {
	return "这是一个测试功能"
}
func (c *TestCmd) Run() {
	fmt.Println("承担")
}
