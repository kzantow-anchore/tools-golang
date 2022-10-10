// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdx_yaml

import (
	"io"

	"sigs.k8s.io/yaml"

	"github.com/spdx/tools-golang/v2/v2_2"
)

// Write takes an SPDX Document and an io.Writer, and writes the document to the writer in YAML format.
func Write(doc *v2_2.Document, w io.Writer) error {
	buf, err := yaml.Marshal(doc)
	if err != nil {
		return err
	}

	_, err = w.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
