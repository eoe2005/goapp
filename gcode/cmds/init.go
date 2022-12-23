package cmds

type CmdInterFace interface {
	Help() string
	Run()
}
