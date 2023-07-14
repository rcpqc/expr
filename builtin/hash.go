package builtin

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func init() {
	Variables["md5"] = md5_
	Variables["sha1"] = sha1_
	Variables["sha256"] = sha256_
	Variables["hex"] = hex.EncodeToString
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
