// Package builder is used to create tools-golang data structures for a given
// directory path's contents, with hashes, etc. filled in and with empty
// license data.
// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package builder

import (
	"fmt"

	"github.com/spdx/tools-golang/common/spdx"
	v2_12 "github.com/spdx/tools-golang/v2/v2_1"
	v2_22 "github.com/spdx/tools-golang/v2/v2_2"
	"github.com/spdx/tools-golang/v2/v2_3"

	"github.com/spdx/tools-golang/builder/builder2v1"
	"github.com/spdx/tools-golang/builder/builder2v2"
	"github.com/spdx/tools-golang/builder/builder2v3"
)

// ===== 2.1 builder =====

// Config2_1 is a collection of configuration settings for builder
// (for version 2.1 SPDX Documents). A few mandatory fields are set here
// so that they can be repeatedly reused in multiple calls to Build2_1.
type Config2_1 struct {
	// NamespacePrefix should be a URI representing a prefix for the
	// namespace with which the SPDX Document will be associated.
	// It will be used in the DocumentNamespace field in the CreationInfo
	// section, followed by the per-Document package name and a random UUID.
	NamespacePrefix string

	// CreatorType should be one of "Person", "Organization" or "Tool".
	// If not one of those strings, it will be interpreted as "Person".
	CreatorType string

	// Creator will be filled in for the given CreatorType.
	Creator string

	// PathsIgnored lists certain paths to be omitted from the built document.
	// Each string should be a path, relative to the package's dirRoot,
	// to a specific file or (for all files in a directory) ending in a slash.
	// Prefix the string with "**" to omit all instances of that file /
	// directory, regardless of where it is in the file tree.
	PathsIgnored []string

	// TestValues is used to pass fixed values for testing purposes
	// only, and should be set to nil for production use. It is only
	// exported so that it will be accessible within builder2v1.
	TestValues map[string]string
}

// Build2_1 creates an SPDX Document (version 2.1), returning that document or
// error if any is encountered. Arguments:
//   - packageName: name of package / directory
//   - dirRoot: path to directory to be analyzed
//   - config: Config object
func Build2_1(packageName string, dirRoot string, config *Config2_1) (*v2_12.Document, error) {
	// build Package section first -- will include Files and make the
	// package verification code available
	pkg, err := builder2v1.BuildPackageSection2_1(packageName, dirRoot, config.PathsIgnored)
	if err != nil {
		return nil, err
	}

	ci, err := builder2v1.BuildCreationInfoSection2_1(config.CreatorType, config.Creator, config.TestValues)
	if err != nil {
		return nil, err
	}

	rln, err := builder2v1.BuildRelationshipSection2_1(packageName)
	if err != nil {
		return nil, err
	}

	doc := &v2_12.Document{
		SPDXVersion:       "SPDX-2.1",
		DataLicense:       "CC0-1.0",
		SPDXIdentifier:    spdx.ElementID("DOCUMENT"),
		DocumentName:      packageName,
		DocumentNamespace: fmt.Sprintf("%s%s-%s", config.NamespacePrefix, packageName, pkg.PackageVerificationCode),
		CreationInfo:      ci,
		Packages:          []*v2_12.Package{pkg},
		Relationships:     []*v2_12.Relationship{rln},
	}

	return doc, nil
}

// ===== 2.2 builder =====

// Config2_2 is a collection of configuration settings for builder
// (for version 2.2 SPDX Documents). A few mandatory fields are set here
// so that they can be repeatedly reused in multiple calls to Build2_2.
type Config2_2 struct {
	// NamespacePrefix should be a URI representing a prefix for the
	// namespace with which the SPDX Document will be associated.
	// It will be used in the DocumentNamespace field in the CreationInfo
	// section, followed by the per-Document package name and a random UUID.
	NamespacePrefix string

	// CreatorType should be one of "Person", "Organization" or "Tool".
	// If not one of those strings, it will be interpreted as "Person".
	CreatorType string

	// Creator will be filled in for the given CreatorType.
	Creator string

	// PathsIgnored lists certain paths to be omitted from the built document.
	// Each string should be a path, relative to the package's dirRoot,
	// to a specific file or (for all files in a directory) ending in a slash.
	// Prefix the string with "**" to omit all instances of that file /
	// directory, regardless of where it is in the file tree.
	PathsIgnored []string

	// TestValues is used to pass fixed values for testing purposes
	// only, and should be set to nil for production use. It is only
	// exported so that it will be accessible within builder2v2.
	TestValues map[string]string
}

// Build2_2 creates an SPDX Document (version 2.2), returning that document or
// error if any is encountered. Arguments:
//   - packageName: name of package / directory
//   - dirRoot: path to directory to be analyzed
//   - config: Config object
func Build2_2(packageName string, dirRoot string, config *Config2_2) (*v2_22.Document, error) {
	// build Package section first -- will include Files and make the
	// package verification code available
	pkg, err := builder2v2.BuildPackageSection2_2(packageName, dirRoot, config.PathsIgnored)
	if err != nil {
		return nil, err
	}

	ci, err := builder2v2.BuildCreationInfoSection2_2(config.CreatorType, config.Creator, config.TestValues)
	if err != nil {
		return nil, err
	}

	rln, err := builder2v2.BuildRelationshipSection2_2(packageName)
	if err != nil {
		return nil, err
	}

	doc := &v2_22.Document{
		SPDXVersion:       "SPDX-2.2",
		DataLicense:       "CC0-1.0",
		SPDXIdentifier:    spdx.ElementID("DOCUMENT"),
		DocumentName:      packageName,
		DocumentNamespace: fmt.Sprintf("%s%s-%s", config.NamespacePrefix, packageName, pkg.PackageVerificationCode),
		CreationInfo:      ci,
		Packages:          []*v2_22.Package{pkg},
		Relationships:     []*v2_22.Relationship{rln},
	}

	return doc, nil
}

// ===== 2.3 builder =====

// Config2_3 is a collection of configuration settings for builder
// (for version 2.3 SPDX Documents). A few mandatory fields are set here
// so that they can be repeatedly reused in multiple calls to Build2_3.
type Config2_3 struct {
	// NamespacePrefix should be a URI representing a prefix for the
	// namespace with which the SPDX Document will be associated.
	// It will be used in the DocumentNamespace field in the CreationInfo
	// section, followed by the per-Document package name and a random UUID.
	NamespacePrefix string

	// CreatorType should be one of "Person", "Organization" or "Tool".
	// If not one of those strings, it will be interpreted as "Person".
	CreatorType string

	// Creator will be filled in for the given CreatorType.
	Creator string

	// PathsIgnored lists certain paths to be omitted from the built document.
	// Each string should be a path, relative to the package's dirRoot,
	// to a specific file or (for all files in a directory) ending in a slash.
	// Prefix the string with "**" to omit all instances of that file /
	// directory, regardless of where it is in the file tree.
	PathsIgnored []string

	// TestValues is used to pass fixed values for testing purposes
	// only, and should be set to nil for production use. It is only
	// exported so that it will be accessible within builder2v3.
	TestValues map[string]string
}

// Build2_3 creates an SPDX Document (version 2.3), returning that document or
// error if any is encountered. Arguments:
//   - packageName: name of package / directory
//   - dirRoot: path to directory to be analyzed
//   - config: Config object
func Build2_3(packageName string, dirRoot string, config *Config2_3) (*v2_3.Document, error) {
	// build Package section first -- will include Files and make the
	// package verification code available
	pkg, err := builder2v3.BuildPackageSection2_3(packageName, dirRoot, config.PathsIgnored)
	if err != nil {
		return nil, err
	}

	ci, err := builder2v3.BuildCreationInfoSection2_3(config.CreatorType, config.Creator, config.TestValues)
	if err != nil {
		return nil, err
	}

	rln, err := builder2v3.BuildRelationshipSection2_3(packageName)
	if err != nil {
		return nil, err
	}

	doc := &v2_3.Document{
		SPDXVersion:       "SPDX-2.3",
		DataLicense:       "CC0-1.0",
		SPDXIdentifier:    spdx.ElementID("DOCUMENT"),
		DocumentName:      packageName,
		DocumentNamespace: fmt.Sprintf("%s%s-%s", config.NamespacePrefix, packageName, pkg.PackageVerificationCode),
		CreationInfo:      ci,
		Packages:          []*v2_3.Package{pkg},
		Relationships:     []*v2_3.Relationship{rln},
	}

	return doc, nil
}
