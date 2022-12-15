// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdx

import (
	"fmt"

	tv "github.com/spdx/tools-golang/tagvalue/lib"
)

// Review is a Review section of an SPDX Document for version 2.3 of the spec.
// DEPRECATED in version 2.0 of spec; retained here for compatibility.
type Review struct {

	// DEPRECATED in version 2.0 of spec
	// 13.1: Reviewer
	// Cardinality: optional, one
	Reviewer string
	// including AnnotatorType: one of "Person", "Organization" or "Tool"
	ReviewerType string

	// DEPRECATED in version 2.0 of spec
	// 13.2: Review Date: YYYY-MM-DDThh:mm:ssZ
	// Cardinality: conditional (mandatory, one) if there is a Reviewer
	ReviewDate string

	// DEPRECATED in version 2.0 of spec
	// 13.3: Review Comment
	// Cardinality: optional, one
	ReviewComment string
}

type reviewTagValue struct {
	Reviewer      string
	ReviewDate    string
	ReviewComment string
}

func (d Review) GetTagValue() (interface{}, error) {
	return reviewTagValue{
		Reviewer:      fmt.Sprintf("%s: %s", d.ReviewerType, d.Reviewer),
		ReviewDate:    d.ReviewDate,
		ReviewComment: d.ReviewComment,
	}, nil
}

func (d Review) FromTagValue(i interface{}) error {
	//TODO implement me
	panic("implement me")
}

var _ tv.TagValueHandler = (*Review)(nil)
