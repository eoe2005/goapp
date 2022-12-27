package cmds

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

type kafkaconf struct {
	Topic string   `json:"topic"`
	Host  []string `json:"host"`
}
type kafkamap map[string]kafkaconf
type KafkaCmd struct {
	confMap kafkamap
}

func (c *KafkaCmd) Help() string {
	return "kafka 测试工具"
}
func (c *KafkaCmd) init() {
	c.confMap = kafkamap{}
	loadConf("kafka", &c.confMap)
}
func (c *KafkaCmd) Run() {
	c.init()
	for {
		cmd := readDef("请输入命令:[publish|consumer|conf]", "conf")
		if cmd == "publish" {
			c.write()
		} else if cmd == "consumer" {
			c.read()
		} else if cmd == "conf" {
			c.readConf()
		}
	}
}
func (c *KafkaCmd) readConf() kafkaconf {
	for {
		fmt.Println("配置列表")
		for k, v := range c.confMap {
			fmt.Printf("%s [%s]%s\n", k, v.Topic, strings.Join(v.Host, ","))
		}
		cmd := readDef("选择配置[默认新建]:", "")
		if cmd == "" {
			topic := mustread("请输入TOPIC:")
			hosts := mustread("请输入地址：")
			name := mustread("请输入配置名字：")
			hs := strings.Split(hosts, ",")
			if len(hs) == 0 || hs[0] == "" {
				continue
			}
			ret := kafkaconf{
				Topic: topic,
				Host:  hs,
			}
			w := &kafka.Writer{
				Topic:                  topic,
				Addr:                   kafka.TCP(hs...),
				AllowAutoTopicCreation: true,
			}
			e := w.Close()
			if e != nil {
				fmt.Println("配置失败")
				continue
			}
			c.confMap[name] = ret
			saveConf("kafka", c.confMap)
			return ret

		} else {
			r, ok := c.confMap[cmd]
			if ok {
				return r
			}
		}
	}
	return kafkaconf{}
}
func (c *KafkaCmd) read() {
	conf := c.readConf()
	ctx := context.Background()
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        conf.Host,
		GroupID:        conf.Topic + fmt.Sprintf("%d", time.Now().UnixNano()),
		Topic:          conf.Topic,
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // flushes commits to Kafka every second
		StartOffset:    kafka.LastOffset,
	})
	defer r.Close()
	for {
		msg, e := r.FetchMessage(ctx)
		if e != nil {
			fmt.Printf("读取失败 %s\n", e.Error())
			return
		}
		fmt.Printf("[%s] %s -> %s\n", msg.Topic, string(msg.Key), string(msg.Value))
	}
}

func (c *KafkaCmd) write() {
	conf := c.readConf()
	ctx := context.Background()
	w := &kafka.Writer{
		Topic:    conf.Topic,
		Addr:     kafka.TCP(conf.Host...),
		Balancer: &kafka.LeastBytes{},
	}
	defer w.Close()
	for {
		k := readDef("输入key[exit退出]:", "")
		if k == "exit" {
			return
		}
		val := ""
		for val == "" {
			val = readDef("输入消息内容:", "")
		}
		e := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(k),
			Value: []byte(val),
		})
		if e == nil {
			continue
		}
		fmt.Printf("发送失败 %s\n", e.Error())
	}

}
