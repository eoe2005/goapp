package cmds

import "fmt"

type KafkaCmd struct {
}

func (c *KafkaCmd) Help() string {
	return "kafka这是一个测试功能"
}
func (c *KafkaCmd) Run() {
	fmt.Println("承担22")
}
