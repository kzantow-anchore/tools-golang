package convert

import (
	"fmt"
	"reflect"
)

type Converter func(from interface{}, to interface{}) error

type ConvertFrom interface {
	ConvertFrom(interface{}, Converter) error
}

const convertFromName = "ConvertFrom"

var (
	nilValue        = reflect.ValueOf(nil)
	convertFromType = reflect.TypeOf((*ConvertFrom)(nil))
)

// Convert takes two objects, e.g. v2_1.Document and &v2_2.Document{} and attempts to
// map all the properties from one to the other
func Convert(from interface{}, to interface{}) error {
	fromValue := reflect.ValueOf(from)
	fromType := fromValue.Type()

	// handle incoming pointers
	for fromType.Kind() == reflect.Ptr {
		fromValue = fromValue.Elem()
		fromType = fromValue.Type()
	}

	toValue := reflect.ValueOf(to)
	toType := toValue.Type()

	if toType.Kind() != reflect.Ptr {
		return fmt.Errorf("to struct provided was not a pointer, unable to set values: %v", to)
	}

	// handle incoming pointers
	for toType.Kind() == reflect.Ptr {
		toValue = toValue.Elem()
		toType = toValue.Type()
	}

	// Assume structs at this point
	for i := 0; i < fromType.NumField(); i++ {
		fromField := fromType.Field(i)
		fromFieldValue := fromValue.Field(i)

		toField, exists := toType.FieldByName(fromField.Name)
		if !exists {
			continue
		}

		toFieldValue := toValue.FieldByIndex(toField.Index)

		newValue, err := getValue(fromFieldValue, toField.Type)

		if err != nil {
			return err
		}

		msg := fmt.Sprintf("from value: %v to value: %v", fromFieldValue.Interface(), newValue.Interface())
		fmt.Println(msg)

		toFieldValue.Set(newValue)
	}

	return nil
}

func getValue(fromValue reflect.Value, targetType reflect.Type) (reflect.Value, error) {
	fromType := fromValue.Type()

	var toValue reflect.Value

	// handle incoming pointers
	if fromType.Kind() == reflect.Ptr {
		fromValue = fromValue.Elem()
		fromType = fromValue.Type()
	}

	toType := targetType
	if toType.Kind() == reflect.Ptr {
		toType = toType.Elem()
	}

	switch {
	case fromType.Kind() == reflect.Struct:
		// this always creates a pointer type
		toValue = reflect.New(toType)

		for i := 0; i < fromType.NumField(); i++ {
			fromField := fromType.Field(i)
			fromFieldValue := fromValue.Field(i)

			toField, exists := toType.FieldByName(fromField.Name)
			if !exists {
				continue
			}
			toFieldType := toField.Type

			toFieldValue := toValue.FieldByIndex(toField.Index)

			newValue, err := getValue(fromFieldValue, toFieldType)
			if err != nil {
				return nilValue, err
			}

			toFieldValue.Set(newValue)
		}

		// allow structs to implement a custom convert function from previous/next version struct
		if toType.Implements(convertFromType) {
			convertFrom := toValue.MethodByName(convertFromName)
			if !convertFrom.IsValid() {
				return nilValue, fmt.Errorf("unable to get ConvertFrom method")
			}
			args := []reflect.Value{}
			out := convertFrom.Call(args)
			err := out[0].Interface()
			if err != nil {
				return nilValue, fmt.Errorf("an error occurred calling %s.%s: %v", toType.Name(), convertFromName, err)
			}
		}
	case fromType.Kind() == reflect.Slice:
		length := fromValue.Len()
		if toType.Kind() == reflect.Slice {
			elementType := toType.Elem()
			toValue = reflect.MakeSlice(toType, length, length)
			for i := 0; i < length; i++ {
				v, err := getValue(fromValue.Index(i), elementType)
				if err != nil {
					return nilValue, err
				}
				toValue.Index(i).Set(v)
			}
		} else {
			// TODO try to convert a slice to something else
		}
	default:
		// TODO determine if there is another conversion for the rest of the types
		toValue = fromValue
	}

	// handle non-pointer returns -- the reflect.New earlier always creates a pointer
	if toType.Kind() != reflect.Ptr {
		toValue = fromPointer(toValue)
	}

	// handle elements which are now pointers
	if targetType.Kind() == reflect.Ptr {
		toValue = toPointer(toValue)
	}

	return toValue, nil
}

func toPointer(val reflect.Value) reflect.Value {
	typ := val.Type()
	if typ.Kind() != reflect.Ptr {
		// this creates a pointer type inherently
		ptrVal := reflect.New(typ)
		ptrVal.Elem().Set(val)
		val = ptrVal
	}
	return val
}

func fromPointer(val reflect.Value) reflect.Value {
	if val.Type().Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
