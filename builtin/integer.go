package builtin

import "strconv"

func init() {
	Functions["int"] = func(v int) int { return v }
	Functions["int8"] = func(v int8) int8 { return v }
	Functions["int16"] = func(v int16) int16 { return v }
	Functions["int32"] = func(v int32) int32 { return v }
	Functions["int64"] = func(v int64) int64 { return v }
	Functions["uint"] = func(v uint) uint { return v }
	Functions["uint8"] = func(v uint8) uint8 { return v }
	Functions["uint16"] = func(v uint16) uint16 { return v }
	Functions["uint32"] = func(v uint32) uint32 { return v }
	Functions["uint64"] = func(v uint64) uint64 { return v }
	Functions["itos"] = func(v int64) string { return strconv.FormatInt(v, 10) }
	Functions["utos"] = func(v uint64) string { return strconv.FormatUint(v, 10) }
}
