// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package common

import (
	"encoding/json"
	"fmt"
	"strings"

	tv "github.com/spdx/tools-golang/tagvalue/lib"
)

// Creator is a wrapper around the Creator SPDX field. The SPDX field contains two values, which requires special
// handling in order to marshal/unmarshal it to/from Go data types.
type Creator struct {
	Creator string
	// CreatorType should be one of "Person", "Organization", or "Tool"
	CreatorType string
}

func (d Creator) ToTagValue() (string, error) {
	return fmt.Sprintf("%s: %s", d.CreatorType, d.Creator), nil
}

func (d *Creator) FromTagValue(s string) error {
	parts := strings.Split(s, ": ")
	if len(parts) == 2 {
		d.CreatorType = parts[0]
		d.Creator = parts[1]
	}
	return nil
}

var _ tv.ToValue = (*Creator)(nil)
var _ tv.FromValue = (*Creator)(nil)

// UnmarshalJSON takes an annotator in the typical one-line format and parses it into a Creator struct.
// This function is also used when unmarshalling YAML
func (c *Creator) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.Trim(str, "\"")
	fields := strings.SplitN(str, ": ", 2)

	if len(fields) != 2 {
		return fmt.Errorf("failed to parse Creator '%s'", str)
	}

	c.CreatorType = fields[0]
	c.Creator = fields[1]

	return nil
}

// MarshalJSON converts the receiver into a slice of bytes representing a Creator in string form.
// This function is also used with marshalling to YAML
func (c Creator) MarshalJSON() ([]byte, error) {
	if c.Creator != "" {
		return json.Marshal(fmt.Sprintf("%s: %s", c.CreatorType, c.Creator))
	}

	return []byte{}, nil
}
