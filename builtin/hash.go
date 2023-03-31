package builtin

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func init() {
	Functions["md5"] = md5_
	Functions["sha1"] = sha1_
	Functions["sha256"] = sha256_
	Functions["hex"] = func(data []byte) string { return hex.EncodeToString(data) }
}
func md5_(data []byte) []byte {
	sum := md5.Sum(data)
	return sum[:]
}
func sha1_(data []byte) []byte {
	sum := sha1.Sum(data)
	return sum[:]
}
func sha256_(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}
