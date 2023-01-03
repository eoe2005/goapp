package serveragent

import "net"

func ok(con net.Conn) {
	con.Write([]byte("done"))
}
func err(con net.Conn, msg string) {
	con.Write([]byte(msg))
}
func readData(con net.Conn, errmsg string) string {
	head := make([]byte, 1024)
	l, e := con.Read(head)
	if e != nil {
		err(con, errmsg)
		return ""
	}
	ok(con)
	return string(head[:l])
}
func checkoutCmd(con net.Conn, name, errmsg string) bool {
	if readData(con, errmsg) != name {
		con.Write([]byte(errmsg))
		return false
	}
	return true
}

func Agent(con net.Conn) {
	defer con.Close()
	head := make([]byte, 1024)
	l, e := con.Read(head)
	if e != nil {
		return
	}
	cmd := string(head[:l])
	switch cmd {
	case "pub_go":
		ok(con)
		publish(con)
	}

}
