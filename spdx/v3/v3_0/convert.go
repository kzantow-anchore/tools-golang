package v3_0

import (
	converter "github.com/anchore/go-struct-converter"
	"github.com/spdx/tools-golang/spdx/v2/common"
	"github.com/spdx/tools-golang/spdx/v2/v2_3"
	"slices"
	"strings"
)

func (d *SpdxDocument) ConvertFrom(doc any) error {
	idMap := map[string]any{}

	if doc, ok := doc.(v2_3.Document); ok {
		d.SpdxId = string(doc.SPDXIdentifier)
		d.Comment = doc.DocumentComment
		d.CreationInfo = convert23CreationInfo(doc.CreationInfo)
		d.Name = doc.DocumentName
		d.DataLicense = &AnyLicenseInfo{
			Name: doc.DataLicense,
		}

		for _, pkg := range doc.Packages {
			newPkg := convert23Package(idMap, pkg)
			if newPkg != nil {
				d.RootElements = append(d.RootElements, newPkg)
				idMap[string(pkg.PackageSPDXIdentifier)] = newPkg
			}
		}

		for _, file := range doc.Files {
			newFile := convert23File(idMap, file)
			if newFile != nil {
				d.RootElements = append(d.RootElements, newFile)
				idMap[string(file.FileSPDXIdentifier)] = newFile
			}
		}

		rels := relMap{}
		for _, rel := range doc.Relationships {
			newRel := convert23Relationship(idMap, rel)
			if newRel != nil {
				rels.add(newRel)
			}
		}
		for _, relTypes := range rels {
			for _, to := range relTypes {
				d.RootElements = append(d.RootElements, to)
			}
		}
	}

	return nil
}

type relMap map[IElement]map[RelationshipType]*Relationship

func (r relMap) add(relationship *Relationship) {
	relTypes := r[relationship.GetFrom()]
	if relTypes == nil {
		relTypes = map[RelationshipType]*Relationship{}
		r[relationship.GetFrom()] = relTypes
	}
	existing := relTypes[relationship.RelationshipType]
	if existing == nil {
		relTypes[relationship.RelationshipType] = relationship
		return
	}
	existing.To = appendUnique(existing.To, relationship.To...)
}

func appendUnique[T comparable](existing []T, adding ...T) []T {
	for _, add := range adding {
		if slices.Contains(existing, add) {
			continue
		}
		existing = append(existing, add)
	}
	return existing
}

func convert23Relationship(idMap map[string]any, rel *v2_3.Relationship) *Relationship {
	if rel == nil {
		return nil
	}
	from, _ := idMap[string(rel.RefA.ElementRefID)].(IElement)
	to, _ := idMap[string(rel.RefB.ElementRefID)].(IElement)
	if from == nil || to == nil {
		return nil
	}
	return &Relationship{
		Comment:          rel.RelationshipComment,
		From:             from,
		RelationshipType: convert23RelationshipType(rel.Relationship),
		To:               []IElement{to},
	}
}

func convert23RelationshipType(relationship string) RelationshipType {
	switch strings.ToLower(relationship) {
	case "contains", "contained_by":
		return RelationshipType_Contains
	case "depends_on", "dependency_of":
		return RelationshipType_DependsOn
	}
	return RelationshipType{}
}

func convert23CreationInfo(info *v2_3.CreationInfo) *CreationInfo {
	return &CreationInfo{
		Comment:     info.CreatorComment,
		Created:     info.Created,
		CreatedBy:   convert23Creators(info.Creators),
		SpecVersion: info.LicenseListVersion,
	}
}

func convert23Creators(creators []common.Creator) []IAgent {
	var out []IAgent
	for _, c := range creators {
		out = append(out, convert23Agent(c.CreatorType, c.Creator))
	}
	return out
}

func convert23Agent(agentType string, agent string) IAgent {
	switch strings.ToLower(agentType) {
	case "person":
		return &Person{
			Name: agent,
		}
	case "organization", "org":
		return &Organization{
			Name: agent,
		}
	}
	return nil
}

func convert23File(idMap map[string]any, file *v2_3.File) IFile {
	if file == nil {
		return nil
	}
	return &File{
		SpdxId:          string(file.FileSPDXIdentifier),
		Comment:         file.FileComment,
		Name:            file.FileName,
		AttributionText: file.FileAttributionTexts,
		CopyrightText:   file.FileCopyrightText,
		FileKind:        FileKindType_File,
	}
}

func convert23Package(idMap map[string]any, pkg *v2_3.Package) *Package {
	if pkg == nil {
		return nil
	}
	return &Package{
		SpdxId:              string(pkg.PackageSPDXIdentifier),
		Comment:             pkg.PackageComment,
		Description:         pkg.PackageDescription,
		ExternalIdentifiers: convert23ExternalIdentifiers(pkg.PackageExternalReferences),
		Name:                pkg.PackageName,
		Summary:             pkg.PackageSummary,
		VerifiedUsing:       convert23VerifiedUsing(pkg.PackageVerificationCode),
		BuiltTime:           pkg.BuiltDate,
		OriginatedBy:        convert23PackageOriginator(pkg.PackageOriginator),
		ReleaseTime:         pkg.ReleaseDate,
		SuppliedBy:          convert23Supplier(pkg.PackageSupplier),
		ValidUntilTime:      pkg.ValidUntilDate,
		CopyrightText:       pkg.PackageCopyrightText,
		PrimaryPurpose:      convert23PrimaryPurpose(pkg.PrimaryPackagePurpose),
		DownloadLocation:    pkg.PackageDownloadLocation,
		HomePage:            pkg.PackageHomePage,
		PackageUrl:          convert23PackageUrl(pkg.PackageExternalReferences),
		PackageVersion:      pkg.PackageVersion,
		SourceInfo:          pkg.PackageSourceInfo,
	}
}

func convert23PackageUrl(references []*v2_3.PackageExternalReference) string {
	for _, ref := range references {
		if ref.RefType == common.TypePackageManagerPURL {
			return ref.Locator
		}
	}
	return ""
}

func convert23PrimaryPurpose(purpose string) SoftwarePurpose {
	switch purpose {
	case "container":
		return SoftwarePurpose_Container
	case "library":
		return SoftwarePurpose_Library
	case "application":
		return SoftwarePurpose_Application
	}
	return SoftwarePurpose{}
}

func convert23Supplier(supplier *common.Supplier) IAgent {
	if supplier == nil {
		return nil
	}
	return convert23Agent(supplier.SupplierType, supplier.Supplier)
}

func convert23PackageOriginator(originator *common.Originator) []IAgent {
	if originator == nil {
		return nil
	}
	return []IAgent{convert23Agent(originator.OriginatorType, originator.Originator)}
}

func convert23VerifiedUsing(verificationCode *common.PackageVerificationCode) []IIntegrityMethod {
	return nil // TODO
}

func convert23ExternalIdentifiers(references []*v2_3.PackageExternalReference) []IExternalIdentifier {
	var out []IExternalIdentifier
	for _, r := range references {
		typ := ExternalIdentifierType{}
		switch r.RefType {
		case common.TypeSecurityCPE22Type:
			typ = ExternalIdentifierType_Cpe22
		case common.TypeSecurityCPE23Type:
			typ = ExternalIdentifierType_Cpe23
		case common.TypePackageManagerPURL:
			typ = ExternalIdentifierType_PackageUrl
		default:
			continue // unknown
		}
		out = append(out, &ExternalIdentifier{
			Comment:                r.ExternalRefComment,
			ExternalIdentifierType: typ,
			Identifier:             r.Locator,
		})
	}
	return out
}

var _ converter.ConvertFrom = (*SpdxDocument)(nil)
