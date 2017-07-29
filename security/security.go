package security

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

const (
	salt = "bobi_funny"
)

//加盐md5
func Md5Salt(s string) string {
	return Md5(Md5(s + salt))
}

//生成32位md5字串
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func Guid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return Md5(base64.URLEncoding.EncodeToString(b))
}
