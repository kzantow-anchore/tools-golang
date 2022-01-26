// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v2

import (
	"fmt"
	"reflect"

	"github.com/spdx/tools-golang/spdx"
)

func (spec JSONSpdxDocument) parseJsonOtherLicenses2_2(key string, value interface{}, doc *spdxDocument2_2) (err error) {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		otherlicenses := reflect.ValueOf(value)
		for i := 0; i < otherlicenses.Len(); i++ {
			ifc := otherlicenses.Index(i).Interface()
			licensemap, err := requireMap(ifc)
			if err != nil {
				return fmt.Errorf("invalid value for otherLicenses, expected map[string] but got: %+v", ifc)
			}
			license := spdx.OtherLicense2_2{}
			// Remove loop all properties are mandatory in annotations
			for k, v := range licensemap {
				switch k {
				case "licenseId":
					license.LicenseIdentifier, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for licenseId, expected string but got: %+v", v)
					}
				case "extractedText":
					license.ExtractedText, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for extractedText, expected string but got: %+v", v)
					}
				case "name":
					license.LicenseName, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for name, expected string but got: %+v", v)
					}
				case "comment":
					license.LicenseComment, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for comment, expected string but got: %+v", v)
					}
				case "seeAlsos":
					if reflect.TypeOf(v).Kind() == reflect.Slice {
						texts := reflect.ValueOf(v)
						for i := 0; i < texts.Len(); i++ {
							ifc = texts.Index(i).Interface()
							s, err := requireString(ifc)
							if err != nil {
								return fmt.Errorf("invalid value for seeAlsos, expected string but got: %+v", ifc)
							}
							license.LicenseCrossReferences = append(license.LicenseCrossReferences, s)
						}
					}
				default:
					return fmt.Errorf("received unknown tag %v in Annotation section", k)
				}
			}
			doc.OtherLicenses = append(doc.OtherLicenses, &license)
		}

	}
	return err
}
