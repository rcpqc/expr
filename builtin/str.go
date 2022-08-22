package builtin

import (
	"fmt"
	"strconv"
)

func init() {
	Functions["btoi"] = btoi
	Functions["btof"] = btof
	Functions["stoi"] = stoi
	Functions["stof"] = stof
	Functions["itos"] = itos
	Functions["slen"] = slen
	Functions["format"] = format
}

func btoi(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func btof(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func stoi(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func stof(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func itos(i int64) string {
	return strconv.FormatInt(i, 10)
}

func slen(s string) int64 {
	return int64(len(s))
}

func format(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
