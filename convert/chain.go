package convert

import (
	converter "github.com/anchore/go-struct-converter"

	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/v2_1"
	"github.com/spdx/tools-golang/v2_2"
)

var DocumentChain = converter.NewChain(
	v2_1.Document{},
	v2_2.Document{},
	spdx.Document{},
)

// ConvertDocument converts the provided SPDX document to the latest verison
func ConvertDocument(doc interface{}) (spdx.Document, error) {
	latest := spdx.Document{}
	err := DocumentChain.Convert(doc, &latest)
	if err != nil {
		return spdx.Document{}, err
	}
	return latest, nil
}

var PackageChain = converter.NewChain(
	v2_1.Package{},
	v2_2.Package{},
	spdx.Package{},
)

// ConvertPackage converts the provided SPDX document to the latest verison
func ConvertPackage(pkg interface{}) (spdx.Package, error) {
	latest := spdx.Package{}
	err := PackageChain.Convert(pkg, &latest)
	if err != nil {
		return spdx.Package{}, err
	}
	return latest, nil
}

var FileChain = converter.NewChain(
	v2_1.File{},
	v2_2.File{},
	spdx.File{},
)

// ConvertFile converts the provided SPDX document to the latest verison
func ConvertFile(file interface{}) (spdx.File, error) {
	latest := spdx.File{}
	err := FileChain.Convert(file, &latest)
	if err != nil {
		return spdx.File{}, err
	}
	return latest, nil
}
