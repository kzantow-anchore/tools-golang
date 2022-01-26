// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v2

import (
	"fmt"
	"reflect"

	"github.com/spdx/tools-golang/spdx"
)

func (spec JSONSpdxDocument) parseJsonRelationships2_2(key string, value interface{}, doc *spdxDocument2_2) error {

	//FIXME : NOASSERTION and NONE in relationship B value not compatible
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		relationships := reflect.ValueOf(value)
		for i := 0; i < relationships.Len(); i++ {
			ifc := relationships.Index(i).Interface()
			relationship, err := requireMap(ifc)
			if err != nil {
				return fmt.Errorf("invalid value for relationship, expected map[string] but got: %+v", ifc)
			}
			rel := spdx.Relationship2_2{}
			// Parse ref A of the relationship
			s, err := requireMapString(relationship, "spdxElementId")
			if err != nil {
				return err
			}
			aid, err := extractDocElementID(s)
			if err != nil {
				return fmt.Errorf("%s", err)
			}
			rel.RefA = aid

			// Parse the refB of the relationship
			// NONE and NOASSERTION are permitted on right side
			permittedSpecial := []string{"NONE", "NOASSERTION"}
			s, err = requireMapString(relationship, "relatedSpdxElement")
			if err != nil {
				return err
			}
			bid, err := extractDocElementSpecial(s, permittedSpecial)
			if err != nil {
				return fmt.Errorf("%s", err)
			}
			rel.RefB = bid
			// Parse relationship type
			if relationship["relationshipType"] == nil {
				return fmt.Errorf("%s , %d", "RelationshipType propty missing in relationship number", i)
			}
			rel.Relationship, err = requireMapString(relationship, "relationshipType")
			if err != nil {
				return err
			}

			// Parse the relationship comment
			if relationship["comment"] != nil {
				rel.RelationshipComment, err = requireMapString(relationship, "comment")
				if err != nil {
					return err
				}
			}

			doc.Relationships = append(doc.Relationships, &rel)
		}

	}
	return nil
}
