package core

type pkg struct {
	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Version string `json:"version,omitempty"`

	Config *config `json:"config,omitempty"`
}
