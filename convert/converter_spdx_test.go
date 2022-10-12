package convert

import (
	"testing"

	"github.com/spdx/tools-golang/spdx/v2_2"
	"github.com/spdx/tools-golang/spdx/v2_3"
)

func Test_ConvertSPDXDocuments(t *testing.T) {
	v2_2doc := v2_2.Document{
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
			},
		},
	}

	v2_3doc := v2_3.Document{}

	err := Convert(v2_2doc, &v2_3doc)
	if err != nil {
		t.Fatalf("unable to convert: %v", err)
	}

	if len(v2_3doc.Packages) != 1 {
		t.Errorf("Incorrect Packages length: %v", len(v2_3doc.Packages))
	}

	pkg := v2_3doc.Packages[0]
	files := pkg.Files
	if len(files) != 2 {
		t.Errorf("Incorrect Files length: %v", len(files))
	}

	if files[0].FileName != "File 1" {
		t.Errorf("Incorrect File name: %v", files[0].FileName)
	}

	if files[1].FileName != "File 2" {
		t.Errorf("Incorrect File name: %v", files[1].FileName)
	}
}
