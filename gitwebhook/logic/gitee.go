package logic

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type GitEEProject struct {
	Token    string
	UserName string
	Repo     string
	Branch   string
	IsGo     bool
	baseDir  string
	RunCmd   string
	RunArgs  []string
}

// 程序运行
func (g GitEEProject) Run() {
	g.baseDir = "/tmp/" + g.UserName + "/" + g.Repo + "/"
	os.RemoveAll(g.baseDir)
	os.MkdirAll(g.baseDir, 0777)
	e := g.loadDir("master", "")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	c := exec.Command("go", "mod", "tidy")
	c.Dir = g.baseDir
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	e1 := c.Run()
	if e1 != nil {
		fmt.Println(e1.Error())
		return
	}

	c = exec.Command("go", "build", "-o", "../../"+g.UserName+"_"+g.Repo)
	c.Dir = g.baseDir
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	e1 = c.Run()
	if e1 != nil {
		fmt.Println(e1.Error())
		return
	}
	outfile, _ := filepath.Abs(g.baseDir + "../../" + g.UserName + "_" + g.Repo)
	exec.Command("killall", "-9", outfile)
	args := []string{outfile}
	args = append(args, g.RunArgs...)
	if g.RunCmd == "" {
		c = exec.Command("setsid", args...)
	} else {
		c = exec.Command(g.RunCmd, args...)
	}

	c.Dir = g.baseDir
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	e1 = c.Run()
	if e1 != nil {
		fmt.Println(e1.Error())
		return
	}
}

func (g GitEEProject) loadDir(sha, dir string) error {
	ret := loadTreeObj{}
	e := HttpGet(fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/git/trees/%s?access_token=%s", g.UserName, g.Repo, sha, g.Token), &ret)
	if e != nil {
		return e
	}
	for _, item := range ret.Tree {
		if item.Type == "tree" {
			os.MkdirAll(g.baseDir+dir+item.Path, 0777)
			e1 := g.loadDir(item.Sha, dir+item.Path+"/")
			if e1 != nil {
				return e1
			}
		} else {
			finfo := fileObj{}
			fmt.Println("load file ->" + dir + item.Path)
			e2 := HttpGet(item.URL+"?access_token="+g.Token, &finfo)
			if e2 != nil {
				return e2
			}
			c, l := base64.StdEncoding.DecodeString(finfo.Content)
			if l != nil {
				fmt.Println("decode err ->" + finfo.Content)
				return l
			}
			ee := ioutil.WriteFile(g.baseDir+dir+item.Path, c, 0777)
			if ee != nil {
				fmt.Println("wfile err ->" + g.baseDir + dir + item.Path)
				return ee
			}

		}
	}

	return nil
}

type loadTreeObj struct {
	Sha       string        `json:"sha"`
	URL       string        `json:"url"`
	Tree      []treeItemObj `json:"tree"`
	Truncated bool          `json:"truncated"`
}

type treeItemObj struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Sha  string `json:"sha"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}
type fileObj struct {
	Sha      string `json:"sha"`
	Size     int    `json:"size"`
	URL      string `json:"url"`
	Content  string `json:"content"`
	Encoding string `json:"encoding"`
}
