package logic

import "os"

func GetGiteeToken() string {
	return os.Getenv("gitee_token")
}
func SetGiteeToken(token string) {
	os.Setenv("gitee_token", token)
}
