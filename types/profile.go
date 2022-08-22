package types

import (
	"reflect"
	"strings"
)

// Profile type's Profile
type Profile struct {
	tagIndex map[string]int
}

// NewProfile construct type's profile
func NewProfile(t reflect.Type, tagkey string) *Profile {
	val, _ := LoadOrCreate(t, func(t reflect.Type) interface{} {
		return (&Profile{}).init(t, tagkey)
	})
	return val.(*Profile)
}

// init initialize profile
func (o *Profile) init(t reflect.Type, tagkey string) *Profile {
	o.tagIndex = map[string]int{}
	for i := 0; i < t.NumField(); i++ {
		if !t.Field(i).IsExported() {
			continue
		}
		tag := t.Field(i).Tag.Get(tagkey)
		if tag == "-" {
			continue
		}
		if tag == "" {
			tag = strings.ToLower(t.Field(i).Name)
		}
		o.tagIndex[tag] = i
	}
	return o
}

// FieldFromTagName get struct's field by tagname
func (o *Profile) FieldFromTagName(rv reflect.Value, tag string) reflect.Value {
	idx, ok := o.tagIndex[tag]
	if !ok || idx < 0 || idx >= rv.NumField() {
		return reflect.Value{}
	}
	return rv.Field(idx)
}
