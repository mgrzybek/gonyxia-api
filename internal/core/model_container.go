package core

type Container struct {
	Name string `json:"name,omitempty"`

	Ready bool `json:"ready,omitempty"`
}
