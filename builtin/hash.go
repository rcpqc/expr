package builtin

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
)

func init() {
	Functions["md5"] = _md5
	Functions["sha1"] = _sha1
	Functions["sha256"] = _sha256
}
func _md5(data []byte) []byte {
	sum := md5.Sum(data)
	return sum[:]
}
func _sha1(data []byte) []byte {
	sum := sha1.Sum(data)
	return sum[:]
}
func _sha256(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}
