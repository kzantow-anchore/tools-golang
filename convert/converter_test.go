package convert

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Convert(t *testing.T) {
	type s1 struct {
		Value string
	}

	type s2 struct {
		Value string
	}

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
			name: "string slice to ptrs slice",
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
		{
			name: "string slice ptr to slice",
			fromStruct: struct {
				PtrToStrings *[]string
			}{
				PtrToStrings: &[]string{"the thing 1", "the thing 2"},
			},
			toStruct: struct {
				PtrToStrings []string
			}{
				PtrToStrings: []string{"the thing 1", "the thing 2"},
			},
		},
		{
			name: "string slice ptr to slice",
			fromStruct: struct {
				PtrToStrings []string
			}{
				PtrToStrings: []string{"the thing 1", "the thing 2"},
			},
			toStruct: struct {
				PtrToStrings *[]string
			}{
				PtrToStrings: &[]string{"the thing 1", "the thing 2"},
			},
		},
		{
			name: "struct to struct",
			fromStruct: struct {
				Struct s1
			}{
				Struct: s1{
					Value: "something",
				},
			},
			toStruct: struct {
				Struct s2
			}{
				Struct: s2{
					Value: "something",
				},
			},
		},
		{
			name: "struct to ptr",
			fromStruct: struct {
				Struct s1
			}{
				Struct: s1{
					Value: "something",
				},
			},
			toStruct: struct {
				Struct *s2
			}{
				Struct: &s2{
					Value: "something",
				},
			},
		},
		{
			name: "struct ptr to struct",
			fromStruct: struct {
				Struct *s1
			}{
				Struct: &s1{
					Value: "something",
				},
			},
			toStruct: struct {
				Struct s2
			}{
				Struct: s2{
					Value: "something",
				},
			},
		},
		{
			name: "struct ptr to struct ptr",
			fromStruct: struct {
				Struct *s1
			}{
				Struct: &s1{
					Value: "something",
				},
			},
			toStruct: struct {
				Struct *s2
			}{
				Struct: &s2{
					Value: "something",
				},
			},
		},
		{
			name: "map to map",
			fromStruct: struct {
				Map map[s1]string
			}{
				Map: map[s1]string{
					s1{Value: "some-key"}: "some-value",
				},
			},
			toStruct: struct {
				Map map[s2]string
			}{
				Map: map[s2]string{
					s2{Value: "some-key"}: "some-value",
				},
			},
		},
		{
			name: "map to map of ptr",
			fromStruct: struct {
				Map map[s1]string
			}{
				Map: map[s1]string{
					s1{Value: "some-key"}: "some-value",
				},
			},
			toStruct: struct {
				Map map[s2]*string
			}{
				Map: map[s2]*string{
					s2{Value: "some-key"}: s("some-value"),
				},
			},
		},
		{
			name: "map key ptr to map",
			fromStruct: struct {
				Map map[*s1]string
			}{
				Map: map[*s1]string{
					&s1{Value: "some-key"}: "some-value",
				},
			},
			toStruct: struct {
				Map map[s2]string
			}{
				Map: map[s2]string{
					s2{Value: "some-key"}: "some-value",
				},
			},
		},
		{
			name: "map ptr to map",
			fromStruct: struct {
				Map *map[*s1]string
			}{
				Map: &map[*s1]string{
					&s1{Value: "some-key"}: "some-value",
				},
			},
			toStruct: struct {
				Map map[s2]string
			}{
				Map: map[s2]string{
					s2{Value: "some-key"}: "some-value",
				},
			},
		},
		{
			name: "map string to int",
			fromStruct: struct {
				Value string
			}{
				Value: "12",
			},
			toStruct: struct {
				Value int
			}{
				Value: 12,
			},
		},
		{
			name: "map int to string",
			fromStruct: struct {
				Value int
			}{
				Value: 12,
			},
			toStruct: struct {
				Value string
			}{
				Value: "12",
			},
		},
		{
			name: "map string slice to string",
			fromStruct: struct {
				Value []string
			}{
				Value: []string{"thing 1"},
			},
			toStruct: struct {
				Value string
			}{
				Value: "thing 1",
			},
		},
		{
			name: "map string to string slice",
			fromStruct: struct {
				Value string
			}{
				Value: "thing 1",
			},
			toStruct: struct {
				Value []string
			}{
				Value: []string{"thing 1"},
			},
		},
		{
			name: "map int slice to string",
			fromStruct: struct {
				Value []int
			}{
				Value: []int{84},
			},
			toStruct: struct {
				Value string
			}{
				Value: "84",
			},
		},
		{
			name: "map string to uint slice",
			fromStruct: struct {
				Value string
			}{
				Value: "63",
			},
			toStruct: struct {
				Value []uint
			}{
				Value: []uint{63},
			},
		},
		{
			name: "map string to uint16 slice",
			fromStruct: struct {
				Value string
			}{
				Value: "63",
			},
			toStruct: struct {
				Value []uint16
			}{
				Value: []uint16{63},
			},
		},
		{
			name: "test top-level conversion functions",
			fromStruct: T3{
				T1: T1{
					Same:     "same value",
					OldValue: "old value",
				},
			},
			toStruct: T4{
				T2: T2{
					Same:     "same value",
					NewValue: "old value",
				},
			},
		},
		{
			name: "test nested conversion functions",
			fromStruct: T5{
				Version: "2.2",
				Embedded: T3{
					T1: T1{
						Same:     "same value",
						OldValue: "old value",
					},
				},
			},
			toStruct: T6{
				Version: "2.3",
				Embedded: T4{
					T2: T2{
						Same:     "same value",
						NewValue: "old value",
					},
				},
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

// -------- structs to test conversion functions:

type T1 struct {
	Same     string
	OldValue string
}

type T2 struct {
	Same     string
	NewValue string
}

func (t *T2) ConvertFrom(i interface{}) error {
	t1 := i.(T1)
	t.NewValue = t1.OldValue
	return nil
}

var _ From = (*T2)(nil)

type T3 struct {
	T1 T1
}

type T4 struct {
	T2 T2
}

func (t *T4) ConvertFrom(i interface{}) error {
	t3 := i.(T3)
	return Convert(t3.T1, &t.T2)
}

var _ From = (*T4)(nil)

type T5 struct {
	Version  string
	Embedded T3
}

type T6 struct {
	Version  string
	Embedded T4
}

func (t *T6) ConvertFrom(_ interface{}) error {
	t.Version = "2.3"
	return nil
}

var _ From = (*T6)(nil)
