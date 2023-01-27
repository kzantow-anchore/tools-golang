package tv

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

func Write(obj interface{}, writer io.Writer) error {
	return write(obj, "", writer)
}

func write(obj interface{}, tag string, writer io.Writer) error {
	if obj, ok := obj.(TagValuePrefix); ok {
		_, err := fmt.Fprintf(writer, "%s\n", obj.TagValuePrefix())
		if err != nil {
			return err
		}
	}

	if obj, ok := obj.(TagValueHandler); ok {
		obj, err := obj.GetTagValue()
		if err != nil {
			return err
		}
		return write(obj, tag, writer)
	}

	if obj, ok := obj.(ToValue); ok {
		str, err := obj.ToTagValue()
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(writer, "%s: %s\n", tag, str)
		if err != nil {
			return err
		}
		return nil
	}

	val := reflect.ValueOf(obj)
	typ := val.Type()

	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}

	switch typ.Kind() {
	case reflect.Struct:
		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			tag := f.Name
			tags := strings.Split(f.Tag.Get("tv"), ",")
			if tags[0] != "" {
				tag = tags[0]
			}

			if tag == "-" {
				continue
			}

			fVal := val.Field(i)

			if fVal.IsZero() {
				continue
			}

			field := fVal.Interface()

			err := write(field, tag, writer)
			if err != nil {
				return err
			}
		}

	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i).Interface()
			err := write(v, tag, writer)
			if err != nil {
				return err
			}
		}

	case reflect.Map:
		// FIXME

	default:
		v := fmt.Sprintf("%v", obj)
		if strings.Contains(v, "\n") {
			v = fmt.Sprintf("<text>%s</text>", v)
		}
		_, err := fmt.Fprintf(writer, "%s: %s\n", tag, v)
		if err != nil {
			return err
		}
	}

	return nil
}
