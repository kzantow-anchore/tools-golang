package convert

import (
	"reflect"
	"testing"

	"github.com/spdx/tools-golang/spdx/common"
	"github.com/spdx/tools-golang/spdx/v2_2"
	"github.com/spdx/tools-golang/spdx/v2_3"
)

func Test_ConvertSPDXDocuments(t *testing.T) {
	tests := []struct {
		name     string
		source   interface{}
		expected interface{}
	}{
		{
			name: "basic v2_2 to v2_3",
			source: v2_2.Document{
				Packages: []*v2_2.Package{
					{
						PackageName: "Pkg 1",
						Files: []*v2_2.File{
							{
								FileName: "File 1",
							},
							{
								FileName: "File 2",
							},
						},
						PackageVerificationCode: common.PackageVerificationCode{
							Value: "verification code value",
							ExcludedFiles: []string{
								"a",
								"b",
							},
						},
					},
				},
			},
			expected: v2_3.Document{
				Packages: []*v2_3.Package{
					{
						PackageName: "Pkg 1",
						Files: []*v2_3.File{
							{
								FileName: "File 1",
							},
							{
								FileName: "File 2",
							},
						},
						PackageVerificationCode: &common.PackageVerificationCode{
							Value: "verification code value",
							ExcludedFiles: []string{
								"a",
								"b",
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outType := reflect.TypeOf(test.expected)
			outInstance := reflect.New(outType).Interface()
			err := Convert(test.source, outInstance)
			if err != nil {
				t.Fatalf("error converting: %v", err)
			}
			outInstance = reflect.ValueOf(outInstance).Elem().Interface()

			if !reflect.DeepEqual(test.expected, outInstance) {
				t.Fatalf("structs do not match: %+v %+v", test.expected, outInstance)
			}
		})
	}
}
