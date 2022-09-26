package builtin

import "time"

func init() {
	Functions["time"] = _time
	Functions["tnow"] = tnow
	Functions["tunix"] = tunix
	Functions["tfmt"] = tfmt
	Functions["tprs"] = tprs
	Functions["tenc"] = tenc
	Functions["tdec"] = tdec
}

func _time(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
func tnow() time.Time {
	return time.Now()
}
func tunix(t time.Time) int64 {
	return t.Unix()
}
func tfmt(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}
func tprs(s string, layout string) int64 {
	t, _ := time.Parse(layout, s)
	return t.Unix()
}
func tenc(t time.Time, layout string) string {
	return t.Format(layout)
}
func tdec(s string, layout string) time.Time {
	t, _ := time.Parse(layout, s)
	return t
}
