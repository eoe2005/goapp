package code

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"io/ioutil"

	"github.com/klauspost/compress/flate"
)

func GzipEncodeBase64(src string) string {
	var b bytes.Buffer
	gw, _ := gzip.NewWriterLevel(&b, 9)
	// gw := gzip.NewWriter(&b)
	gw.Write([]byte(src))
	// gw.Flush()
	gw.Close()
	// fmt.Println(len(src), len(b.Bytes()), b.Bytes())
	return base64.StdEncoding.EncodeToString(b.Bytes())
}
func GzipDecodeBase64(src string) string {
	dd, _ := base64.RawStdEncoding.DecodeString(src)
	b := bytes.NewReader(dd)
	gw, _ := gzip.NewReader(b)
	date, _ := ioutil.ReadAll(gw)
	return string(date)
}
func ZibEncodeBase64(src string) string {
	var b bytes.Buffer
	gw, _ := zlib.NewWriterLevel(&b, 1)
	// gw := gzip.NewWriter(&b)
	gw.Write([]byte(src))
	gw.Flush()
	gw.Close()
	// fmt.Println(len(src), len(b.Bytes()), b.Bytes())
	return base64.StdEncoding.EncodeToString(b.Bytes())
}
func ZlibDecodeBase64(src string) string {
	dd, _ := base64.RawStdEncoding.DecodeString(src)
	b := bytes.NewReader(dd)
	gw, _ := zlib.NewReader(b)
	date, _ := ioutil.ReadAll(gw)
	return string(date)
}
func FlateEncodeBase64(src string) string {
	var b bytes.Buffer
	gw, _ := flate.NewWriter(&b, 1)
	// gw := gzip.NewWriter(&b)
	gw.Write([]byte(src))
	gw.Flush()
	gw.Close()
	// fmt.Println(len(src), len(b.Bytes()), b.Bytes())
	return base64.StdEncoding.EncodeToString(b.Bytes())
}
func FlateDecodeBase64(src string) string {
	// dd, _ := base64.RawStdEncoding.DecodeString(src)
	// b := bytes.NewReader(dd)
	// gw, _ := flate.NewReader(&b)
	// date, _ := ioutil.ReadAll(gw)
	// return string(date)
	return ""
}
