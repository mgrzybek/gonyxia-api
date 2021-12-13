package core

type quotas struct {
	Enabled bool `json:"enabled,omitempty"`

	AllowUserModification bool `json:"allowUserModification,omitempty"`

	Default *Quota `json:"default,omitempty"`
}
