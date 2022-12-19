package convert

import (
	"fmt"
	"reflect"

	"github.com/spdx/tools-golang/spdx/common"
)

// FromPtr accepts a document or a document pointer and returns the direct struct reference
func FromPtr(doc common.Document) common.Document {
	value := reflect.ValueOf(doc)
	for value.Type().Kind() == reflect.Ptr {
		value = value.Elem()
	}
	return value.Interface()
}

func Describe(o interface{}) string {
	value := reflect.ValueOf(o)
	typ := value.Type()
	prefix := ""
	for typ.Kind() == reflect.Ptr {
		prefix += "*"
		value = value.Elem()
		typ = value.Type()
	}
	str := limit(fmt.Sprintf("%+v", value.Interface()), 300)
	name := fmt.Sprintf("%s.%s%s", typ.PkgPath(), prefix, typ.Name())
	return fmt.Sprintf("%s: %s", name, str)
}

func limit(text string, length int) string {
	if length <= 0 || len(text) <= length+3 {
		return text
	}
	r := []rune(text)
	r = r[:length]
	return string(r) + "..."
}
