package builtin

import "time"

func init() {
	Functions["timefmt"] = timefmt
}

func timefmt(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}
