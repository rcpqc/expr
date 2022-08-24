package builtin

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var firstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var allCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func init() {
	Functions["stoi"] = stoi
	Functions["stof"] = stof
	Functions["str"] = str
	Functions["slen"] = slen
	Functions["sfmt"] = sfmt
	Functions["snake"] = Snake
}

func stoi(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
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

// Snake translate to snake case
func Snake(s string) string {
	snake := firstCap.ReplaceAllString(s, "${1}_${2}")
	snake = allCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
