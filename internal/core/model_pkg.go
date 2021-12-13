package core

type Pkg struct {
	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Version string `json:"version,omitempty"`

	Config *Config `json:"config,omitempty"`
}
