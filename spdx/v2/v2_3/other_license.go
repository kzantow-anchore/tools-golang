// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package v2_3

import (
	"fmt"

	"github.com/spdx/tools-golang/spdx/v2/common"
)

// OtherLicense is an Other License Information section of an SPDX Document
type OtherLicense struct {
	// 10.1: License Identifier: "LicenseRef-[idstring]"
	// Cardinality: conditional (mandatory, one) if license is not
	//              on SPDX License List
	LicenseIdentifier common.LicenseID `json:"licenseId" tv:"LicenseID"`

	// 10.2: Extracted Text
	// Cardinality: conditional (mandatory, one) if there is a
	//              License Identifier assigned
	ExtractedText string `json:"extractedText"`

	// 10.3: License Name: single line of text or "NOASSERTION"
	// Cardinality: conditional (mandatory, one) if license is not
	//              on SPDX License List
	LicenseName string `json:"name,omitempty"`

	// 10.4: License Cross Reference
	// Cardinality: conditional (optional, one or many) if license
	//              is not on SPDX License List
	LicenseCrossReferences []string `json:"seeAlsos,omitempty" tv:"LicenseCrossReference"`

	// 10.5: License Comment
	// Cardinality: optional, one
	LicenseComment string `json:"comment,omitempty"`
}

func (o OtherLicense) OtherLicense() string {
	return fmt.Sprintf("\n# ----- Other License: %s -----\n", o.LicenseName)
}
