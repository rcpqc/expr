package builtin

import (
	"fmt"
	"strconv"
)

func init() {
	Functions["stoi"] = stoi
	Functions["stou"] = stou
	Functions["stof"] = stof
	Functions["str"] = str
	Functions["slen"] = slen
	Functions["sfmt"] = sfmt
}

func stoi(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func stou(s string) uint64 {
	u, _ := strconv.ParseUint(s, 10, 64)
	return u
}

func stof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func str(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func slen(s string) int64 {
	return int64(len(s))
}

func sfmt(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
