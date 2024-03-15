package gfs

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

func DecodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func EncodeBase64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// Md5Reader
//f, _ := os.Open("./_example/1.jpeg")
//f.Seek(0, 0)
//gfs.Md5Reader(f)
func Md5Reader(fr io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, fr); err != nil {
		return "", fmt.Errorf(`io.Copy failed`)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
