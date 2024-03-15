package gfs

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func Md5String(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1String(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256String(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func Base64Encode(src []byte) string {
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

// Base64Reader
//f, _ := os.Open("./_example/1.jpeg")
//f.Seek(0, 0)
//gfs.Base64Reader(f)
func Base64Reader(fi io.Reader) (string, error) {
	fd, err := io.ReadAll(fi)
	if err != nil {
		return "", err
	}
	return Base64Encode(fd), nil
}
