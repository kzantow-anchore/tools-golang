// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package writer

import (
	"fmt"
	"io"

	"github.com/spdx/tools-golang/v2/v2_2"
)

func renderReview(rev *v2_2.Review, w io.Writer) error {
	if rev.Reviewer != "" && rev.ReviewerType != "" {
		fmt.Fprintf(w, "Reviewer: %s: %s\n", rev.ReviewerType, rev.Reviewer)
	}
	if rev.ReviewDate != "" {
		fmt.Fprintf(w, "ReviewDate: %s\n", rev.ReviewDate)
	}
	if rev.ReviewComment != "" {
		fmt.Fprintf(w, "ReviewComment: %s\n", textify(rev.ReviewComment))
	}

	fmt.Fprintf(w, "\n")

	return nil
}
