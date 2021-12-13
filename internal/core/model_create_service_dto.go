package core

type createServiceDto struct {
	CatalogID string `json:"catalogId,omitempty"`

	PackageName string `json:"packageName,omitempty"`

	PackageVersion string `json:"packageVersion,omitempty"`

	Name string `json:"name,omitempty"`

	Options *interface{} `json:"options,omitempty"`

	DryRun bool `json:"dryRun,omitempty"`
}
