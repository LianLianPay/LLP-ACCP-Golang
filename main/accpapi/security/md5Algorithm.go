package security

import (
	"crypto/md5"
)

func Md5Digest(str string) []byte {
	hash := md5.New()
	hash.Write([]byte(str))
	return hash.Sum(nil)
}
