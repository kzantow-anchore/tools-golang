// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdx_json

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/spdx/tools-golang/v2/v2_2"
)

// Read takes in an io.Reader and returns an SPDX document.
func Read(content io.Reader) (*v2_2.Document, error) {
	// convert io.Reader to a slice of bytes and call the parser
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(content)
	if err != nil {
		return nil, err
	}

	var doc v2_2.Document
	err = json.Unmarshal(buf.Bytes(), &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}
