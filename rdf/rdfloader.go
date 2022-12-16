// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdx_rdf

import (
	"errors"
	"fmt"
	"io"

	"github.com/spdx/gordf/rdfloader"
	gordfParser "github.com/spdx/gordf/rdfloader/parser"
	"github.com/spdx/tools-golang/common"
	"github.com/spdx/tools-golang/convert"
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/spdx/v2_2"
	v2_2_reader "github.com/spdx/tools-golang/spdx/v2_2/rdf/reader"
	"github.com/spdx/tools-golang/spdx/v2_3"
	v2_3_reader "github.com/spdx/tools-golang/spdx/v2_3/rdf/reader"
)

func Read(content io.Reader) (*spdx.Document, error) {
	var rdfParserObj, err = rdfloader.LoadFromReaderObject(content)
	if err != nil {
		return nil, err
	}

	version, err := getSpdxVersion(rdfParserObj)
	if err != nil {
		return nil, err
	}

	var data interface{}
	switch version {
	case v2_2.Version:
		data, err = v2_2_reader.LoadFromGoRDFParser(rdfParserObj)
	case v2_3.Version:
		data, err = v2_3_reader.LoadFromGoRDFParser(rdfParserObj)
	default:
		return nil, fmt.Errorf("unsupported SPDX version: '%v'", version)
	}

	if err != nil {
		return nil, err
	}

	out, err := convert.Document(data.(common.Document))
	return &out, err
}

func getSpdxVersion(parser *gordfParser.Parser) (string, error) {
	version := ""
	for _, node := range parser.Triples {
		if node.Predicate.ID == "http://spdx.org/rdf/terms#specVersion" {
			version = node.Object.ID
			break
		}
	}
	if version == "" {
		return "", errors.New("unable to determine version from RDF document")
	}
	return version, nil
}
