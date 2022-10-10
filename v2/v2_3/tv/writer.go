// Package tvsaver is used to save tools-golang data structures
// as SPDX tag-value documents.
// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package tv

import (
	"io"

	"github.com/spdx/tools-golang/v2/v2_3"
	"github.com/spdx/tools-golang/v2/v2_3/tv/writer"
)

// Save takes an io.Writer and an SPDX Document (version 2.3),
// and writes it to the writer in tag-value format. It returns error
// if any error is encountered.
func Write(doc *v2_3.Document, w io.Writer) error {
	return writer.RenderDocument(doc, w)
}
