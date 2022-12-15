// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdx_json

import (
	"encoding/json"
	"io"

	"github.com/spdx/tools-golang/common"
)

// Write takes an SPDX Document and an io.Writer, and writes the document to the writer in JSON format.
func Write(doc common.Document, w io.Writer) error {
	buf, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
