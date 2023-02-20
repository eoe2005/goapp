package logic

type GiteeParam struct {
	LangType   string   `json:"lang_type"`
	RunCmd     string   `json:"run_cmd"`
	RunArgs    []string `json:"run_args"`
	Branch     string   `json:"branch"`
	Repository struct {
		Name      string `json:"name"`
		Path      string `json:"path"`
		NameSpace string `json:"namespace"`
	} `json:"repository"`
}

func (p *GiteeParam) Run() {
	switch p.LangType {
	case "go":
		p.runGo()
	}

}

func (p *GiteeParam) runGo() {
	if p.Branch == "" {
		p.Branch = "master"
	}
	ge := GitEEProject{
		Token:    GetGiteeToken(),
		UserName: p.Repository.NameSpace,
		Repo:     p.Repository.Name,
		Branch:   p.Branch,
		IsGo:     true,
		RunCmd:   p.RunCmd,
		RunArgs:  p.RunArgs,
	}
	ge.Run()
}
