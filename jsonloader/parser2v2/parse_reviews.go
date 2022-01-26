// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v2

import (
	"fmt"
	"reflect"

	"github.com/spdx/tools-golang/spdx"
)

func (spec JSONSpdxDocument) parseJsonReviews2_2(key string, value interface{}, doc *spdxDocument2_2) (err error) {
	//FIXME: Reviewer type property of review not specified in the spec
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		reviews := reflect.ValueOf(value)
		for i := 0; i < reviews.Len(); i++ {
			ifc := reviews.Index(i).Interface()
			reviewMap, err := requireMap(ifc)
			if err != nil {
				return fmt.Errorf("invalid data type for reviews; expected map[string], got: %+v", ifc)
			}
			review := spdx.Review2_2{}
			// Remove loop all properties are mandatory in annotations
			for k, v := range reviewMap {
				switch k {
				case "reviewer":
					subkey, subvalue, err := extractSubs(v)
					if err != nil {
						return err
					}
					if subkey != "Person" && subkey != "Organization" && subkey != "Tool" {
						return fmt.Errorf("unrecognized Reviewer type %v", subkey)
					}
					review.ReviewerType = subkey
					review.Reviewer = subvalue
				case "comment":
					review.ReviewComment, err = requireString(v)
					if err != nil {
						return fmt.Errorf("review comment required but got error: %w", err)
					}
				case "reviewDate":
					review.ReviewDate, err = requireString(v)
					if err != nil {
						return fmt.Errorf("review reviewDate required but got error: %w", err)
					}
				default:
					return fmt.Errorf("received unknown property %v in Review Section section", k)
				}
			}
			doc.Reviews = append(doc.Reviews, &review)
		}

	}
	return nil
}
