package cmds

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

//加载配置文件
func loadConf(file string, obj any) error {
	file = getFilePath(file)
	if file == "" {
		return errors.New("文件不存在")
	}
	data, e := ioutil.ReadFile(file)
	if e != nil {
		return e
	}
	return json.Unmarshal(data, obj)
}

//保存配置文件
func saveConf(file string, obj any) error {
	file = getFilePath(file)
	if file == "" {
		return errors.New("文件不存在")
	}
	data, e := json.Marshal(obj)
	if e != nil {
		return e
	}
	return ioutil.WriteFile(file, data, 0777)
}
func getFilePath(file string) string {
	dir, e := os.UserHomeDir()
	if e != nil {
		fmt.Println(e)
		return ""
	}
	dir = dir + "/.goapps/gcode"
	e = os.MkdirAll(dir, 0777)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	return dir + "/" + file + ".conf"
}
