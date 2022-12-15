package convert

import (
	converter "github.com/anchore/go-struct-converter"

	"github.com/spdx/tools-golang/common"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/v2_1"
	"github.com/spdx/tools-golang/v2_2"
)

var DocumentChain = converter.NewChain(
	v2_1.Document{},
	v2_2.Document{},
	spdx.Document{},
)

// Document converts the provided SPDX document to the latest verison
func Document(doc common.Document) (spdx.Document, error) {
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

// Package converts the provided SPDX document to the latest verison
func Package(pkg common.Package) (spdx.Package, error) {
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

// File converts the provided SPDX document to the latest verison
func File(file common.File) (spdx.File, error) {
	latest := spdx.File{}
	err := FileChain.Convert(file, &latest)
	if err != nil {
		return spdx.File{}, err
	}
	return latest, nil
}
