// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v2

import (
	"fmt"
	"reflect"

	"github.com/spdx/tools-golang/spdx"
)

func (spec JSONSpdxDocument) parseJsonSnippets2_2(key string, value interface{}, doc *spdxDocument2_2) (err error) {

	if reflect.TypeOf(value).Kind() == reflect.Slice {
		snippets := reflect.ValueOf(value)
		for i := 0; i < snippets.Len(); i++ {
			ifc := snippets.Index(i).Interface()
			snippetMap, err := requireMap(ifc)
			if err != nil {
				return fmt.Errorf("invalid value for relationship, expected map[string] but got: %+v", ifc)
			}
			// create a new package
			snippet := &spdx.Snippet2_2{}
			//extract the SPDXID of the package
			spdxId, err := requireMapString(snippetMap, "SPDXID")
			if err != nil {
				return fmt.Errorf("invalid value for SPDXID: %w", err)
			}
			eID, err := extractElementID(spdxId)
			if err != nil {
				return fmt.Errorf("%s", err)
			}
			snippet.SnippetSPDXIdentifier = eID
			//range over all other properties now
			for k, v := range snippetMap {
				switch k {
				case "SPDXID", "snippetFromFile":
					//redundant case
				case "name":
					snippet.SnippetName, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for snippet.name, expected string but got: %+v", v)
					}
				case "copyrightText":
					snippet.SnippetCopyrightText, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for snippet.copyrightText, expected string but got: %+v", v)
					}
				case "licenseComments":
					snippet.SnippetLicenseComments, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for snippet.licenseComments, expected string but got: %+v", v)
					}
				case "licenseConcluded":
					snippet.SnippetLicenseConcluded, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for snippet.licenseConcluded, expected string but got: %+v", v)
					}
				case "licenseInfoInSnippets":
					if reflect.TypeOf(v).Kind() == reflect.Slice {
						info := reflect.ValueOf(v)
						for i := 0; i < info.Len(); i++ {
							s, err := requireString(info.Index(i).Interface())
							if err != nil {
								return err
							}
							snippet.LicenseInfoInSnippet = append(snippet.LicenseInfoInSnippet, s)
						}
					}
				case "attributionTexts":
					if reflect.TypeOf(v).Kind() == reflect.Slice {
						info := reflect.ValueOf(v)
						for i := 0; i < info.Len(); i++ {
							s, err := requireString(info.Index(i).Interface())
							if err != nil {
								return err
							}
							snippet.SnippetAttributionTexts = append(snippet.SnippetAttributionTexts, s)
						}
					}
				case "comment":
					snippet.SnippetComment, err = requireString(v)
					if err != nil {
						return fmt.Errorf("invalid value for snippet.licenseConcluded, expected string but got: %+v", v)
					}
				case "ranges":
					//TODO: optimise this logic
					if reflect.TypeOf(v).Kind() == reflect.Slice {
						info := reflect.ValueOf(v)
						for i := 0; i < info.Len(); i++ {
							ranges, err := requireMap(info.Index(i).Interface())
							if err != nil {
								return err
							}
							rangeStart, err := requireMapMap(ranges, "startPointer")
							if err != nil {
								return err
							}
							rangeEnd, err := requireMapMap(ranges, "endPointer")
							if err != nil {
								return err
							}
							if rangeStart["lineNumber"] != nil && rangeEnd["lineNumber"] != nil {
								snippet.SnippetLineRangeStart, err = requireMapFloatInt(rangeStart, "lineNumber")
								if err != nil {
									return err
								}
								snippet.SnippetLineRangeEnd, err = requireMapFloatInt(rangeEnd, "lineNumber")
								if err != nil {
									return err
								}
							} else {
								snippet.SnippetLineRangeStart, err = requireMapFloatInt(rangeStart, "offset")
								if err != nil {
									return err
								}
								snippet.SnippetLineRangeEnd, err = requireMapFloatInt(rangeEnd, "offset")
								if err != nil {
									return err
								}
							}
						}
					}
				default:
					return fmt.Errorf("received unknown tag %v in files section", k)
				}
			}
			snippetFromFile, err := requireMapString(snippetMap, "snippetFromFile")
			if err != nil {
				return err
			}
			fileID, err := extractDocElementID(snippetFromFile)
			if err != nil {
				return err
			}
			snippet.SnippetFromFileSPDXIdentifier = fileID
			if doc.UnpackagedFiles[fileID.ElementRefID].Snippets == nil {
				doc.UnpackagedFiles[fileID.ElementRefID].Snippets = make(map[spdx.ElementID]*spdx.Snippet2_2)
			}
			doc.UnpackagedFiles[fileID.ElementRefID].Snippets[eID] = snippet
		}

	}
	return err
}
