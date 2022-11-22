// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdx_json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/spdx/tools-golang/convert"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/v2_1"
	"github.com/spdx/tools-golang/v2_2"
)

func Read(content io.Reader) (*spdx.Document, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(content)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(buf.Bytes(), &data)
	if err != nil {
		return nil, err
	}

	val, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("not a valid SPDX JSON document")
	}

	version, ok := val["spdxVersion"]
	if !ok {
		return nil, fmt.Errorf("JSON document does not contain spdxVersion field")
	}

	switch version {
	case "SPDX-2.1":
		var doc v2_1.Document
		err = json.Unmarshal(buf.Bytes(), &doc)
		if err != nil {
			return nil, err
		}
		data = doc
	case "SPDX-2.2":
		var doc v2_2.Document
		err = json.Unmarshal(buf.Bytes(), &doc)
		if err != nil {
			return nil, err
		}
		data = doc
	case spdx.Version:
		var doc spdx.Document
		err = json.Unmarshal(buf.Bytes(), &doc)
		if err != nil {
			return nil, err
		}
		data = doc
	default:
		return nil, fmt.Errorf("unsupported SDPX version: %s", version)
	}

	out, err := convert.ConvertDocument(data)
	return &out, err
}

// Load2_2 takes in an io.Reader and returns an SPDX document.
func Load2_2(content io.Reader) (*v2_2.Document, error) {
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

// Load2_3 takes in an io.Reader and returns an SPDX document.
func Load2_3(content io.Reader) (*spdx.Document, error) {
	// convert io.Reader to a slice of bytes and call the parser
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(content)
	if err != nil {
		return nil, err
	}

	var doc spdx.Document
	err = json.Unmarshal(buf.Bytes(), &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}
