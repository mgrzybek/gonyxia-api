package core

type container struct {
	Name string `json:"name,omitempty"`

	Ready bool `json:"ready,omitempty"`
}
