// Package parser2v2 contains functions to read, load and parse
// SPDX tag-value files, version 2.2.
// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package reader

import (
	"fmt"

	"github.com/spdx/tools-golang/common/spdx"
	"github.com/spdx/tools-golang/common/tv"
	"github.com/spdx/tools-golang/v2/v2_2"
)

// ParseTagValues takes a list of (tag, value) pairs, parses it and returns
// a pointer to a parsed SPDX Document.
func ParseTagValues(tvs []tv.TagValuePair) (*v2_2.Document, error) {
	parser := tvParser2_2{}
	for _, tv := range tvs {
		err := parser.parsePair(tv.Tag, tv.Value)
		if err != nil {
			return nil, err
		}
	}
	if parser.file != nil && parser.file.FileSPDXIdentifier == nullSpdxElementId2_2 {
		return nil, fmt.Errorf("file with FileName %s does not have SPDX identifier", parser.file.FileName)
	}
	if parser.pkg != nil && parser.pkg.PackageSPDXIdentifier == nullSpdxElementId2_2 {
		return nil, fmt.Errorf("package with PackageName %s does not have SPDX identifier", parser.pkg.PackageName)
	}
	return parser.doc, nil
}

func (parser *tvParser2_2) parsePair(tag string, value string) error {
	switch parser.st {
	case psStart2_2:
		return parser.parsePairFromStart(tag, value)
	case psCreationInfo2_2:
		return parser.parsePairFromCreationInfo(tag, value)
	case psPackage2_2:
		return parser.parsePairFromPackage(tag, value)
	case psFile2_2:
		return parser.parsePairFromFile(tag, value)
	case psSnippet2_2:
		return parser.parsePairFromSnippet(tag, value)
	case psOtherLicense2_2:
		return parser.parsePairFromOtherLicense(tag, value)
	case psReview2_2:
		return parser.parsePairFromReview(tag, value)
	default:
		return fmt.Errorf("parser state %v not recognized when parsing (%s, %s)", parser.st, tag, value)
	}
}

func (parser *tvParser2_2) parsePairFromStart(tag string, value string) error {
	// fail if not in Start parser state
	if parser.st != psStart2_2 {
		return fmt.Errorf("got invalid state %v in parsePairFromStart2_2", parser.st)
	}

	// create an SPDX Document data struct if we don't have one already
	if parser.doc == nil {
		parser.doc = &v2_2.Document{ExternalDocumentReferences: []v2_2.ExternalDocumentRef{}}
	}

	switch tag {
	case "DocumentComment":
		parser.doc.DocumentComment = value
	case "SPDXVersion":
		parser.doc.SPDXVersion = value
	case "DataLicense":
		parser.doc.DataLicense = value
	case "SPDXID":
		eID, err := extractElementID(value)
		if err != nil {
			return err
		}
		parser.doc.SPDXIdentifier = eID
	case "DocumentName":
		parser.doc.DocumentName = value
	case "DocumentNamespace":
		parser.doc.DocumentNamespace = value
	case "ExternalDocumentRef":
		documentRefID, uri, alg, checksum, err := extractExternalDocumentReference(value)
		if err != nil {
			return err
		}
		edr := v2_2.ExternalDocumentRef{
			DocumentRefID: documentRefID,
			URI:           uri,
			Checksum:      spdx.Checksum{Algorithm: spdx.ChecksumAlgorithm(alg), Value: checksum},
		}
		parser.doc.ExternalDocumentReferences = append(parser.doc.ExternalDocumentReferences, edr)
	default:
		// move to Creation Info parser state
		parser.st = psCreationInfo2_2
		return parser.parsePairFromCreationInfo(tag, value)
	}

	return nil
}
