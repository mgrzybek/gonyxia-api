package core

// Region refers to a k8s cluster
type Region struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Location *location `json:"location,omitempty"`

	Services *Services `json:"services,omitempty"`

	OnyxiaAPI *onyxiaAPI `json:"onyxiaAPI,omitempty"`

	Data *data `json:"data,omitempty"`
}
