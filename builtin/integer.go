package builtin

import "strconv"

func init() {
	variables["int"] = func(v int) int { return v }
	variables["int8"] = func(v int8) int8 { return v }
	variables["int16"] = func(v int16) int16 { return v }
	variables["int32"] = func(v int32) int32 { return v }
	variables["int64"] = func(v int64) int64 { return v }
	variables["uint"] = func(v uint) uint { return v }
	variables["uint8"] = func(v uint8) uint8 { return v }
	variables["uint16"] = func(v uint16) uint16 { return v }
	variables["uint32"] = func(v uint32) uint32 { return v }
	variables["uint64"] = func(v uint64) uint64 { return v }
	variables["itos"] = func(v int64) string { return strconv.FormatInt(v, 10) }
	variables["utos"] = func(v uint64) string { return strconv.FormatUint(v, 10) }
}
