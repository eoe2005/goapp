package serveragent

import (
	"io"
	"net"
	"os"
)

func publish(c net.Conn) {
	for !checkoutCmd(c, "testoo1", "token错误") {
	}
	fname := ""
	for fname == "" {
		fname = readData(c, "文件名字")
	}
	if !readFile(c, fname) {
		err(c, "保存文件失败")
	}

}

func readFile(con net.Conn, filepath string) bool {

	os.Remove(filepath)
	fd, e := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND, 0777)
	if e != nil {
		return false
	}
	defer fd.Close()
	for {
		max := 4096
		data := make([]byte, max)
		l, e := con.Read(data)
		if e == io.EOF {
			return false
		}
		fd.Write(data[0:l])
		if l < max {
			break
		}
	}

	return true
}
