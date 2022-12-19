// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package writer

import (
	"bytes"
	common2 "github.com/spdx/tools-golang/spdx/common"
	"testing"

	"github.com/spdx/tools-golang/spdx"
)

// ===== Snippet section Saver tests =====
func TestSaverSnippetSavesText(t *testing.T) {
	sn := &spdx.Snippet{
		SnippetSPDXIdentifier:         common2.ElementID("Snippet17"),
		SnippetFromFileSPDXIdentifier: common2.MakeDocElementID("", "File292").ElementRefID,
		Ranges: []common2.SnippetRange{
			{
				StartPointer: common2.SnippetRangePointer{LineNumber: 3},
				EndPointer:   common2.SnippetRangePointer{LineNumber: 8},
			},
			{
				StartPointer: common2.SnippetRangePointer{Offset: 17},
				EndPointer:   common2.SnippetRangePointer{Offset: 209},
			},
		},
		SnippetLicenseConcluded: "GPL-2.0-or-later",
		LicenseInfoInSnippet: []string{
			"GPL-2.0-or-later",
			"MIT",
		},
		SnippetLicenseComments:  "this is a comment(s) about the snippet license",
		SnippetCopyrightText:    "Copyright (c) John Doe 20x6",
		SnippetComment:          "this is a snippet comment",
		SnippetName:             "from John's program",
		SnippetAttributionTexts: []string{"some attributions"},
	}

	// what we want to get, as a buffer of bytes
	want := bytes.NewBufferString(`SnippetSPDXID: SPDXRef-Snippet17
SnippetFromFileSPDXID: SPDXRef-File292
SnippetLineRange: 3:8
SnippetByteRange: 17:209
SnippetLicenseConcluded: GPL-2.0-or-later
LicenseInfoInSnippet: GPL-2.0-or-later
LicenseInfoInSnippet: MIT
SnippetLicenseComments: this is a comment(s) about the snippet license
SnippetCopyrightText: Copyright (c) John Doe 20x6
SnippetComment: this is a snippet comment
SnippetName: from John's program
SnippetAttributionText: some attributions

`)

	// render as buffer of bytes
	var got bytes.Buffer
	err := renderSnippet(sn, &got)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// check that they match
	c := bytes.Compare(want.Bytes(), got.Bytes())
	if c != 0 {
		t.Errorf("Expected %v, got %v", want.String(), got.String())
	}
}

func TestSaverSnippetOmitsOptionalFieldsIfEmpty(t *testing.T) {
	sn := &spdx.Snippet{
		SnippetSPDXIdentifier:         common2.ElementID("Snippet17"),
		SnippetFromFileSPDXIdentifier: common2.MakeDocElementID("", "File292").ElementRefID,
		Ranges: []common2.SnippetRange{
			{
				StartPointer: common2.SnippetRangePointer{Offset: 17},
				EndPointer:   common2.SnippetRangePointer{Offset: 209},
			},
		},
		SnippetLicenseConcluded: "GPL-2.0-or-later",
		SnippetCopyrightText:    "Copyright (c) John Doe 20x6",
	}

	// what we want to get, as a buffer of bytes
	want := bytes.NewBufferString(`SnippetSPDXID: SPDXRef-Snippet17
SnippetFromFileSPDXID: SPDXRef-File292
SnippetByteRange: 17:209
SnippetLicenseConcluded: GPL-2.0-or-later
SnippetCopyrightText: Copyright (c) John Doe 20x6

`)

	// render as buffer of bytes
	var got bytes.Buffer
	err := renderSnippet(sn, &got)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// check that they match
	c := bytes.Compare(want.Bytes(), got.Bytes())
	if c != 0 {
		t.Errorf("Expected %v, got %v", want.String(), got.String())
	}
}

func TestSaverSnippetWrapsCopyrightMultiline(t *testing.T) {
	sn := &spdx.Snippet{
		SnippetSPDXIdentifier:         common2.ElementID("Snippet17"),
		SnippetFromFileSPDXIdentifier: common2.MakeDocElementID("", "File292").ElementRefID,
		Ranges: []common2.SnippetRange{
			{
				StartPointer: common2.SnippetRangePointer{Offset: 17},
				EndPointer:   common2.SnippetRangePointer{Offset: 209},
			},
		},
		SnippetLicenseConcluded: "GPL-2.0-or-later",
		SnippetCopyrightText: `Copyright (c) John Doe 20x6
Copyright (c) John Doe 20x6`,
	}

	// what we want to get, as a buffer of bytes
	want := bytes.NewBufferString(`SnippetSPDXID: SPDXRef-Snippet17
SnippetFromFileSPDXID: SPDXRef-File292
SnippetByteRange: 17:209
SnippetLicenseConcluded: GPL-2.0-or-later
SnippetCopyrightText: <text>Copyright (c) John Doe 20x6
Copyright (c) John Doe 20x6</text>

`)

	// render as buffer of bytes
	var got bytes.Buffer
	err := renderSnippet(sn, &got)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	// check that they match
	c := bytes.Compare(want.Bytes(), got.Bytes())
	if c != 0 {
		t.Errorf("Expected %v, got %v", want.String(), got.String())
	}
}
