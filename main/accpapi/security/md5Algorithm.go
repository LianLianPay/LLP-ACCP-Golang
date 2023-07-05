package security

import (
	"crypto/md5"
)

func Md5Digest(str []byte) []byte {
	hash := md5.New()

	// 将字符串转换为字节数组并写入哈希对象
	hash.Write(str)

	// 计算哈希值
	hashValue := hash.Sum(nil)
	return hashValue
}
