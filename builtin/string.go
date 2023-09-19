package builtin

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	variables["stoi"] = stoi
	variables["stou"] = stou
	variables["stof"] = stof
	variables["str"] = str
	variables["slen"] = slen
	variables["sfmt"] = fmt.Sprintf
	variables["split"] = strings.Split
	variables["sjoin"] = strings.Join
	variables["sfind"] = strings.Index
	variables["slower"] = strings.ToLower
	variables["supper"] = strings.ToUpper
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

func str(v any) string {
	return fmt.Sprintf("%v", v)
}

func slen(s string) int64 {
	return int64(len(s))
}
