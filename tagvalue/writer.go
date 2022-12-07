package tagvalue

import (
	"io"

	"github.com/spdx/tools-golang/spdx/common"
	tv "github.com/spdx/tools-golang/tagvalue/lib"
)

func Write(doc common.Document, writer io.Writer) error {
	return tv.Write(doc, writer)
}
