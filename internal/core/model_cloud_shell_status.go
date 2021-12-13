package core

type cloudShellStatus struct {
	Status string `json:"status,omitempty"`

	URL string `json:"url,omitempty"`

	PackageToDeploy *pkg `json:"packageToDeploy,omitempty"`

	CatalogID string `json:"catalogId,omitempty"`
}
