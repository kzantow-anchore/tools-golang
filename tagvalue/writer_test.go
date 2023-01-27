package tagvalue

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/spdx/tools-golang/spdx"
	"github.com/spdx/tools-golang/spdx/v2/v2_3/tagvalue/writer"
)

func Test_Writer(t *testing.T) {
	tests := []struct {
		name     string
		source   spdx.Document
		expected string
	}{
		{
			name: "full document",
			source: spdx.Document{
				SPDXVersion:       spdx.Version,
				DataLicense:       spdx.DataLicense,
				SPDXIdentifier:    "DOCUMENT",
				DocumentName:      "doc-name",
				DocumentNamespace: "doc-namespace",
				ExternalDocumentReferences: []spdx.ExternalDocumentRef{
					{
						DocumentRefID: "example-1",
						URI:           "https://raw.githubusercontent.com/spdx/spdx-examples/master/example1/spdx/example1.spdx",
						Checksum: spdx.Checksum{
							Algorithm: spdx.SHA1,
							Value:     "b92c6fb161b39991d96613f7b8c348422cf53c58",
						},
					},
					{
						DocumentRefID: "example-maven",
						URI:           "https://raw.githubusercontent.com/spdx/spdx-examples/master/example8/spdx/examplemaven-0.0.1.spdx.json",
						Checksum: spdx.Checksum{
							Algorithm: spdx.SHA256,
							Value:     "d784b605183745c00ee47b8acafcb0b11c2f21dcad7662254e5247ad894eeca7",
						},
					},
					{
						DocumentRefID: "spdx-tool-1.2",
						URI:           "https://spdx.org/spdxdocs/spdx-tools-v1.2-3F2504E0-4F89-41D3-9A0C-0305E82C3301",
						Checksum: spdx.Checksum{
							Algorithm: spdx.SHA1,
							Value:     "d6a770ba38583ed4bb4525bd96e50461655d2759",
						},
					},
				},
				DocumentComment: "some doc comment with \n multiple lines",
				CreationInfo: &spdx.CreationInfo{
					LicenseListVersion: "license list version",
					Creators: []spdx.Creator{
						{
							Creator:     "Some Org, Inc.",
							CreatorType: "Organization",
						},
						{
							Creator:     "John Doe",
							CreatorType: "Person",
						},
					},
					Created:        "2010-01-29T18:30:22Z",
					CreatorComment: "Some creator comment",
				},
				Packages: []*spdx.Package{
					{
						IsUnpackaged:              true,
						PackageName:               "package name 1",
						PackageSPDXIdentifier:     "id-1",
						PackageVersion:            "version-1",
						PackageFileName:           "file 1",
						PackageSupplier:           nil,
						PackageOriginator:         nil,
						PackageDownloadLocation:   "",
						FilesAnalyzed:             true,
						IsFilesAnalyzedTagPresent: true,
						PackageVerificationCode: &spdx.PackageVerificationCode{
							Value:         "d6a770ba38583ed4bb4525bd96e50461655d2758",
							ExcludedFiles: []string{"a", "b"},
						},
						PackageChecksums: []spdx.Checksum{
							{
								Algorithm: spdx.SHA1,
								Value:     "d6a770ba38583ed4bb4525bd96e50461655d2758",
							},
							{
								Algorithm: spdx.SHA256,
								Value:     "d784b605183745c00ee47b8acafcb0b11c2f21dcad7662254e5247ad894eeca7",
							},
						},
						PackageHomePage:             "https://spdx.org/spdxdocs/example-home-page",
						PackageSourceInfo:           "source info 1",
						PackageLicenseConcluded:     "MIT",
						PackageLicenseInfoFromFiles: []string{"a", "b"},
						PackageLicenseDeclared:      "MIT",
						PackageLicenseComments:      "license comments 1",
						PackageCopyrightText:        "copyright text 1",
						PackageSummary:              "summary 1",
						PackageDescription:          "description 1",
						PackageComment:              "comment 1",
						PackageExternalReferences: []*spdx.PackageExternalReference{
							{
								Category:           "SECURITY",
								RefType:            "cpe23Type",
								Locator:            "cpe:2.3:a:pivotal_software:spring_framework:4.1.0:*:*:*:*:*:*:*",
								ExternalRefComment: "External ref comment 1",
							},
							{
								Category:           "OTHER",
								RefType:            "swh",
								Locator:            "swh:1:cnt:94a9ed024d3859793618152ea559a168bbcbb5e2",
								ExternalRefComment: "External ref comment 2",
							},
						},
						Files: []*spdx.File{
							{
								FileName:           "file 1",
								FileSPDXIdentifier: "file-id-1",
								FileTypes:          []string{"SOURCE", "APPLICATION"},
								Checksums: []spdx.Checksum{
									{
										Algorithm: spdx.MD5,
										Value:     "f583fb61254335c08df989b315081fb6",
									},
									{
										Algorithm: spdx.SHA256,
										Value:     "23031f795e508b4c51f45a208613469a4140213062316ce2f7dae6b79945877f",
									},
								},
								LicenseConcluded:   "MIT",
								LicenseInfoInFiles: []string{"f1", "f2", "f3"},
								LicenseComments:    "comments 1",
								FileCopyrightText:  "copy text 1",
								ArtifactOfProjects: []*spdx.ArtifactOfProject{
									{
										Name:     "name 1",
										HomePage: "http://some-url.org/page-1",
										URI:      "uri 1",
									},
									{
										Name:     "name 2",
										HomePage: "http://some-other-url.com/",
										URI:      "uri 2",
									},
								},
								FileComment:      "comment 1",
								FileNotice:       "notice 1",
								FileContributors: []string{"c1", "c2"},
								FileDependencies: []string{"dep1", "dep2", "dep3"},
								Snippets: map[spdx.ElementID]*spdx.Snippet{
									spdx.ElementID("e1"): {
										SnippetSPDXIdentifier:         "id1",
										SnippetFromFileSPDXIdentifier: "file1",
										Ranges: []spdx.SnippetRange{
											{
												StartPointer: spdx.SnippetRangePointer{
													Offset:             1,
													LineNumber:         2,
													FileSPDXIdentifier: "f1",
												},
												EndPointer: spdx.SnippetRangePointer{
													Offset:             3,
													LineNumber:         4,
													FileSPDXIdentifier: "f2",
												},
											},
											{
												StartPointer: spdx.SnippetRangePointer{
													Offset:             5,
													LineNumber:         6,
													FileSPDXIdentifier: "f3",
												},
												EndPointer: spdx.SnippetRangePointer{
													Offset:             7,
													LineNumber:         8,
													FileSPDXIdentifier: "f4",
												},
											},
										},
										SnippetLicenseConcluded: "GPL-2.0",
										LicenseInfoInSnippet:    []string{"a", "b"},
										SnippetLicenseComments:  "license comment 1",
										SnippetCopyrightText:    "copy 1",
										SnippetComment:          "comment 1",
										SnippetName:             "name 1",
									},
									spdx.ElementID("e2"): {
										SnippetSPDXIdentifier:         "id2",
										SnippetFromFileSPDXIdentifier: "file2",
										Ranges: []spdx.SnippetRange{
											{
												StartPointer: spdx.SnippetRangePointer{
													Offset:             5,
													LineNumber:         6,
													FileSPDXIdentifier: "f3",
												},
												EndPointer: spdx.SnippetRangePointer{
													Offset:             7,
													LineNumber:         8,
													FileSPDXIdentifier: "f4",
												},
											},
											{
												StartPointer: spdx.SnippetRangePointer{
													Offset:             9,
													LineNumber:         10,
													FileSPDXIdentifier: "f13",
												},
												EndPointer: spdx.SnippetRangePointer{
													Offset:             11,
													LineNumber:         12,
													FileSPDXIdentifier: "f14",
												},
											},
										},
										SnippetLicenseConcluded: "MIT",
										LicenseInfoInSnippet:    []string{"a", "b"},
										SnippetLicenseComments:  "license comment 1",
										SnippetCopyrightText:    "copy 1",
										SnippetComment:          "comment 1",
										SnippetName:             "name 1",
									},
								},
								//Annotations: []spdx.Annotation{
								//	{
								//		Annotator: spdx.Annotator{
								//			Annotator:     "ann 1",
								//			AnnotatorType: spdx.AnnotatorTypePerson,
								//		},
								//		AnnotationDate: "2020-01-29T18:30:22Z",
								//		AnnotationType: spdx.AnnotationTypeOther,
								//		AnnotationSPDXIdentifier: spdx.DocElementID{
								//			DocumentRefID: "doc-ref-1",
								//			ElementRefID:  "elem-id-1",
								//		},
								//		AnnotationComment: "comment 1",
								//	},
								//	{
								//		Annotator: spdx.Annotator{
								//			Annotator:     "ann 2",
								//			AnnotatorType: spdx.AnnotatorTypeOrganization,
								//		},
								//		AnnotationDate: "2022-01-29T18:30:22Z",
								//		AnnotationType: spdx.AnnotationTypeReview,
								//		AnnotationSPDXIdentifier: spdx.DocElementID{
								//			DocumentRefID: "doc-ref-2",
								//			ElementRefID:  "elem-id-2",
								//			SpecialID:     "spec 2",
								//		},
								//		AnnotationComment: "comment 2",
								//	},
								//},
							},
						},
						Annotations: []spdx.Annotation{
							{
								Annotator: spdx.Annotator{
									Annotator:     "ann 1",
									AnnotatorType: spdx.AnnotatorTypeTool,
								},
								AnnotationDate: "2020-01-29T18:30:22Z",
								AnnotationType: spdx.AnnotationTypeReview,
								AnnotationSPDXIdentifier: spdx.DocElementID{
									DocumentRefID: "doc-ref-1",
									ElementRefID:  "elem-id-1",
								},
								AnnotationComment: "comment 1",
							},
							{
								Annotator: spdx.Annotator{
									Annotator:     "ann 2",
									AnnotatorType: spdx.AnnotatorTypePerson,
								},
								AnnotationDate: "2022-01-29T18:30:22Z",
								AnnotationType: spdx.AnnotationTypeReview,
								AnnotationSPDXIdentifier: spdx.DocElementID{
									DocumentRefID: "doc-ref-2",
									ElementRefID:  "elem-id-2",
									SpecialID:     "spec 2",
								},
								AnnotationComment: "comment 2",
							},
						},
					},
				},
				Files: []*spdx.File{
					{
						FileName:           "file 1",
						FileSPDXIdentifier: "file-id-2",
						FileTypes:          []string{"SOURCE", "APPLICATION"},
						Checksums: []spdx.Checksum{
							{
								Algorithm: spdx.MD5,
								Value:     "f583fb61254335c08df989b315081fb6",
							},
							{
								Algorithm: spdx.SHA256,
								Value:     "23031f795e508b4c51f45a208613469a4140213062316ce2f7dae6b79945877f",
							},
						},
						LicenseConcluded:   "MIT AND GPL-2.0",
						LicenseInfoInFiles: []string{"f1", "f2", "f3"},
						LicenseComments:    "comments 1",
						FileCopyrightText:  "copy 1",
						ArtifactOfProjects: []*spdx.ArtifactOfProject{
							{
								Name:     "name 1",
								HomePage: "http://some-url.org/page-1",
								URI:      "http://some-url.org/page-1/uri",
							},
							{
								Name:     "name 2",
								HomePage: "http://some-other-url.com/",
								URI:      "http://some-other-url.com/uri",
							},
						},
						FileComment:      "comment 1",
						FileNotice:       "notice 1",
						FileContributors: []string{"c1", "c2"},
						FileDependencies: []string{"d1", "d2", "d3", "d4"},
						Snippets:         nil, // already have snippets elsewhere
						Annotations:      nil, // already have annotations elsewhere
					},
					{
						FileName:           "file 2",
						FileSPDXIdentifier: "id 2",
						FileTypes:          []string{"SOURCE", "TEXT"},
						Checksums: []spdx.Checksum{
							{
								Algorithm: spdx.MD5,
								Value:     "3b3d84497fe614195e799aeafdd0740b",
							},
							{
								Algorithm: spdx.SHA1,
								Value:     "a518255d3dcf2c6f3fa8a4de15b6aafc77221b29",
							},
						},
						LicenseConcluded:   "Apache-2.0",
						LicenseInfoInFiles: []string{"f1", "f2", "f3"},
						LicenseComments:    "comments 2",
						FileCopyrightText:  "copy 2",
						ArtifactOfProjects: []*spdx.ArtifactOfProject{
							{
								Name:     "name 2",
								HomePage: "http://some-other-url.com/",
								URI:      "http://some-url.org/page-1/uri",
							},
							{
								Name:     "name 4",
								HomePage: "http://some-fourth.edu/",
								URI:      "http://some-fourth.edu/uri",
							},
						},
						FileComment:      "comment 2",
						FileNotice:       "notice 2",
						FileContributors: []string{"c1", "c2"},
						FileDependencies: []string{"d1", "d2", "d3", "d4"},
						Snippets:         nil, // already have snippets elsewhere
						Annotations:      nil, // already have annotations elsewhere
					},
				},
				OtherLicenses: []*spdx.OtherLicense{
					{
						LicenseIdentifier:      "LGPL-3.0",
						ExtractedText:          "text 1",
						LicenseName:            "name 1",
						LicenseCrossReferences: []string{"x1", "x2", "x3"},
						LicenseComment:         "comment 1",
					},
					{
						LicenseIdentifier:      "Apache-2.0",
						ExtractedText:          "text 2",
						LicenseName:            "name 2",
						LicenseCrossReferences: []string{"x4", "x5", "x6"},
						LicenseComment:         "comment 2",
					},
				},
				Relationships: []*spdx.Relationship{
					{
						RefA: spdx.DocElementID{
							DocumentRefID: "doc-ref-1",
							ElementRefID:  "elem-id-1",
						},
						RefB: spdx.DocElementID{
							DocumentRefID: "doc-ref-2",
							ElementRefID:  "elem-id-2",
						},
						Relationship:        spdx.RelationshipContains,
						RelationshipComment: "comment 1",
					},
					{
						RefA: spdx.DocElementID{
							DocumentRefID: "doc-ref-3",
							ElementRefID:  "elem-id-3",
						},
						RefB: spdx.DocElementID{
							DocumentRefID: "doc-ref-4",
							ElementRefID:  "elem-id-4",
							SpecialID:     "special-id-4",
						},
						Relationship:        spdx.RelationshipCopyOf,
						RelationshipComment: "comment 2",
					},
				},
				Annotations: []*spdx.Annotation{
					{
						Annotator: spdx.Annotator{
							Annotator:     "annotator 1",
							AnnotatorType: spdx.AnnotatorTypePerson,
						},
						AnnotationDate: "2020-01-29T18:30:22Z",
						AnnotationType: spdx.AnnotationTypeReview,
						AnnotationSPDXIdentifier: spdx.DocElementID{
							DocumentRefID: "doc-ref-1",
							ElementRefID:  "elem-id-1",
							SpecialID:     "spec 1",
						},
						AnnotationComment: "comment 1",
					},
					{
						Annotator: spdx.Annotator{
							Annotator:     "annotator 2",
							AnnotatorType: spdx.AnnotatorTypeOrganization,
						},
						AnnotationDate: "2022-01-29T18:30:22Z",
						AnnotationType: spdx.AnnotationTypeOther,
						AnnotationSPDXIdentifier: spdx.DocElementID{
							DocumentRefID: "doc-ref-2",
							ElementRefID:  "elem-id-2",
							SpecialID:     "spec 2",
						},
						AnnotationComment: "comment 2",
					},
				},
				Snippets: []spdx.Snippet{
					{
						SnippetSPDXIdentifier:         "id1",
						SnippetFromFileSPDXIdentifier: "file1",
						Ranges: []spdx.SnippetRange{
							{
								StartPointer: spdx.SnippetRangePointer{
									Offset:             1,
									LineNumber:         2,
									FileSPDXIdentifier: "f1",
								},
								EndPointer: spdx.SnippetRangePointer{
									Offset:             3,
									LineNumber:         4,
									FileSPDXIdentifier: "f2",
								},
							},
							{
								StartPointer: spdx.SnippetRangePointer{
									Offset:             5,
									LineNumber:         6,
									FileSPDXIdentifier: "f3",
								},
								EndPointer: spdx.SnippetRangePointer{
									Offset:             7,
									LineNumber:         8,
									FileSPDXIdentifier: "f4",
								},
							},
						},
						SnippetLicenseConcluded: "MIT",
						LicenseInfoInSnippet:    []string{"a", "b"},
						SnippetLicenseComments:  "license comment 1",
						SnippetCopyrightText:    "copy 1",
						SnippetComment:          "comment 1",
						SnippetName:             "name 1",
					},
					{
						SnippetSPDXIdentifier:         "id2",
						SnippetFromFileSPDXIdentifier: "file2",
						Ranges: []spdx.SnippetRange{
							{
								StartPointer: spdx.SnippetRangePointer{
									Offset:             5,
									LineNumber:         6,
									FileSPDXIdentifier: "f3",
								},
								EndPointer: spdx.SnippetRangePointer{
									Offset:             7,
									LineNumber:         8,
									FileSPDXIdentifier: "f4",
								},
							},
							{
								StartPointer: spdx.SnippetRangePointer{
									Offset:             9,
									LineNumber:         10,
									FileSPDXIdentifier: "f13",
								},
								EndPointer: spdx.SnippetRangePointer{
									Offset:             11,
									LineNumber:         12,
									FileSPDXIdentifier: "f14",
								},
							},
						},
						SnippetLicenseConcluded: "GPL-2.0",
						LicenseInfoInSnippet:    []string{"a", "b"},
						SnippetLicenseComments:  "license comment 1",
						SnippetCopyrightText:    "copy 1",
						SnippetComment:          "comment 1",
						SnippetName:             "name 1",
					},
				},
				Reviews: []*spdx.Review{
					{
						Reviewer:      "reviewer 1",
						ReviewerType:  spdx.AnnotatorTypePerson,
						ReviewDate:    "2020-01-29T18:30:22Z",
						ReviewComment: "comment 1",
					},
					{
						Reviewer:      "reviewer 2",
						ReviewerType:  spdx.AnnotatorTypeTool,
						ReviewDate:    "2022-01-29T18:30:22Z",
						ReviewComment: "comment 2",
					},
				},
			},
			expected: `
`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := writer.RenderDocument(&test.source, buf)
			require.NoError(t, err)

			b := buf.Bytes()

			//b = regexp.MustCompile("ExternalDocumentRef:.*\n").ReplaceAll(b, []byte(""))
			err = ioutil.WriteFile("sample.orig.spdx", b, fs.ModePerm)

			buf = new(bytes.Buffer)
			err = Write(test.source, buf)
			require.NoError(t, err)

			b = buf.Bytes()

			//b = regexp.MustCompile("ExternalDocumentRef:.*\n").ReplaceAll(b, []byte(""))
			err = ioutil.WriteFile("sample.spdx", b, fs.ModePerm)

			require.NoError(t, err)

			require.Equal(t, test.expected, string(b))
		})
	}
}
