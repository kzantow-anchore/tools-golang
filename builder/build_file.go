// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package builder

import (
	"fmt"
	common2 "github.com/spdx/tools-golang/spdx/common"
	"path/filepath"

	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/utils"
)

// BuildFileSection creates an SPDX File, returning that
// file or error if any is encountered. Arguments:
//   - filePath: path to file, relative to prefix
//   - prefix: relative directory for filePath
//   - fileNumber: integer index (unique within package) to use in identifier
func BuildFileSection(filePath string, prefix string, fileNumber int) (*spdx.File, error) {
	// build the full file path
	p := filepath.Join(prefix, filePath)

	// make sure we can get the file and its hashes
	ssha1, ssha256, smd5, err := utils.GetHashesForFilePath(p)
	if err != nil {
		return nil, err
	}

	// build the identifier
	i := fmt.Sprintf("File%d", fileNumber)

	// now build the File section
	f := &spdx.File{
		FileName:           filePath,
		FileSPDXIdentifier: common2.ElementID(i),
		Checksums: []common2.Checksum{
			{
				Algorithm: common2.SHA1,
				Value:     ssha1,
			},
			{
				Algorithm: common2.SHA256,
				Value:     ssha256,
			},
			{
				Algorithm: common2.MD5,
				Value:     smd5,
			},
		},
		LicenseConcluded:   "NOASSERTION",
		LicenseInfoInFiles: []string{"NOASSERTION"},
		FileCopyrightText:  "NOASSERTION",
	}

	return f, nil
}
