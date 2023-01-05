package code

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"io/ioutil"
)

func GzipEncodeBase64(src string) string {
	var b bytes.Buffer
	gw, _ := gzip.NewWriterLevel(&b, zlib.BestSpeed)
	// gw := gzip.NewWriterLevel(&b)
	gw.Write([]byte(src))
	gw.Flush()
	gw.Close()
	// fmt.Println(len(b.Bytes()), len(string(b.Bytes())))
	return base64.RawStdEncoding.EncodeToString(b.Bytes())
}
func GzipDecodeBase64(src string) string {
	dd, _ := base64.RawStdEncoding.DecodeString(src)
	b := bytes.NewReader(dd)
	gw, _ := gzip.NewReader(b)
	date, _ := ioutil.ReadAll(gw)
	return string(date)
}
