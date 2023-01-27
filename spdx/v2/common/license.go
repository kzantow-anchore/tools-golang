package common

import (
	"encoding/json"
	"fmt"
	"strings"

	tv "github.com/spdx/tools-golang/tagvalue/lib"
)

const licenseIDPrefix = "LicenseRef-"

type LicenseID string

func (d LicenseID) ToTagValue() (string, error) {
	return prependLicensePrefix(string(d)), nil
}

func (d *LicenseID) FromTagValue(s string) error {
	*d = LicenseID(trimLicensePrefix(s))
	return nil
}

var _ tv.ToValue = (*LicenseID)(nil)
var _ tv.FromValue = (*LicenseID)(nil)

// UnmarshalJSON takes an OtherLicense in the typical one-line format and parses it into an OtherLicense struct.
// This function is also used when unmarshalling YAML
func (d *LicenseID) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = strings.Trim(s, "\"")
	*d = LicenseID(trimLicensePrefix(s))
	return nil
}

// MarshalJSON converts the receiver into a slice of bytes representing an OtherLicense in string form.
// This function is also used when marshalling to YAML
func (d LicenseID) MarshalJSON() ([]byte, error) {
	return json.Marshal(prependLicensePrefix(string(d)))
}

func prependLicensePrefix(licenseID string) string {
	switch licenseID {
	case "NONE", "NOASSERTION":
	default:
		if !strings.HasPrefix(licenseID, licenseIDPrefix) {
			licenseID = fmt.Sprintf("%s%s", licenseIDPrefix, licenseID)
		}
	}
	return licenseID
}

func trimLicensePrefix(license string) string {
	return strings.TrimPrefix(license, licenseIDPrefix)
}
