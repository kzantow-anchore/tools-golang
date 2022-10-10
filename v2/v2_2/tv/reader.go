// Package tvloader is used to load and parse SPDX tag-value documents
// into tools-golang data structures.
// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package tv

import (
	"io"

	"github.com/spdx/tools-golang/common/tv"
	"github.com/spdx/tools-golang/v2/v2_2"
	"github.com/spdx/tools-golang/v2/v2_2/tv/reader"
)

// Read takes an io.Reader and returns a fully-parsed SPDX Document
// (version 2.2) if parseable, or error if any error is encountered.
func Read(content io.Reader) (*v2_2.Document, error) {
	tvPairs, err := tv.ReadTagValues(content)
	if err != nil {
		return nil, err
	}

	doc, err := reader.ParseTagValues(tvPairs)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
