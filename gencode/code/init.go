package code

type CmdFunc func(string) string

type Cmd struct {
	Help   string
	Handle CmdFunc
}
