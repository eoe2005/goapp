package main

import (
	"flag"

	"github.com/eoe2005/goapp/gitwebhook/logic"
	"github.com/gin-gonic/gin"
)

var (
	giteetoken = ""
	serverPort = "9999"
)

func main() {
	flag.StringVar(&giteetoken, "gitee", "", "gitee token")
	flag.StringVar(&serverPort, "server_port", "9999", "server port")
	flag.Parse()
	logic.SetGiteeToken(giteetoken)
	run()

}

func run() {
	ge := gin.New()
	ge.POST("/gitee", func(c *gin.Context) {
		p := logic.GiteeParam{}
		e := c.Bind(&p)
		if e == nil {
			c.JSON(200, gin.H{"code": 0})
			return
		}
		go p.Run()
		c.JSON(200, gin.H{"code": 0})

	})
	ge.Run(":" + serverPort)
}
