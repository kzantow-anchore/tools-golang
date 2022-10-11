package convert

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Convert(t *testing.T) {
	tests := []struct {
		name       string
		fromStruct interface{}
		toStruct   interface{}
	}{
		{
			name: "missing properties are omitted",
			fromStruct: struct {
				Value string
			}{
				Value: "the value",
			},
			toStruct: struct {
				Other string
			}{},
		},
		{
			name: "string equals",
			fromStruct: struct {
				Value string
			}{
				Value: "the value",
			},
			toStruct: struct {
				Value string
			}{
				Value: "the value",
			},
		},
		{
			name: "string ptr equals",
			fromStruct: struct {
				String *string
			}{
				String: s("the value"),
			},
			toStruct: struct {
				String *string
			}{
				String: s("the value"),
			},
		},
		{
			name: "int equals",
			fromStruct: struct {
				Int int
			}{
				Int: 2,
			},
			toStruct: struct {
				Int int
			}{
				Int: 2,
			},
		},
		{
			name: "bool equals",
			fromStruct: struct {
				Int bool
			}{
				Int: true,
			},
			toStruct: struct {
				Int bool
			}{
				Int: true,
			},
		},
		{
			name: "string ptr equals",
			fromStruct: struct {
				StringPtr *string
			}{
				StringPtr: s("the value"),
			},
			toStruct: struct {
				StringPtr *string
			}{
				StringPtr: s("the value"),
			},
		},
		{
			name: "string to ptr equals",
			fromStruct: struct {
				StringPtr string
			}{
				StringPtr: "the value",
			},
			toStruct: struct {
				StringPtr *string
			}{
				StringPtr: s("the value"),
			},
		},
		{
			name: "string from ptr equals",
			fromStruct: struct {
				StringPtr *string
			}{
				StringPtr: s("the value"),
			},
			toStruct: struct {
				StringPtr string
			}{
				StringPtr: "the value",
			},
		},
		{
			name: "string slice",
			fromStruct: struct {
				Strings []string
			}{
				Strings: []string{"the name"},
			},
			toStruct: struct {
				Strings []string
			}{
				Strings: []string{"the name"},
			},
		},
		{
			name: "string ptr slice",
			fromStruct: struct {
				StringsPtr []*string
			}{
				StringsPtr: []*string{s("thing 1"), s("thing 2")},
			},
			toStruct: struct {
				StringsPtr []*string
			}{
				StringsPtr: []*string{s("thing 1"), s("thing 2")},
			},
		},
		{
			name: "string ptr to string slice",
			fromStruct: struct {
				StringsPtrToStr []*string
			}{
				StringsPtrToStr: []*string{s("thing 1"), s("thing 2")},
			},
			toStruct: struct {
				StringsPtrToStr []string
			}{
				StringsPtrToStr: []string{"thing 1", "thing 2"},
			},
		},
		{
			name: "string to ptrs slice",
			fromStruct: struct {
				StringsToPtrStr []string
			}{
				StringsToPtrStr: []string{"thing 1", "thing 2"},
			},
			toStruct: struct {
				StringsToPtrStr []*string
			}{
				StringsToPtrStr: []*string{s("thing 1"), s("thing 2")},
			},
		},
		{
			name: "string slice ptr",
			fromStruct: struct {
				PtrToStrings *[]string
			}{
				PtrToStrings: &[]string{"the thing 1", "the thing 2"},
			},
			toStruct: struct {
				PtrToStrings *[]string
			}{
				PtrToStrings: &[]string{"the thing 1", "the thing 2"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			typ := reflect.TypeOf(test.toStruct)
			newInstance := reflect.New(typ)
			result := newInstance.Interface()

			err := Convert(test.fromStruct, result)
			if err != nil {
				t.Fatalf("error during conversion: %v", err)
			}

			str := fmt.Sprintf("got: %+v", result)
			fmt.Println(str)

			// need to align elem vs. pointer of the result
			result = reflect.ValueOf(result).Elem().Interface()

			to := test.toStruct
			if !reflect.DeepEqual(to, result) {
				t.Fatalf("Convert output does not match: %+v %+v", to, result)
			}
		})
	}
}

func s(s string) *string {
	return &s
}
