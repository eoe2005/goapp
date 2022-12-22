package db

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type finfo struct {
	Name     string
	TypeName string
	IsNull   bool
	Default  string
	Comment  string
}
type tableInfo struct {
	Name    string
	Comment string
	Pkg     map[string]int
	Fields  []finfo
}

func Test(sql string) {
	buildGo(sql)
}
func buildGo(sql string) {
	tinfo := tableInfo{
		Pkg:    map[string]int{},
		Fields: []finfo{},
	}
	lines := strings.Split(sql, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.TrimRight(line, ",")
		line = strings.ToLower(line)
		if strings.HasPrefix(line, "create") {
			sp := strings.Split(line, "`")
			tinfo.Name = sp[1]
			continue

		} else if strings.HasPrefix(line, ")") {
			sp := strings.Split(line, "='")
			if len(sp) == 2 {
				aa := strings.Split(sp[1], "'")
				tinfo.Comment = aa[0]
			}
			continue
		}

		if strings.Index(line, "key ") >= 0 {
			continue
		}

		sp := strings.Split(line, "`")
		f := finfo{
			Name:   sp[1],
			IsNull: true,
		}

		sp = strings.Split(strings.TrimSpace(sp[2]), " ")

		f.TypeName = sp[0]
		switch f.TypeName {
		case "datetime", "time", "data", "timestamp":
			tinfo.Pkg["time"] = 1
		}

		sp = sp[1:]
		l := len(sp)
		for i := 0; i < l; i++ {
			name := sp[i]
			switch name {
			case "unsigned":
				f.TypeName += " " + name
			case "not":
				i++
				f.IsNull = false
			case "default":
				f.Default = strings.ReplaceAll(sp[i+1], "'", "")
				i++
			case "comment":
				f.Comment = strings.ReplaceAll(sp[i+1], "'", "")
				i++
			}
		}

		tinfo.Fields = append(tinfo.Fields, f)

	}

	buildGoByTabInfo(tinfo)
}

func buildGoByTabInfo(tinfo tableInfo) {
	dirPath, _ := os.Getwd()
	if len(os.Args) == 2 {
		p := os.Args[1]
		if strings.HasPrefix(p, "/") {
			dirPath = p
		} else {
			dirPath += "/" + p
		}
	}
	dirname, _ := filepath.Abs(dirPath)
	fname := dirname + "/" + tinfo.Name + ".go"
	pname := strings.ReplaceAll(path.Base(dirname), "-", "")

	mname := buildName(tinfo.Name) + "Model"

	data := []byte("")
	buf := bytes.NewBuffer(data)

	buf.WriteString(fmt.Sprintf("package %s\n\n", pname))
	for k, _ := range tinfo.Pkg {
		buf.WriteString(fmt.Sprintf("import \"%s\"\n\n", k))
	}
	buf.WriteString(fmt.Sprintf("// %s\n", tinfo.Comment))
	buf.WriteString(fmt.Sprintf("type %s struct {\n", mname))

	for _, f := range tinfo.Fields {
		buf.WriteString(fmt.Sprintf("\t%s\t", buildName(f.Name)))
		if strings.HasPrefix(f.TypeName, "time") {
			buf.WriteString("time.Time")
		} else if strings.HasPrefix(f.TypeName, "datetime") {
			buf.WriteString("time.Time")
		} else if strings.HasPrefix(f.TypeName, "date") {
			buf.WriteString("time.Time")
		} else if strings.HasPrefix(f.TypeName, "bigint") {
			if strings.HasSuffix(f.TypeName, "unsigned") {
				buf.WriteString("uint64")
			} else {
				buf.WriteString("int64")
			}
		} else if strings.HasPrefix(f.TypeName, "int") {
			if strings.HasSuffix(f.TypeName, "unsigned") {
				buf.WriteString("uint32")
			} else {
				buf.WriteString("int32")
			}
		} else if strings.HasPrefix(f.TypeName, "float") {
			buf.WriteString("float32")
		} else if strings.HasPrefix(f.TypeName, "double") {
			buf.WriteString("float64")
		} else if strings.HasPrefix(f.TypeName, "tinyint") {
			buf.WriteString("int")
		} else {
			buf.WriteString("string")
		}
		dv := ""
		en := ""
		if !f.IsNull {
			en = "NOT NULL;"
			dv = "default:" + f.Default + ";"
		}
		buf.WriteString(fmt.Sprintf("\t`gorm:\"column:%s;type:%s;%s%scomment:%s", f.Name, f.TypeName, dv, en, f.Comment))
		buf.WriteString(fmt.Sprintf("\" json:\"%s\" form:\"%s\" get:\"%s\" post:\"%s\"`\n", f.Name, f.Name, f.Name, f.Name))
	}

	buf.WriteString(fmt.Sprintf("}\n\nfunc (m *%s) TableName() string {\n\treturn \"%s\"\n}\n", mname, tinfo.Name))

	// fmt.Println("", fname, pname, string(buf.Bytes()), mname)
	ioutil.WriteFile(fname, buf.Bytes(), 0777)
}

func buildName(name string) string {
	ns := strings.Split(name, "_")
	nss := []string{}
	for _, i := range ns {
		nss = append(nss, strings.ToUpper(i[0:1])+i[1:])
	}
	return strings.Join(nss, "")
}
