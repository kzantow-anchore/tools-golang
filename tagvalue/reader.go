// Package tvloader is used to load and parse SPDX tag-value documents
// into tools-golang data structures.
// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package spdx_tagvalue

import (
	"fmt"
	"io"

	"github.com/spdx/tools-golang/common"
	"github.com/spdx/tools-golang/convert"
	"github.com/spdx/tools-golang/spdx"
	spdx_reader "github.com/spdx/tools-golang/spdx/tagvalue/reader"
	"github.com/spdx/tools-golang/tagvalue/reader"
	"github.com/spdx/tools-golang/v2_1"
	v2_1_reader "github.com/spdx/tools-golang/v2_1/tagvalue/reader"
	"github.com/spdx/tools-golang/v2_2"
	v2_2_reader "github.com/spdx/tools-golang/v2_2/tagvalue/reader"
)

// Read takes an io.Reader and returns a fully-parsed SPDX Document
// if parseable, or error if any error is encountered.
func Read(content io.Reader) (*spdx.Document, error) {
	tvPairs, err := reader.ReadTagValues(content)
	if err != nil {
		return nil, err
	}

	if len(tvPairs) == 0 {
		return nil, fmt.Errorf("no tag values found")
	}

	version := ""
	for _, pair := range tvPairs {
		if pair.Tag == "SPDXVersion" {
			version = pair.Value
			break
		}
	}

	var data interface{}
	switch version {
	case v2_1.Version:
		data, err = v2_1_reader.ParseTagValues(tvPairs)
	case v2_2.Version:
		data, err = v2_2_reader.ParseTagValues(tvPairs)
	case spdx.Version:
		data, err = spdx_reader.ParseTagValues(tvPairs)
	default:
		return nil, fmt.Errorf("unsupported SPDX version: '%v'", version)
	}

	if err != nil {
		return nil, err
	}

	out, err := convert.Document(data.(common.Document))
	return &out, err
}
