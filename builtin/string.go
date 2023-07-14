package builtin

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	Variables["stoi"] = stoi
	Variables["stou"] = stou
	Variables["stof"] = stof
	Variables["str"] = str
	Variables["slen"] = slen
	Variables["sfmt"] = fmt.Sprintf
	Variables["split"] = strings.Split
	Variables["sjoin"] = strings.Join
	Variables["sfind"] = strings.Index
	Variables["slower"] = strings.ToLower
	Variables["supper"] = strings.ToUpper
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
