// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdxlib

import (
	"testing"

	"github.com/spdx/tools-golang/common/spdx"
	v2_12 "github.com/spdx/tools-golang/v2/v2_1"
	v2_22 "github.com/spdx/tools-golang/v2/v2_2"
)

// ===== 2.1 tests =====

func Test2_1CanGetIDsOfDescribedPackages(t *testing.T) {
	// set up document and some packages and relationships
	doc := &v2_12.Document{
		SPDXVersion:    "SPDX-2.1",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_12.CreationInfo{},
		Packages: []*v2_12.Package{
			{PackageName: "pkg1", PackageSPDXIdentifier: "p1"},
			{PackageName: "pkg2", PackageSPDXIdentifier: "p2"},
			{PackageName: "pkg3", PackageSPDXIdentifier: "p3"},
			{PackageName: "pkg4", PackageSPDXIdentifier: "p4"},
			{PackageName: "pkg5", PackageSPDXIdentifier: "p5"},
		},
		Relationships: []*v2_12.Relationship{
			&v2_12.Relationship{
				RefA:         spdx.MakeDocElementID("", "DOCUMENT"),
				RefB:         spdx.MakeDocElementID("", "p1"),
				Relationship: "DESCRIBES",
			},
			&v2_12.Relationship{
				RefA:         spdx.MakeDocElementID("", "DOCUMENT"),
				RefB:         spdx.MakeDocElementID("", "p5"),
				Relationship: "DESCRIBES",
			},
			// inverse relationship -- should also get detected
			&v2_12.Relationship{
				RefA:         spdx.MakeDocElementID("", "p4"),
				RefB:         spdx.MakeDocElementID("", "DOCUMENT"),
				Relationship: "DESCRIBED_BY",
			},
			// different relationship
			&v2_12.Relationship{
				RefA:         spdx.MakeDocElementID("", "p1"),
				RefB:         spdx.MakeDocElementID("", "p2"),
				Relationship: "DEPENDS_ON",
			},
		},
	}

	// request IDs for DESCRIBES / DESCRIBED_BY relationships
	describedPkgIDs, err := GetDescribedPackageIDs2_1(doc)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	// should be three of the five IDs, returned in alphabetical order
	if len(describedPkgIDs) != 3 {
		t.Fatalf("expected %d packages, got %d", 3, len(describedPkgIDs))
	}
	if describedPkgIDs[0] != spdx.ElementID("p1") {
		t.Errorf("expected %v, got %v", spdx.ElementID("p1"), describedPkgIDs[0])
	}
	if describedPkgIDs[1] != spdx.ElementID("p4") {
		t.Errorf("expected %v, got %v", spdx.ElementID("p4"), describedPkgIDs[1])
	}
	if describedPkgIDs[2] != spdx.ElementID("p5") {
		t.Errorf("expected %v, got %v", spdx.ElementID("p5"), describedPkgIDs[2])
	}
}

func Test2_1GetDescribedPackagesReturnsSinglePackageIfOnlyOne(t *testing.T) {
	// set up document and one package, but no relationships
	// b/c only one package
	doc := &v2_12.Document{
		SPDXVersion:    "SPDX-2.1",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_12.CreationInfo{},
		Packages: []*v2_12.Package{
			{PackageName: "pkg1", PackageSPDXIdentifier: "p1"},
		},
	}

	// request IDs for DESCRIBES / DESCRIBED_BY relationships
	describedPkgIDs, err := GetDescribedPackageIDs2_1(doc)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	// should return the single package
	if len(describedPkgIDs) != 1 {
		t.Fatalf("expected %d package, got %d", 1, len(describedPkgIDs))
	}
	if describedPkgIDs[0] != spdx.ElementID("p1") {
		t.Errorf("expected %v, got %v", spdx.ElementID("p1"), describedPkgIDs[0])
	}
}

func Test2_1FailsToGetDescribedPackagesIfMoreThanOneWithoutDescribesRelationship(t *testing.T) {
	// set up document and multiple packages, but no DESCRIBES relationships
	doc := &v2_12.Document{
		SPDXVersion:    "SPDX-2.1",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_12.CreationInfo{},
		Packages: []*v2_12.Package{
			{PackageName: "pkg1", PackageSPDXIdentifier: "p1"},
			{PackageName: "pkg2", PackageSPDXIdentifier: "p2"},
			{PackageName: "pkg3", PackageSPDXIdentifier: "p3"},
			{PackageName: "pkg4", PackageSPDXIdentifier: "p4"},
			{PackageName: "pkg5", PackageSPDXIdentifier: "p5"},
		},
		Relationships: []*v2_12.Relationship{
			// different relationship
			&v2_12.Relationship{
				RefA:         spdx.MakeDocElementID("", "p1"),
				RefB:         spdx.MakeDocElementID("", "p2"),
				Relationship: "DEPENDS_ON",
			},
		},
	}

	_, err := GetDescribedPackageIDs2_1(doc)
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}

func Test2_1FailsToGetDescribedPackagesIfMoreThanOneWithNilRelationships(t *testing.T) {
	// set up document and multiple packages, but no relationships slice
	doc := &v2_12.Document{
		SPDXVersion:    "SPDX-2.1",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_12.CreationInfo{},
		Packages: []*v2_12.Package{
			{PackageName: "pkg1", PackageSPDXIdentifier: "p1"},
			{PackageName: "pkg2", PackageSPDXIdentifier: "p2"},
		},
	}

	_, err := GetDescribedPackageIDs2_1(doc)
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}

func Test2_1FailsToGetDescribedPackagesIfZeroPackagesInMap(t *testing.T) {
	// set up document but no packages
	doc := &v2_12.Document{
		SPDXVersion:    "SPDX-2.1",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_12.CreationInfo{},
		Packages:       []*v2_12.Package{},
	}

	_, err := GetDescribedPackageIDs2_1(doc)
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}

func Test2_1FailsToGetDescribedPackagesIfNilMap(t *testing.T) {
	// set up document but no packages
	doc := &v2_12.Document{
		SPDXVersion:    "SPDX-2.1",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_12.CreationInfo{},
	}

	_, err := GetDescribedPackageIDs2_1(doc)
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}

// ===== 2.2 tests =====

func Test2_2CanGetIDsOfDescribedPackages(t *testing.T) {
	// set up document and some packages and relationships
	doc := &v2_22.Document{
		SPDXVersion:    "SPDX-2.2",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_22.CreationInfo{},
		Packages: []*v2_22.Package{
			{PackageName: "pkg1", PackageSPDXIdentifier: "p1"},
			{PackageName: "pkg2", PackageSPDXIdentifier: "p2"},
			{PackageName: "pkg3", PackageSPDXIdentifier: "p3"},
			{PackageName: "pkg4", PackageSPDXIdentifier: "p4"},
			{PackageName: "pkg5", PackageSPDXIdentifier: "p5"},
		},
		Relationships: []*v2_22.Relationship{
			&v2_22.Relationship{
				RefA:         spdx.MakeDocElementID("", "DOCUMENT"),
				RefB:         spdx.MakeDocElementID("", "p1"),
				Relationship: "DESCRIBES",
			},
			&v2_22.Relationship{
				RefA:         spdx.MakeDocElementID("", "DOCUMENT"),
				RefB:         spdx.MakeDocElementID("", "p5"),
				Relationship: "DESCRIBES",
			},
			// inverse relationship -- should also get detected
			&v2_22.Relationship{
				RefA:         spdx.MakeDocElementID("", "p4"),
				RefB:         spdx.MakeDocElementID("", "DOCUMENT"),
				Relationship: "DESCRIBED_BY",
			},
			// different relationship
			&v2_22.Relationship{
				RefA:         spdx.MakeDocElementID("", "p1"),
				RefB:         spdx.MakeDocElementID("", "p2"),
				Relationship: "DEPENDS_ON",
			},
		},
	}

	// request IDs for DESCRIBES / DESCRIBED_BY relationships
	describedPkgIDs, err := GetDescribedPackageIDs2_2(doc)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	// should be three of the five IDs, returned in alphabetical order
	if len(describedPkgIDs) != 3 {
		t.Fatalf("expected %d packages, got %d", 3, len(describedPkgIDs))
	}
	if describedPkgIDs[0] != spdx.ElementID("p1") {
		t.Errorf("expected %v, got %v", spdx.ElementID("p1"), describedPkgIDs[0])
	}
	if describedPkgIDs[1] != spdx.ElementID("p4") {
		t.Errorf("expected %v, got %v", spdx.ElementID("p4"), describedPkgIDs[1])
	}
	if describedPkgIDs[2] != spdx.ElementID("p5") {
		t.Errorf("expected %v, got %v", spdx.ElementID("p5"), describedPkgIDs[2])
	}
}

func Test2_2GetDescribedPackagesReturnsSinglePackageIfOnlyOne(t *testing.T) {
	// set up document and one package, but no relationships
	// b/c only one package
	doc := &v2_22.Document{
		SPDXVersion:    "SPDX-2.2",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_22.CreationInfo{},
		Packages: []*v2_22.Package{
			{PackageName: "pkg1", PackageSPDXIdentifier: "p1"},
		},
	}

	// request IDs for DESCRIBES / DESCRIBED_BY relationships
	describedPkgIDs, err := GetDescribedPackageIDs2_2(doc)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	// should return the single package
	if len(describedPkgIDs) != 1 {
		t.Fatalf("expected %d package, got %d", 1, len(describedPkgIDs))
	}
	if describedPkgIDs[0] != spdx.ElementID("p1") {
		t.Errorf("expected %v, got %v", spdx.ElementID("p1"), describedPkgIDs[0])
	}
}

func Test2_2FailsToGetDescribedPackagesIfMoreThanOneWithoutDescribesRelationship(t *testing.T) {
	// set up document and multiple packages, but no DESCRIBES relationships
	doc := &v2_22.Document{
		SPDXVersion:    "SPDX-2.2",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_22.CreationInfo{},
		Packages: []*v2_22.Package{
			{PackageName: "pkg1", PackageSPDXIdentifier: "p1"},
			{PackageName: "pkg2", PackageSPDXIdentifier: "p2"},
			{PackageName: "pkg3", PackageSPDXIdentifier: "p3"},
			{PackageName: "pkg4", PackageSPDXIdentifier: "p4"},
			{PackageName: "pkg5", PackageSPDXIdentifier: "p5"},
		},
		Relationships: []*v2_22.Relationship{
			// different relationship
			&v2_22.Relationship{
				RefA:         spdx.MakeDocElementID("", "p1"),
				RefB:         spdx.MakeDocElementID("", "p2"),
				Relationship: "DEPENDS_ON",
			},
		},
	}

	_, err := GetDescribedPackageIDs2_2(doc)
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}

func Test2_2FailsToGetDescribedPackagesIfMoreThanOneWithNilRelationships(t *testing.T) {
	// set up document and multiple packages, but no relationships slice
	doc := &v2_22.Document{
		SPDXVersion:    "SPDX-2.2",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_22.CreationInfo{},
		Packages: []*v2_22.Package{
			{PackageName: "pkg1", PackageSPDXIdentifier: "p1"},
			{PackageName: "pkg2", PackageSPDXIdentifier: "p2"},
		},
	}

	_, err := GetDescribedPackageIDs2_2(doc)
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}

func Test2_2FailsToGetDescribedPackagesIfZeroPackagesInMap(t *testing.T) {
	// set up document but no packages
	doc := &v2_22.Document{
		SPDXVersion:    "SPDX-2.2",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_22.CreationInfo{},
		Packages:       []*v2_22.Package{},
	}

	_, err := GetDescribedPackageIDs2_2(doc)
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}

func Test2_2FailsToGetDescribedPackagesIfNilMap(t *testing.T) {
	// set up document but no packages
	doc := &v2_22.Document{
		SPDXVersion:    "SPDX-2.2",
		DataLicense:    "CC0-1.0",
		SPDXIdentifier: spdx.ElementID("DOCUMENT"),
		CreationInfo:   &v2_22.CreationInfo{},
	}

	_, err := GetDescribedPackageIDs2_2(doc)
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}
