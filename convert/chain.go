package convert

import (
	"github.com/anchore/go-struct-converter"
	"github.com/spdx/tools-golang/spdx/common"

	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/spdx/v2_1"
	"github.com/spdx/tools-golang/spdx/v2_2"
	"github.com/spdx/tools-golang/spdx/v2_3"
)

var DocumentChain = converter.NewChain(
	v2_1.Document{},
	v2_2.Document{},
	v2_3.Document{},
)

// Document converts the provided SPDX document to the latest version
func Document(doc common.Document) (spdx.Document, error) {
	if doc, ok := doc.(spdx.Document); ok {
		return doc, nil
	}
	latest := spdx.Document{}
	err := DocumentChain.Convert(doc, &latest)
	if err != nil {
		return spdx.Document{}, err
	}
	return latest, nil
}
