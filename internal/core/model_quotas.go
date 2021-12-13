package core

type Quotas struct {
	Enabled bool `json:"enabled,omitempty"`

	AllowUserModification bool `json:"allowUserModification,omitempty"`

	Default_ *Quota `json:"default,omitempty"`
}
