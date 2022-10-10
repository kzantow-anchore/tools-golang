// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package reader

import (
	"github.com/spdx/tools-golang/common/spdx"
	v2_22 "github.com/spdx/tools-golang/v2/v2_2"
)

type tvParser2_2 struct {
	// document into which data is being parsed
	doc *v2_22.Document

	// current parser state
	st tvParserState2_2

	// current SPDX item being filled in, if any
	pkg       *v2_22.Package
	pkgExtRef *v2_22.PackageExternalReference
	file      *v2_22.File
	fileAOP   *v2_22.ArtifactOfProject
	snippet   *v2_22.Snippet
	otherLic  *v2_22.OtherLicense
	rln       *v2_22.Relationship
	ann       *v2_22.Annotation
	rev       *v2_22.Review
	// don't need creation info pointer b/c only one,
	// and we can get to it via doc.CreationInfo
}

// parser state (SPDX document version 2.2)
type tvParserState2_2 int

const (
	// at beginning of document
	psStart2_2 tvParserState2_2 = iota

	// in document creation info section
	psCreationInfo2_2

	// in package data section
	psPackage2_2

	// in file data section (including "unpackaged" files)
	psFile2_2

	// in snippet data section (including "unpackaged" files)
	psSnippet2_2

	// in other license section
	psOtherLicense2_2

	// in review section
	psReview2_2
)

const nullSpdxElementId2_2 = spdx.ElementID("")
