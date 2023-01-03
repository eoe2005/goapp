package main

import (
	"fmt"
	"goapps/utils"
	"os"
	"os/exec"
	"path"
)

type Service interface {
	Add(filepath string)
	Del(filepath string)
}
type AlpineService struct {
}
type UbuntuService struct {
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("参数错误 %s add|del exepath", os.Args[0])
		os.Exit(-1)
	}
	act := os.Args[1]
	exeFilePath := os.Args[2]
	var s Service
	switch utils.GetOs() {
	case utils.ALPINE:
		s = AlpineService{}
		break
	case utils.UBUNTU:

	}
	fmt.Println(utils.GetOs())
	if s == nil {
		panic("不支持的操作系统")
	}
	switch act {
	case "add":
		s.Add(exeFilePath)
	case "del":
		s.Del(exeFilePath)
	default:
		panic("不支持的命令")

	}
}

func (s AlpineService) Add(filepath string) {
	sname := path.Base(filepath)
	fcontent := `#!/sbin/openrc-run
name="busybox $SVCNAME"
command="%s"
command_args=""
command_background="yes"
pidfile="/run/$RC_SVCNAME.pid"
depend() {
	need net localmount
	after firewall
}
`
	f, e := os.OpenFile("/etc/init.d/"+sname, os.O_CREATE|os.O_WRONLY, 0766)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf(fcontent, filepath))
	exec.Command("rc-update", "add", sname)
}
func (s AlpineService) Del(filepath string) {
	sname := path.Base(filepath)
	exec.Command("rc-update", "add", sname)
	os.Remove("/etc/init.d/" + sname)
}

func (s UbuntuService) Add(filepath string) {
	sname := path.Base(filepath)
	fcontent := `[Unit]
	Description=%s
	After=network.target
	After=syslog.target
	
	[Service]
	Type=simple
	LimitNOFILE=65535
	ExecStart=%s
	ExecReload=/bin/kill -USR1 $MAINPID
	RestartSec=5
	Restart=always
	
	[Install]
	WantedBy=multi-user.target
`
	f, e := os.OpenFile("/etc/systemd/system/"+sname+".service", os.O_CREATE|os.O_WRONLY, 0766)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf(fcontent, sname, filepath))
	exec.Command("systemctl", "reload-daemon")
	exec.Command("systemctl", "enable", sname)
	exec.Command("systemctl", sname, "start")

}
func (s UbuntuService) Del(filepath string) {
	sname := path.Base(filepath)
	exec.Command("systemctl", "disable", sname)
	exec.Command("systemctl", sname, "stop")
	os.Remove("/etc/systemd/system/" + sname + ".service")
	exec.Command("systemctl", "reload-daemon")
}
