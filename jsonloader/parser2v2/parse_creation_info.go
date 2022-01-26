// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v2

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spdx/tools-golang/spdx"
)

func (spec JSONSpdxDocument) parseJsonCreationInfo2_2(key string, value interface{}, doc *spdxDocument2_2) (err error) {
	// create an SPDX Creation Info data struct if we don't have one already

	if doc.CreationInfo == nil {
		doc.CreationInfo = &spdx.CreationInfo2_2{
			ExternalDocumentReferences: map[string]spdx.ExternalDocumentRef2_2{},
		}
	}
	ci := doc.CreationInfo
	switch key {
	case "dataLicense":
		ci.DataLicense, err = requireString(value)
		if err != nil {
			return fmt.Errorf("invalid value for dataLicense, expected string but got: %+v", value)
		}
	case "spdxVersion":
		ci.SPDXVersion, err = requireString(value)
		if err != nil {
			return fmt.Errorf("invalid value for spdxVersion, expected string but got: %+v", value)
		}
	case "SPDXID":
		id, err := extractElementID(value)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		ci.SPDXIdentifier = id
	case "documentNamespace":
		ci.DocumentNamespace, err = requireString(value)
		if err != nil {
			return fmt.Errorf("invalid value for documentNamespace, expected string but got: %+v", value)
		}
	case "name":
		ci.DocumentName, err = requireString(value)
		if err != nil {
			return fmt.Errorf("invalid value for document name, expected string but got: %+v", value)
		}
	case "comment":
		ci.DocumentComment, err = requireString(value)
		if err != nil {
			return fmt.Errorf("invalid value for document comment, expected string but got: %+v", value)
		}
	case "creationInfo":
		creationInfo, err := requireMap(value)
		if err != nil {
			return fmt.Errorf("invalid value for document creationInfo, expected map[string] but got: %+v", value)
		}
		for key, val := range creationInfo {
			switch key {
			case "comment":
				ci.CreatorComment, err = requireString(val)
				if err != nil {
					return fmt.Errorf("invalid value for creationInfo.comment, expected string but got: %+v", val)
				}
			case "created":
				ci.Created, err = requireString(val)
				if err != nil {
					return fmt.Errorf("invalid value for creationInfo.created, expected string but got: %+v", val)
				}
			case "licenseListVersion":
				ci.LicenseListVersion, err = requireString(val)
				if err != nil {
					return fmt.Errorf("invalid value for creationInfo.licenseListVersion, expected string but got: %+v", val)
				}
			case "creators":
				err := parseCreators(creationInfo["creators"], ci)
				if err != nil {
					return fmt.Errorf("%s", err)
				}
			}
		}
	case "externalDocumentRefs":
		err := parseExternalDocumentRefs(value, ci)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
	default:
		return fmt.Errorf("unrecognized key %v", key)

	}

	return err
}

// ===== Helper functions =====

func parseCreators(creators interface{}, ci *spdx.CreationInfo2_2) error {
	if reflect.TypeOf(creators).Kind() == reflect.Slice {
		s := reflect.ValueOf(creators)

		for i := 0; i < s.Len(); i++ {
			subkey, subvalue, err := extractSubs(s.Index(i).Interface())
			if err != nil {
				return err
			}
			switch subkey {
			case "Person":
				ci.CreatorPersons = append(ci.CreatorPersons, strings.TrimSuffix(subvalue, " ()"))
			case "Organization":
				ci.CreatorOrganizations = append(ci.CreatorOrganizations, strings.TrimSuffix(subvalue, " ()"))
			case "Tool":
				ci.CreatorTools = append(ci.CreatorTools, subvalue)
			default:
				return fmt.Errorf("unrecognized Creator type: %s", subkey)
			}
		}
	}
	return nil
}

func parseExternalDocumentRefs(references interface{}, ci *spdx.CreationInfo2_2) (err error) {
	if reflect.TypeOf(references).Kind() == reflect.Slice {
		s := reflect.ValueOf(references)

		for i := 0; i < s.Len(); i++ {
			ifc := s.Index(i).Interface()
			ref, err := requireMap(ifc)
			if err != nil {
				return fmt.Errorf("invalid value for external document ref, expected map[string] but got: %+v", ifc)
			}
			documentRefID, err := requireMapString(ref, "externalDocumentId")
			if err != nil {
				return fmt.Errorf("error extracting externalDocumentId: %w", err)
			}
			if !strings.HasPrefix(documentRefID, "DocumentRef-") {
				return fmt.Errorf("expected first element to have DocumentRef- prefix")
			}
			documentRefID = strings.TrimPrefix(documentRefID, "DocumentRef-")
			if documentRefID == "" {
				return fmt.Errorf("document identifier has nothing after prefix")
			}
			checksum, err := requireMapMap(ref, "checksum")
			if err != nil {
				return fmt.Errorf("error extracting checksum: %w", err)
			}
			spdxDocument, err := requireMapString(ref, "spdxDocument")
			if err != nil {
				return fmt.Errorf("error extracting spdxDocument: %w", err)
			}
			algorithm, err := requireMapString(checksum, "algorithm")
			if err != nil {
				return fmt.Errorf("error extracting algorithm: %w", err)
			}
			checksumValue, err := requireMapString(checksum, "checksumValue")
			if err != nil {
				return fmt.Errorf("error extracting checksumValue: %w", err)
			}
			edr := spdx.ExternalDocumentRef2_2{
				DocumentRefID: documentRefID,
				URI:           spdxDocument,
				Alg:           algorithm,
				Checksum:      checksumValue,
			}

			ci.ExternalDocumentReferences[documentRefID] = edr
		}
	}
	return nil
}
