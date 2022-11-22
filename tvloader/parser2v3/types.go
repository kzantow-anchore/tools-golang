// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v3

import (
	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/spdx/common"
)

type tvParser2_3 struct {
	// document into which data is being parsed
	doc *spdx.Document

	// current parser state
	st tvParserState2_3

	// current SPDX item being filled in, if any
	pkg       *spdx.Package
	pkgExtRef *spdx.PackageExternalReference
	file      *spdx.File
	fileAOP   *spdx.ArtifactOfProject
	snippet   *spdx.Snippet
	otherLic  *spdx.OtherLicense
	rln       *spdx.Relationship
	ann       *spdx.Annotation
	rev       *spdx.Review
	// don't need creation info pointer b/c only one,
	// and we can get to it via doc.CreationInfo
}

// parser state (SPDX document version 2.3)
type tvParserState2_3 int

const (
	// at beginning of document
	psStart2_3 tvParserState2_3 = iota

	// in document creation info section
	psCreationInfo2_3

	// in package data section
	psPackage2_3

	// in file data section (including "unpackaged" files)
	psFile2_3

	// in snippet data section (including "unpackaged" files)
	psSnippet2_3

	// in other license section
	psOtherLicense2_3

	// in review section
	psReview2_3
)

const nullSpdxElementId2_3 = common.ElementID("")
