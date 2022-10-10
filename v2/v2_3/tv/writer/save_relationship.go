// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package writer

import (
	"fmt"
	"io"

	"github.com/spdx/tools-golang/common/spdx"
	"github.com/spdx/tools-golang/v2/v2_3"
)

func renderRelationship(rln *v2_3.Relationship, w io.Writer) error {
	rlnAStr := spdx.RenderDocElementID(rln.RefA)
	rlnBStr := spdx.RenderDocElementID(rln.RefB)
	if rlnAStr != "SPDXRef-" && rlnBStr != "SPDXRef-" && rln.Relationship != "" {
		fmt.Fprintf(w, "Relationship: %s %s %s\n", rlnAStr, rln.Relationship, rlnBStr)
	}
	if rln.RelationshipComment != "" {
		fmt.Fprintf(w, "RelationshipComment: %s\n", textify(rln.RelationshipComment))
	}

	return nil
}
