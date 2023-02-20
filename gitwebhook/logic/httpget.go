package logic

import (
	"encoding/json"
	"net/http"
)

func HttpGet(url string, obj any) error {
	r, e := http.Get(url)
	if e != nil {
		return e
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(obj)
}
