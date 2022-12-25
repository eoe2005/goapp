package cmds

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//mysql 配置文件
type mysqlconf map[string]string

type MysqlCmd struct {
	conf mysqlconf
}

func (c *MysqlCmd) Help() string {
	return "mysql这是一个测试功能"
}
func (c *MysqlCmd) Run() {
	c.conf = mysqlconf{}
	e := loadConf("mysql", &c.conf)
	if e != nil {
		fmt.Println(e)
	}
	if len(c.conf) > 0 {
		fmt.Println("可选配置:\n")
		for n, conf := range c.conf {
			fmt.Printf("  %s : %s\n", n, conf)
		}
	}
	cname := readDef("\n输入配置名字[new:del]:", "new")
	switch cname {
	case "new":
		c.getCon("")
	case "del":
	default:
		condns, ok := c.conf[cname]
		if ok {
			c.getCon(condns)

		}
	}

}

func (c *MysqlCmd) getCon(conDns string) *sql.DB {
	if conDns != "" {
		dbcon, err := sql.Open("mysql", conDns)
		if err != nil {
			fmt.Println(err)
			panic("")

		}
		return dbcon
	}
	host := readDef("请输入服务器地址[默认: localhost] -> ", "127.0.0.1")
	port := readDef("请输入端口[默认: 3306] -> ", "3306")
	user := readDef("请输入账号[默认: root] -> ", "root")
	pass := readDef("请输入密码 -> ", "")
	dbcon, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port))
	if err != nil {
		fmt.Println(err)
		return c.getCon("")
	}
	rows, e := dbcon.Query("SHOW DATABASES")
	if e != nil {
		fmt.Println(e.Error())
		return nil
	}
	defer rows.Close()
	dbs := []string{}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		dbs = append(dbs, name)
	}

	dbname := ""
	isOk := true
	for isOk {
		dbname = read(fmt.Sprintf("选择数据[%s] -> ", strings.Join(dbs, ",")))
		if dbname == "" {
			continue
		}

		for _, name := range dbs {
			if name == dbname {
				isOk = false
				break
			}
		}
	}
	dbcon.Exec("USE " + dbname)
	cname := readDef("保存配置 -> ", "")
	if cname != "" {
		c.conf[cname] = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)
		c.save()
	}
	return dbcon
}
func (c *MysqlCmd) save() {
	saveConf("mysql", c.conf)
}
