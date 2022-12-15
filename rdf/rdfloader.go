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
	"github.com/spdx/tools-golang/spdx/rdf/reader"
	spdx_reader "github.com/spdx/tools-golang/spdx/rdf/reader"
	"github.com/spdx/tools-golang/v2_2"
	v2_2_reader "github.com/spdx/tools-golang/v2_2/rdf/reader"
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
	case spdx.Version:
		data, err = spdx_reader.LoadFromGoRDFParser(rdfParserObj)
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

// Takes in a file Reader and returns the pertaining spdx document
// or the error if any is encountered while setting the doc.
func Load2_2(content io.Reader) (*v2_2.Document, error) {
	var rdfParserObj, err = rdfloader.LoadFromReaderObject(content)
	if err != nil {
		return nil, err
	}

	doc, err := v2_2_reader.LoadFromGoRDFParser(rdfParserObj)
	return doc, err
}

// Takes in a file Reader and returns the pertaining spdx document
// or the error if any is encountered while setting the doc.
func Load2_3(content io.Reader) (*spdx.Document, error) {
	var rdfParserObj, err = rdfloader.LoadFromReaderObject(content)
	if err != nil {
		return nil, err
	}

	doc, err := reader.LoadFromGoRDFParser(rdfParserObj)
	return doc, err
}
