package gfs

import "encoding/base64"

func DecodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func EncodeBase64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
