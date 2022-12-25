package cmds

import "fmt"

type RedisCmd struct {
}

func (c *RedisCmd) Help() string {
	return "redis这是一个测试功能"
}
func (c *RedisCmd) Run() {
	fmt.Println("承担")
}
