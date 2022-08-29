package builtin

import "time"

func init() {
	Functions["time"] = _time
	Functions["tfmt"] = tfmt
	Functions["tenc"] = tenc
	Functions["tdec"] = tdec
}

func _time(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

func tfmt(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

func tenc(t time.Time, layout string) string {
	return t.Format(layout)
}

func tdec(s string, layout string) time.Time {
	t, _ := time.Parse(layout, s)
	return t
}
