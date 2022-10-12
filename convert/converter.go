package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

type Converter func(from interface{}, to interface{}) error

type ConvertFrom interface {
	ConvertFrom(interface{}, Converter) error
}

const convertFromName = "ConvertFrom"

var (
	nilValue        = reflect.ValueOf(nil)
	convertFromType = reflect.TypeOf((*ConvertFrom)(nil)).Elem()
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

		if !fromFieldValue.IsValid() || fromFieldValue.IsZero() || fromFieldValue.IsNil() {
			continue
		}

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
	var err error

	fromType := fromValue.Type()

	var toValue reflect.Value

	// handle incoming pointer types
	if fromType.Kind() == reflect.Ptr {
		fromValue = fromValue.Elem()
		if !fromValue.IsValid() {
			return nilValue, nil
		}
		if fromValue.IsZero() || fromValue.IsNil() {
			return nilValue, nil
		}
		fromType = fromValue.Type()
	}

	baseTargetType := targetType
	if targetType.Kind() == reflect.Ptr {
		baseTargetType = targetType.Elem()
	}

	switch {
	case fromType.Kind() == reflect.Struct && baseTargetType.Kind() == reflect.Struct:
		// this always creates a pointer type
		toValue = reflect.New(baseTargetType)
		toValue = toValue.Elem()

		for i := 0; i < fromType.NumField(); i++ {
			fromField := fromType.Field(i)
			fromFieldValue := fromValue.Field(i)

			toField, exists := baseTargetType.FieldByName(fromField.Name)
			if !exists {
				continue
			}
			toFieldType := toField.Type

			toFieldValue := toValue.FieldByName(toField.Name)

			newValue, err := getValue(fromFieldValue, toFieldType)
			if err != nil {
				return nilValue, err
			}

			toFieldValue.Set(newValue)
		}

		// allow structs to implement a custom convert function from previous/next version struct
		if baseTargetType.Implements(convertFromType) {
			convertFrom := toValue.MethodByName(convertFromName)
			if !convertFrom.IsValid() {
				return nilValue, fmt.Errorf("unable to get ConvertFrom method")
			}
			args := []reflect.Value{}
			out := convertFrom.Call(args)
			err := out[0].Interface()
			if err != nil {
				return nilValue, fmt.Errorf("an error occurred calling %s.%s: %v", baseTargetType.Name(), convertFromName, err)
			}
		}
	case fromType.Kind() == reflect.Slice && baseTargetType.Kind() == reflect.Slice:
		length := fromValue.Len()
		targetElementType := baseTargetType.Elem()
		toValue = reflect.MakeSlice(baseTargetType, length, length)
		for i := 0; i < length; i++ {
			v, err := getValue(fromValue.Index(i), targetElementType)
			if err != nil {
				return nilValue, err
			}
			if v.IsValid() {
				toValue.Index(i).Set(v)
			}
		}
	case fromType.Kind() == reflect.Map && baseTargetType.Kind() == reflect.Map:
		keyType := baseTargetType.Key()
		elementType := baseTargetType.Elem()
		toValue = reflect.MakeMap(baseTargetType)
		for _, fromKey := range fromValue.MapKeys() {
			fromVal := fromValue.MapIndex(fromKey)
			k, err := getValue(fromKey, keyType)
			if err != nil {
				return nilValue, err
			}
			v, err := getValue(fromVal, elementType)
			if err != nil {
				return nilValue, err
			}
			if k.IsValid() && v.IsValid() {
				toValue.SetMapIndex(k, v)
			}
		}
	default:
		// TODO determine if there is another conversion for the rest of the types
		toValue = fromValue
	}

	// handle non-pointer returns -- the reflect.New earlier always creates a pointer
	// FIXME: i don't think this is doing anything, baseTargetType is never a ptr
	if baseTargetType.Kind() != reflect.Ptr {
		toValue = fromPtr(toValue)
	}

	toValue, err = convertValueTypes(toValue, baseTargetType)

	if err != nil {
		return nilValue, err
	}

	// handle elements which are now pointers
	if targetType.Kind() == reflect.Ptr {
		toValue = toPtr(toValue)
	}

	return toValue, nil
}

// convertValueTypes takes a value and a target type, and attempts to convert
// between the types - e.g. string -> int. when this function is called the value
func convertValueTypes(value reflect.Value, targetType reflect.Type) (reflect.Value, error) {
	typ := value.Type()
	switch {
	// if the types are the same, just return the value
	case typ.Kind() == targetType.Kind():
		return value, nil
	case value.IsZero() && isPrimitive(targetType):

	case isPrimitive(typ) && isPrimitive(targetType):
		// get a string representation of the value
		str := fmt.Sprintf("%v", value.Interface()) // FIXME is there a better way?
		var err error
		var out interface{}
		switch {
		case isString(targetType):
			out = str
		case isBool(targetType):
			out, err = strconv.ParseBool(str)
		case isInt(targetType):
			out, err = strconv.Atoi(str)
		case isUint(targetType):
			out, err = strconv.ParseUint(str, 10, 64)
		case isFloat(targetType):
			out, err = strconv.ParseFloat(str, 64)
		}

		if err != nil {
			return nilValue, err
		}

		v := reflect.ValueOf(out)

		v = v.Convert(targetType)

		return v, nil
	case isSlice(typ) && isSlice(targetType):
		// TODO -- this should already be handled
	case isSlice(typ):
		// TODO this may be lossy
		if value.Len() > 0 {
			v := value.Index(0)
			v, err := convertValueTypes(v, targetType)
			if err != nil {
				return nilValue, err
			}
			return v, nil
		}
		return convertValueTypes(nilValue, targetType)
	case isSlice(targetType):
		elementType := targetType.Elem()
		v, err := convertValueTypes(value, elementType)
		if err != nil {
			return nilValue, err
		}
		slice := reflect.MakeSlice(targetType, 1, 1)
		slice.Index(0).Set(v)
		return slice, nil
	}

	return nilValue, fmt.Errorf("unable to convert from: %v to %v", value.Interface(), targetType.Name())
}

func isPrimitive(typ reflect.Type) bool {
	return isString(typ) || isBool(typ) || isInt(typ) || isUint(typ) || isFloat(typ)
}

func isString(typ reflect.Type) bool {
	return typ.Kind() == reflect.String
}

func isBool(typ reflect.Type) bool {
	return typ.Kind() == reflect.Bool
}

func isInt(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return true
	}
	return false
}

func isUint(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return true
	}
	return false
}

func isFloat(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Float32,
		reflect.Float64:
		return true
	}
	return false
}

func isSlice(typ reflect.Type) bool {
	return typ.Kind() == reflect.Slice
}

func toPtr(val reflect.Value) reflect.Value {
	typ := val.Type()
	if typ.Kind() != reflect.Ptr {
		// this creates a pointer type inherently
		ptrVal := reflect.New(typ)
		ptrVal.Elem().Set(val)
		val = ptrVal
	}
	return val
}

func fromPtr(val reflect.Value) reflect.Value {
	typ := val.Type()
	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
