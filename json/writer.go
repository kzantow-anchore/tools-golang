// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdx_json

import (
	"encoding/json"
	"io"

	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/v2_2"
)

// Save2_2 takes an SPDX Document (version 2.2) and an io.Writer, and writes the document to the writer in JSON format.
func Save2_2(doc *v2_2.Document, w io.Writer) error {
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

// Save2_3 takes an SPDX Document (version 2.2) and an io.Writer, and writes the document to the writer in JSON format.
func Save2_3(doc *spdx.Document, w io.Writer) error {
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
