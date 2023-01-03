package main

import (
	"net"

	"github.com/eoe2005/goapp/serveragent/serveragent"
)

const PORT = ":9876"

func main() {
	s, e := net.Listen("tcp", PORT)
	if e != nil {
		panic("监听端口失败" + e.Error())
	}
	for {
		con, ce := s.Accept()
		if ce == nil {
			serveragent.Agent(con)
		}
	}
}
