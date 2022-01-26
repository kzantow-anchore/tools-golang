// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v2

import (
	"fmt"
	"strings"

	"github.com/spdx/tools-golang/spdx"
)

// used to extract key / value from embedded substrings
// returns subkey, subvalue, nil if no error, or "", "", error otherwise
func extractSubs(value interface{}) (string, string, error) {
	str, ok := value.(string)
	if !ok {
		return "", "", fmt.Errorf("invalid value for extractSubs, expected string but got: %+v", value)
	}
	// parse the value to see if it's a valid subvalue format
	sp := strings.SplitN(str, ":", 2)
	if len(sp) == 1 {
		return "", "", fmt.Errorf("invalid subvalue format for: %s (no colon found)", value)
	}

	subkey := strings.TrimSpace(sp[0])
	subvalue := strings.TrimSpace(sp[1])

	return subkey, subvalue, nil
}

// used to extract DocumentRef and SPDXRef values from an SPDX Identifier
// which can point either to this document or to a different one
func extractDocElementID(value interface{}) (spdx.DocElementID, error) {

	docRefID := ""
	idStr, ok := value.(string)
	if !ok {
		return spdx.DocElementID{}, fmt.Errorf("invalid document element ID, expected string but got: %+v", value)
	}

	// check prefix to see if it's a DocumentRef ID
	if strings.HasPrefix(idStr, "DocumentRef-") {
		// extract the part that comes between "DocumentRef-" and ":"
		strs := strings.Split(idStr, ":")
		// should be exactly two, part before and part after
		if len(strs) < 2 {
			return spdx.DocElementID{}, fmt.Errorf("no colon found although DocumentRef- prefix present")
		}
		if len(strs) > 2 {
			return spdx.DocElementID{}, fmt.Errorf("more than one colon found")
		}

		// trim the prefix and confirm non-empty
		docRefID = strings.TrimPrefix(strs[0], "DocumentRef-")
		if docRefID == "" {
			return spdx.DocElementID{}, fmt.Errorf("document identifier has nothing after prefix")
		}
		// and use remainder for element ID parsing
		idStr = strs[1]
	}

	// check prefix to confirm it's got the right prefix for element IDs
	if !strings.HasPrefix(idStr, "SPDXRef-") {
		return spdx.DocElementID{}, fmt.Errorf("missing SPDXRef- prefix for element identifier")
	}

	// make sure no colons are present
	if strings.Contains(idStr, ":") {
		// we know this means there was no DocumentRef- prefix, because
		// we would have handled multiple colons above if it was
		return spdx.DocElementID{}, fmt.Errorf("invalid colon in element identifier")
	}

	// trim the prefix and confirm non-empty
	eltRefID := strings.TrimPrefix(idStr, "SPDXRef-")
	if eltRefID == "" {
		return spdx.DocElementID{}, fmt.Errorf("element identifier has nothing after prefix")
	}

	// we're good
	return spdx.DocElementID{DocumentRefID: docRefID, ElementRefID: spdx.ElementID(eltRefID)}, nil
}

// used to extract SPDXRef values from an SPDX Identifier, OR "special" strings
// from a specified set of permitted values. The primary use case for this is
// the right-hand side of Relationships, where beginning in SPDX 2.2 the values
// "NONE" and "NOASSERTION" are permitted. If the value does not match one of
// the specified permitted values, it will fall back to the ordinary
// DocElementID extractor.
func extractDocElementSpecial(value string, permittedSpecial []string) (spdx.DocElementID, error) {
	// check value against special set first
	for _, sp := range permittedSpecial {
		if sp == value {
			return spdx.DocElementID{SpecialID: sp}, nil
		}
	}
	// not found, fall back to regular search
	return extractDocElementID(value)
}

// used to extract SPDXRef values only from an SPDX Identifier which can point
// to this document only. Use extractDocElementID for parsing IDs that can
// refer either to this document or a different one.
func extractElementID(value interface{}) (spdx.ElementID, error) {
	str, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("invalid data type for SPDX Element ID, required string but got: %+v", value)
	}
	// check prefix to confirm it's got the right prefix for element IDs
	if !strings.HasPrefix(str, "SPDXRef-") {
		return spdx.ElementID(""), fmt.Errorf("missing SPDXRef- prefix for element identifier")
	}

	// make sure no colons are present
	if strings.Contains(str, ":") {
		return spdx.ElementID(""), fmt.Errorf("invalid colon in element identifier")
	}

	// trim the prefix and confirm non-empty
	eltRefID := strings.TrimPrefix(str, "SPDXRef-")
	if eltRefID == "" {
		return spdx.ElementID(""), fmt.Errorf("element identifier has nothing after prefix")
	}

	// we're good
	return spdx.ElementID(eltRefID), nil
}

// Requires the provided interface be a string value and returns it or an error if the cast fails
func requireString(v interface{}) (string, error) {
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("string required but got invalid value: %+v", v)
	}
	return s, nil
}

// Requires the provided interface be a map[string]interface{} value and returns it or an error if the cast fails
func requireMap(v interface{}) (map[string]interface{}, error) {
	s, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("map[string]interface{} required but got invalid value: %+v", v)
	}
	return s, nil
}

// Extracts a named string value from the provided map using the provided key
func requireMapString(v map[string]interface{}, key string) (string, error) {
	val, ok := v[key]
	if !ok {
		return "", fmt.Errorf("key not found in map: %s", key)
	}
	s, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("expected string value from map but got: %+v", val)
	}
	return s, nil
}

// Extracts a named map[string] value from the provided map using the provided key
func requireMapMap(v map[string]interface{}, key string) (map[string]interface{}, error) {
	val, ok := v[key]
	if !ok {
		return nil, fmt.Errorf("key not found in map: %s", key)
	}
	s, ok := val.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected map[string]interface{} value from map but got: %+v", val)
	}
	return s, nil
}

// Requires the map key with the given name of a float64 type to be present and casts it to an int
func requireMapFloatInt(v map[string]interface{}, key string) (int, error) {
	val, ok := v[key]
	if !ok {
		return 0, fmt.Errorf("key not found in map: %s", key)
	}
	f, ok := val.(float64)
	if !ok {
		return 0, fmt.Errorf("expected float64 value from map but got: %+v", val)
	}
	return int(f), nil
}
