// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package v2_3

import (
	"fmt"

	"github.com/spdx/tools-golang/spdx/v2/common"
	tv "github.com/spdx/tools-golang/tagvalue/lib"
)

// File is a File section of an SPDX Document
type File struct {
	// 8.1: File Name
	// Cardinality: mandatory, one
	FileName string `json:"fileName"`

	// 8.2: File SPDX Identifier: "SPDXRef-[idstring]"
	// Cardinality: mandatory, one
	FileSPDXIdentifier common.ElementID `json:"SPDXID" tv:"SPDXID"`

	// 8.3: File Types
	// Cardinality: optional, multiple
	FileTypes []string `json:"fileTypes,omitempty" tv:"FileType"`

	// 8.4: File Checksum: may have keys for SHA1, SHA256, MD5, SHA3-256, SHA3-384, SHA3-512, BLAKE2b-256, BLAKE2b-384, BLAKE2b-512, BLAKE3, ADLER32
	// Cardinality: mandatory, one SHA1, others may be optionally provided
	Checksums []common.Checksum `json:"checksums" tv:"FileChecksum"`

	// 8.5: Concluded License: SPDX License Expression, "NONE" or "NOASSERTION"
	// Cardinality: optional, one
	LicenseConcluded string `json:"licenseConcluded,omitempty"`

	// 8.6: License Information in File: SPDX License Expression, "NONE" or "NOASSERTION"
	// Cardinality: optional, one or many
	LicenseInfoInFiles []string `json:"licenseInfoInFiles,omitempty" tv:"LicenseInfoInFile"`

	// 8.7: Comments on License
	// Cardinality: optional, one
	LicenseComments string `json:"licenseComments,omitempty"`

	// 8.8: Copyright Text: copyright notice(s) text, "NONE" or "NOASSERTION"
	// Cardinality: mandatory, one
	FileCopyrightText string `json:"copyrightText"`

	// DEPRECATED in version 2.1 of spec
	// 8.9-8.11: Artifact of Project variables (defined below)
	// Cardinality: optional, one or many
	ArtifactOfProjects []*ArtifactOfProject `json:"artifactOfs,omitempty" tv:"ArtifactOfProjectName"`

	// 8.12: File Comment
	// Cardinality: optional, one
	FileComment string `json:"comment,omitempty"`

	// 8.13: File Notice
	// Cardinality: optional, one
	FileNotice string `json:"noticeText,omitempty"`

	// 8.14: File Contributor
	// Cardinality: optional, one or many
	FileContributors []string `json:"fileContributors,omitempty" tv:"FileContributor"`

	// 8.15: File Attribution Text
	// Cardinality: optional, one or many
	FileAttributionTexts []string `json:"attributionTexts,omitempty" tv:"FileAttributionText"`

	// DEPRECATED in version 2.0 of spec
	// 8.16: File Dependencies
	// Cardinality: optional, one or many
	FileDependencies []string `json:"fileDependencies,omitempty" tv:"FileDependency"`

	// Snippets contained in this File
	// Note that Snippets could be defined in a different Document! However,
	// the only ones that _THIS_ document can contain are this ones that are
	// defined here -- so this should just be an ElementID.
	Snippets map[common.ElementID]*Snippet `json:"-" yaml:"-"`

	Annotations []Annotation `json:"annotations,omitempty"`
}

func (f File) TagValuePrefix() string {
	return fmt.Sprintf("\n# ----- File: %s -----\n", f.FileName)
}

var _ tv.TagValuePrefix = (*File)(nil)

// ArtifactOfProject is a DEPRECATED collection of data regarding
// a Package, as defined in sections 8.9-8.11.
// NOTE: the JSON schema does not define the structure of this object:
// https://github.com/spdx/spdx-spec/blob/development/v2.3.1/schemas/spdx-schema.json#L480
type ArtifactOfProject struct {

	// DEPRECATED in version 2.1 of spec
	// 8.9: Artifact of Project Name
	// Cardinality: conditional, required if present, one per AOP
	Name string `json:"name" tv:"ArtifactOfProjectName"`

	// DEPRECATED in version 2.1 of spec
	// 8.10: Artifact of Project Homepage: URL or "UNKNOWN"
	// Cardinality: optional, one per AOP
	HomePage string `json:"homePage" tv:"ArtifactOfProjectHomePage"`

	// DEPRECATED in version 2.1 of spec
	// 8.11: Artifact of Project Uniform Resource Identifier
	// Cardinality: optional, one per AOP
	URI string `json:"URI" tv:"ArtifactOfProjectURI"`
}
