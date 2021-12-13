package core

type Catalog struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Maintainer string `json:"maintainer,omitempty"`

	Location string `json:"location,omitempty"`

	Status string `json:"status,omitempty"`

	LastUpdateTime int64 `json:"lastUpdateTime,omitempty"`

	Scm string `json:"scm,omitempty"`

	Type_ string `json:"type,omitempty"`
}
