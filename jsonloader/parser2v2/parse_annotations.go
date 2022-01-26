// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v2

import (
	"fmt"
	"reflect"

	"github.com/spdx/tools-golang/spdx"
)

func (spec JSONSpdxDocument) parseJsonAnnotations2_2(key string, value interface{}, doc *spdxDocument2_2, SPDXElementId spdx.DocElementID) (err error) {
	//FIXME: SPDXID property not defined in spec but it is needed
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		annotations := reflect.ValueOf(value)
		for i := 0; i < annotations.Len(); i++ {
			annotation, err := requireMap(annotations.Index(i).Interface())
			ann := spdx.Annotation2_2{AnnotationSPDXIdentifier: SPDXElementId}
			// Remove loop all properties are mandatory in annotations
			for k, v := range annotation {
				switch k {
				case "annotationDate":
					ann.AnnotationDate, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for Annotation annotationDate: %+v", v)
					}
				case "annotationType":
					ann.AnnotationType, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for annotationType: %+v", v)
					}
				case "comment":
					ann.AnnotationComment, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for comment: %+v", v)
					}
				case "annotator":
					subkey, subvalue, err := extractSubs(v)
					if err != nil {
						return err
					}
					if subkey != "Person" && subkey != "Organization" && subkey != "Tool" {
						return fmt.Errorf("unrecognized Annotator type %v", subkey)
					}
					ann.AnnotatorType = subkey
					ann.Annotator = subvalue

				default:
					return fmt.Errorf("received unknown property %v in Annotation section", k)
				}
			}
			doc.Annotations = append(doc.Annotations, &ann)
		}

	}
	return nil
}
