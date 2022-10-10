// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdxlib

import (
	"github.com/spdx/tools-golang/common/spdx"
	v2_12 "github.com/spdx/tools-golang/v2/v2_1"
	v2_22 "github.com/spdx/tools-golang/v2/v2_2"
)

// FilterRelationships2_1 returns a slice of Element IDs returned by the given filter closure. The closure is passed
// one relationship at a time, and it can return an ElementID or nil.
func FilterRelationships2_1(doc *v2_12.Document, filter func(*v2_12.Relationship) *spdx.ElementID) ([]spdx.ElementID, error) {
	elementIDs := []spdx.ElementID{}

	for _, relationship := range doc.Relationships {
		if id := filter(relationship); id != nil {
			elementIDs = append(elementIDs, *id)
		}
	}

	return elementIDs, nil
}

// FilterRelationships2_2 returns a slice of Element IDs returned by the given filter closure. The closure is passed
// one relationship at a time, and it can return an ElementID or nil.
func FilterRelationships2_2(doc *v2_22.Document, filter func(*v2_22.Relationship) *spdx.ElementID) ([]spdx.ElementID, error) {
	elementIDs := []spdx.ElementID{}

	for _, relationship := range doc.Relationships {
		if id := filter(relationship); id != nil {
			elementIDs = append(elementIDs, *id)
		}
	}

	return elementIDs, nil
}
