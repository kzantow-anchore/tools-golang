// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package builder2v2

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/spdx/tools-golang/common/spdx"
	v2_22 "github.com/spdx/tools-golang/v2/v2_2"

	"github.com/spdx/tools-golang/utils"
)

// BuildPackageSection2_2 creates an SPDX Package (version 2.2), returning
// that package or error if any is encountered. Arguments:
//   - packageName: name of package / directory
//   - dirRoot: path to directory to be analyzed
//   - pathsIgnore: slice of strings for filepaths to ignore
func BuildPackageSection2_2(packageName string, dirRoot string, pathsIgnore []string) (*v2_22.Package, error) {
	// build the file section first, so we'll have it available
	// for calculating the package verification code
	filepaths, err := utils.GetAllFilePaths(dirRoot, pathsIgnore)
	osType := runtime.GOOS

	if err != nil {
		return nil, err
	}

	re, ok := regexp.Compile("/+")
	if ok != nil {
		return nil, err
	}
	dirRootLen := 0
	if osType == "windows" {
		dirRootLen = len(dirRoot)
	}

	files := []*v2_22.File{}
	fileNumber := 0
	for _, fp := range filepaths {
		newFilePatch := ""
		if osType == "windows" {
			newFilePatch = filepath.FromSlash("." + fp[dirRootLen:])
		} else {
			newFilePatch = filepath.FromSlash("./" + fp)
		}
		newFile, err := BuildFileSection2_2(re.ReplaceAllLiteralString(newFilePatch, string(filepath.Separator)), dirRoot, fileNumber)
		if err != nil {
			return nil, err
		}
		files = append(files, newFile)
		fileNumber++
	}

	// get the verification code
	code, err := utils.GetVerificationCode2_2(files, "")
	if err != nil {
		return nil, err
	}

	// now build the package section
	pkg := &v2_22.Package{
		PackageName:                 packageName,
		PackageSPDXIdentifier:       spdx.ElementID(fmt.Sprintf("Package-%s", packageName)),
		PackageDownloadLocation:     "NOASSERTION",
		FilesAnalyzed:               true,
		IsFilesAnalyzedTagPresent:   true,
		PackageVerificationCode:     code,
		PackageLicenseConcluded:     "NOASSERTION",
		PackageLicenseInfoFromFiles: []string{},
		PackageLicenseDeclared:      "NOASSERTION",
		PackageCopyrightText:        "NOASSERTION",
		Files:                       files,
	}

	return pkg, nil
}
