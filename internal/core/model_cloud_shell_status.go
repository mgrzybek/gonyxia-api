package core

type CloudShellStatus struct {
	Status string `json:"status,omitempty"`

	Url string `json:"url,omitempty"`

	PackageToDeploy *Pkg `json:"packageToDeploy,omitempty"`

	CatalogId string `json:"catalogId,omitempty"`
}
