package builtin

import "time"

func init() {
	Functions["time"] = _time
	Functions["tnow"] = tnow
	Functions["tfmt"] = tfmt
	Functions["tprs"] = tprs
}

func _time(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
func tnow() int64 {
	return time.Now().Unix()
}
func tfmt(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}
func tprs(s string, layout string) int64 {
	t, _ := time.Parse(layout, s)
	return t.Unix()
}
