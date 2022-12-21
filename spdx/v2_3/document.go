// Package v2_3 Package contains the struct definition for an SPDX Document
// and its constituent parts.
// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package v2_3

import (
	"fmt"
	"strings"

	"github.com/spdx/tools-golang/spdx/common"
	tv "github.com/spdx/tools-golang/tagvalue/lib"
)

const Version = "SPDX-2.3"
const DataLicense = "CC0-1.0"

// ExternalDocumentRef is a reference to an external SPDX document as defined in section 6.6
type ExternalDocumentRef struct {
	// DocumentRefID is the ID string defined in the start of the
	// reference. It should _not_ contain the "DocumentRef-" part
	// of the mandatory ID string.
	DocumentRefID string `json:"externalDocumentId"`

	// URI is the URI defined for the external document
	URI string `json:"spdxDocument"`

	// Checksum is the actual hash data
	Checksum common.Checksum `json:"checksum"`
}

func (d ExternalDocumentRef) ToTagValue() (string, error) {
	return fmt.Sprintf("%s %s %s", common.PrefixDocumentRef(common.ElementID(d.DocumentRefID)), d.URI, d.Checksum), nil
}

func (d *ExternalDocumentRef) FromTagValue(s string) error {
	parts := strings.SplitN(s, " ", 3)
	if len(parts) == 3 {
		elementID, err := common.TrimDocumentRefPrefix(parts[0])
		if err != nil {
			return err
		}
		d.DocumentRefID = string(elementID)
		d.URI = parts[1]
		err = d.Checksum.FromTagValue(parts[2])
		if err != nil {
			return err
		}
	}
	return nil
}

var _ tv.ToValue = (*ExternalDocumentRef)(nil)
var _ tv.FromValue = (*ExternalDocumentRef)(nil)

// Document is an SPDX Document:
// See https://spdx.github.io/spdx-spec/v2.3/document-creation-information
type Document struct {
	// 6.1: SPDX Version; should be in the format "SPDX-<version>"
	// Cardinality: mandatory, one
	SPDXVersion string `json:"spdxVersion"`

	// 6.2: Data License; should be "CC0-1.0"
	// Cardinality: mandatory, one
	DataLicense string `json:"dataLicense"`

	// 6.3: SPDX Identifier; should be "DOCUMENT" to represent
	//      mandatory identifier of SPDXRef-DOCUMENT
	// Cardinality: mandatory, one
	SPDXIdentifier common.ElementID `json:"SPDXID" tv:"SPDXID"`

	// 6.4: Document Name
	// Cardinality: mandatory, one
	DocumentName string `json:"name"`

	// 6.5: Document Namespace
	// Cardinality: mandatory, one
	DocumentNamespace string `json:"documentNamespace"`

	// 6.6: External Document References
	// Cardinality: optional, one or many
	ExternalDocumentReferences []ExternalDocumentRef `json:"externalDocumentRefs,omitempty" tv:"ExternalDocumentRef"`

	// 6.11: Document Comment
	// Cardinality: optional, one
	DocumentComment string `json:"comment,omitempty"`

	CreationInfo  *CreationInfo   `json:"creationInfo"`
	Packages      []*Package      `json:"packages,omitempty"`
	Files         []*File         `json:"files,omitempty"`
	OtherLicenses []*OtherLicense `json:"hasExtractedLicensingInfos,omitempty"`
	Relationships []*Relationship `json:"relationships,omitempty"`
	Annotations   []*Annotation   `json:"annotations,omitempty"`
	Snippets      []Snippet       `json:"snippets,omitempty"`

	// DEPRECATED in version 2.0 of spec
	Reviews []*Review `json:"-" yaml:"-"`
}
