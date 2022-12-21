// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package v2_2

import (
	"fmt"

	"github.com/spdx/tools-golang/spdx/common"
	tv "github.com/spdx/tools-golang/tagvalue/lib"
)

// Relationship is a Relationship section of an SPDX Document for
// version 2.2 of the spec.
type Relationship struct {

	// 11.1: Relationship
	// Cardinality: optional, one or more; one per Relationship
	//              one mandatory for SPDX Document with multiple packages
	// RefA and RefB are first and second item
	// Relationship is type from 11.1.1
	RefA         common.DocElementID `json:"spdxElementId"`
	RefB         common.DocElementID `json:"relatedSpdxElement"`
	Relationship string              `json:"relationshipType"`

	// 11.2: Relationship Comment
	// Cardinality: optional, one
	RelationshipComment string `json:"comment,omitempty"`
}

type relationshipTagValue struct {
	Relationship        string
	RelationshipComment string
}

func (d Relationship) GetTagValue() (interface{}, error) {
	a, err := d.RefA.ToTagValue()
	if err != nil {
		return nil, err
	}
	b, err := d.RefB.ToTagValue()
	if err != nil {
		return nil, err
	}
	return relationshipTagValue{
		Relationship:        fmt.Sprintf("%s %s %s", a, d.Relationship, b),
		RelationshipComment: d.RelationshipComment,
	}, nil
}

func (d Relationship) FromTagValue(i interface{}) error {
	//TODO implement me
	panic("implement me")
}

var _ tv.TagValueHandler = (*Relationship)(nil)
