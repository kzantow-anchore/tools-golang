// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package common

import (
	"encoding/json"
	"fmt"
	"strings"

	tv "github.com/spdx/tools-golang/tagvalue/lib"
)

const (
	AnnotatorTypePerson       = "Person"
	AnnotatorTypeOrganization = "Organization"
	AnnotatorTypeTool         = "Tool"

	AnnotationTypeReview = "REVIEW"
	AnnotationTypeOther  = "OTHER"
)

type Annotator struct {
	Annotator string
	// including AnnotatorType: one of "Person", "Organization" or "Tool"
	AnnotatorType string
}

func (d Annotator) ToTagValue() (string, error) {
	return fmt.Sprintf("%s: %s", d.AnnotatorType, d.Annotator), nil
}

func (d *Annotator) FromTagValue(s string) error {
	parts := strings.Split(s, ": ")
	if len(parts) == 2 {
		d.AnnotatorType = parts[0]
		d.Annotator = parts[1]
	}
	return nil
}

var _ tv.ToValue = (*Annotator)(nil)
var _ tv.FromValue = (*Annotator)(nil)

// UnmarshalJSON takes an annotator in the typical one-line format and parses it into an Annotator struct.
// This function is also used when unmarshalling YAML
func (a *Annotator) UnmarshalJSON(data []byte) error {
	// annotator will simply be a string
	annotatorStr := string(data)
	annotatorStr = strings.Trim(annotatorStr, "\"")

	annotatorFields := strings.SplitN(annotatorStr, ": ", 2)

	if len(annotatorFields) != 2 {
		return fmt.Errorf("failed to parse Annotator '%s'", annotatorStr)
	}

	a.AnnotatorType = annotatorFields[0]
	a.Annotator = annotatorFields[1]

	return nil
}

// MarshalJSON converts the receiver into a slice of bytes representing an Annotator in string form.
// This function is also used when marshalling to YAML
func (a Annotator) MarshalJSON() ([]byte, error) {
	if a.Annotator != "" {
		return json.Marshal(fmt.Sprintf("%s: %s", a.AnnotatorType, a.Annotator))
	}

	return []byte{}, nil
}
