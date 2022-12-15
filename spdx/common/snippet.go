// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package common

import (
	"fmt"

	tv "github.com/spdx/tools-golang/tagvalue/lib"
)

type SnippetRangePointer struct {
	// 5.3: Snippet Byte Range: [start byte]:[end byte]
	// Cardinality: mandatory, one
	Offset int `json:"offset,omitempty"`

	// 5.4: Snippet Line Range: [start line]:[end line]
	// Cardinality: optional, one
	LineNumber int `json:"lineNumber,omitempty"`

	FileSPDXIdentifier ElementID `json:"reference"`
}

type SnippetRange struct {
	StartPointer SnippetRangePointer `json:"startPointer"`
	EndPointer   SnippetRangePointer `json:"endPointer"`
}

type snippetTagValue struct {
	SnippetByteRange string
	SnippetLineRange string
}

func (d SnippetRange) GetTagValue() (interface{}, error) {
	return snippetTagValue{
		SnippetByteRange: fmt.Sprintf("%d:%d", d.StartPointer.Offset, d.EndPointer.Offset),
		SnippetLineRange: fmt.Sprintf("%d:%d", d.StartPointer.LineNumber, d.EndPointer.LineNumber),
	}, nil
}

func (d SnippetRange) FromTagValue(i interface{}) error {
	//TODO implement me
	panic("implement me")
}

var _ tv.TagValueHandler = (*SnippetRange)(nil)
