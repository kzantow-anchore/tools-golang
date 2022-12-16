package convert

import (
	converter "github.com/anchore/go-struct-converter"

	"github.com/spdx/tools-golang/common"
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

// Document converts the provided SPDX document to the latest verison
func Document(doc common.Document) (spdx.Document, error) {
	latest := spdx.Document{}
	err := DocumentChain.Convert(doc, &latest)
	if err != nil {
		return spdx.Document{}, err
	}
	return latest, nil
}
