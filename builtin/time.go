package builtin

import "time"

func init() {
	Functions["tfmt"] = tfmt
}

func tfmt(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}
