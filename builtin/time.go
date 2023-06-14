package builtin

import "time"

func init() {
	Variables["time"] = unix
	Variables["tnow"] = tnow
	Variables["tfmt"] = tfmt
	Variables["tprs"] = tprs
}

func unix(timestamp int64) time.Time {
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
