// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package reader

import (
	"testing"

	"github.com/spdx/tools-golang/v2/v2_2"
)

// ===== Annotation section tests =====
func TestParser2_2FailsIfAnnotationNotSet(t *testing.T) {
	parser := tvParser2_2{
		doc: &v2_2.Document{},
		st:  psCreationInfo2_2,
	}
	err := parser.parsePairForAnnotation("Annotator", "Person: John Doe (jdoe@example.com)")
	if err == nil {
		t.Errorf("expected error when calling parsePairFromAnnotation2_2 without setting ann pointer")
	}
}

func TestParser2_2FailsIfAnnotationTagUnknown(t *testing.T) {
	parser := tvParser2_2{
		doc: &v2_2.Document{},
		st:  psCreationInfo2_2,
	}
	// start with valid annotator
	err := parser.parsePair("Annotator", "Person: John Doe (jdoe@example.com)")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	// parse invalid tag, using parsePairForAnnotation(
	err = parser.parsePairForAnnotation("blah", "oops")
	if err == nil {
		t.Errorf("expected non-nil error, got nil")
	}
}

func TestParser2_2FailsIfAnnotationFieldsWithoutAnnotation(t *testing.T) {
	parser := tvParser2_2{
		doc: &v2_2.Document{},
		st:  psCreationInfo2_2,
	}
	err := parser.parsePair("AnnotationDate", "2018-09-15T17:25:00Z")
	if err == nil {
		t.Errorf("expected error when calling parsePair2_2 for AnnotationDate without Annotator first")
	}
	err = parser.parsePair("AnnotationType", "REVIEW")
	if err == nil {
		t.Errorf("expected error when calling parsePair2_2 for AnnotationType without Annotator first")
	}
	err = parser.parsePair("SPDXREF", "SPDXRef-45")
	if err == nil {
		t.Errorf("expected error when calling parsePair2_2 for SPDXREF without Annotator first")
	}
	err = parser.parsePair("AnnotationComment", "comment whatever")
	if err == nil {
		t.Errorf("expected error when calling parsePair2_2 for AnnotationComment without Annotator first")
	}
}

func TestParser2_2CanParseAnnotationTags(t *testing.T) {
	parser := tvParser2_2{
		doc: &v2_2.Document{},
		st:  psCreationInfo2_2,
	}

	// Annotator without email address
	err := parser.parsePair("Annotator", "Person: John Doe")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if parser.ann.Annotator.Annotator != "John Doe" {
		t.Errorf("got %+v for Annotator, expected John Doe", parser.ann.Annotator.Annotator)
	}
	if parser.ann.Annotator.AnnotatorType != "Person" {
		t.Errorf("got %v for AnnotatorType, expected Person", parser.ann.Annotator.AnnotatorType)
	}

	// Annotation Date
	dt := "2018-09-15T17:32:00Z"
	err = parser.parsePair("AnnotationDate", dt)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if parser.ann.AnnotationDate != dt {
		t.Errorf("got %v for AnnotationDate, expected %v", parser.ann.AnnotationDate, dt)
	}

	// Annotation type
	aType := "REVIEW"
	err = parser.parsePair("AnnotationType", aType)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if parser.ann.AnnotationType != aType {
		t.Errorf("got %v for AnnotationType, expected %v", parser.ann.AnnotationType, aType)
	}

	// SPDX Identifier Reference
	ref := "SPDXRef-30"
	err = parser.parsePair("SPDXREF", ref)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	deID := parser.ann.AnnotationSPDXIdentifier
	if deID.DocumentRefID != "" || deID.ElementRefID != "30" {
		t.Errorf("got %v for SPDXREF, expected %v", parser.ann.AnnotationSPDXIdentifier, "30")
	}

	// Annotation Comment
	cmt := "this is a comment"
	err = parser.parsePair("AnnotationComment", cmt)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if parser.ann.AnnotationComment != cmt {
		t.Errorf("got %v for AnnotationComment, expected %v", parser.ann.AnnotationComment, cmt)
	}
}

func TestParser2_2FailsIfAnnotatorInvalid(t *testing.T) {
	parser := tvParser2_2{
		doc: &v2_2.Document{},
		st:  psCreationInfo2_2,
	}
	err := parser.parsePair("Annotator", "John Doe (jdoe@example.com)")
	if err == nil {
		t.Errorf("expected non-nil error, got nil")
	}
}

func TestParser2_2FailsIfAnnotatorTypeInvalid(t *testing.T) {
	parser := tvParser2_2{
		doc: &v2_2.Document{},
		st:  psCreationInfo2_2,
	}
	err := parser.parsePair("Annotator", "Human: John Doe (jdoe@example.com)")
	if err == nil {
		t.Errorf("expected non-nil error, got nil")
	}
}

func TestParser2_2FailsIfAnnotationRefInvalid(t *testing.T) {
	parser := tvParser2_2{
		doc: &v2_2.Document{},
		st:  psCreationInfo2_2,
	}
	// start with valid annotator
	err := parser.parsePair("Annotator", "Person: John Doe (jdoe@example.com)")
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	err = parser.parsePair("SPDXREF", "blah:other")
	if err == nil {
		t.Errorf("expected non-nil error, got nil")
	}
}
