package builtin

import "strconv"

func init() {
	Variables["int"] = func(v int) int { return v }
	Variables["int8"] = func(v int8) int8 { return v }
	Variables["int16"] = func(v int16) int16 { return v }
	Variables["int32"] = func(v int32) int32 { return v }
	Variables["int64"] = func(v int64) int64 { return v }
	Variables["uint"] = func(v uint) uint { return v }
	Variables["uint8"] = func(v uint8) uint8 { return v }
	Variables["uint16"] = func(v uint16) uint16 { return v }
	Variables["uint32"] = func(v uint32) uint32 { return v }
	Variables["uint64"] = func(v uint64) uint64 { return v }
	Variables["itos"] = func(v int64) string { return strconv.FormatInt(v, 10) }
	Variables["utos"] = func(v uint64) string { return strconv.FormatUint(v, 10) }
}
