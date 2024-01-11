package main

type DescriptionSpec struct {
	TableDescs   []TableDesc `json:"tablesDescs"`
	OutputFormat Format      `json:"outputFormat"`
	ForPreview   bool        `json:"forPreview"`
}
