package types

import (
	"reflect"
	"regexp"
	"strings"
)

// Profile type's Profile
type Profile struct {
	indices map[string]int
	methods map[string]struct{}
}

// NewProfile construct type's profile
func NewProfile(t reflect.Type, tagkey string) *Profile {
	val, _ := LoadOrCreate(t, func(t reflect.Type) any {
		return (&Profile{}).init(t, tagkey)
	})
	return val.(*Profile)
}

var firstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var allCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// snake translate to snake case
func snake(s string) string {
	snake := firstCap.ReplaceAllString(s, "${1}_${2}")
	snake = allCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// init initialize profile
func (o *Profile) init(t reflect.Type, tagkey string) *Profile {
	o.indices = map[string]int{}
	o.methods = map[string]struct{}{}
	// for ptr, methods + element's fields
	// for struct, methods + fields
	for i := 0; i < t.NumMethod(); i++ {
		tag := snake(t.Method(i).Name)
		o.indices[tag] = i
		o.methods[tag] = struct{}{}
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return o
	}
	for i := 0; i < t.NumField(); i++ {
		if !t.Field(i).IsExported() {
			continue
		}
		tag := t.Field(i).Tag.Get(tagkey)
		if tag == "-" {
			continue
		}
		if tag == "" {
			tag = snake(t.Field(i).Name)
		}
		o.indices[tag] = i
	}
	return o
}

// Select get struct's field/method by tagname
func (o *Profile) Select(rv reflect.Value, tag string) (reflect.Value, bool) {
	_, method := o.methods[tag]
	idx, found := o.indices[tag]
	// not found
	if !found {
		return reflect.Value{}, false
	}
	// method or field
	if method {
		return rv.Method(idx), true
	} else {
		rv = reflect.Indirect(rv)
		if !rv.IsValid() {
			return reflect.Value{}, true
		}
		return rv.Field(idx), true
	}
}
