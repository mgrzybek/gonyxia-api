package core

type config struct {
	Type string `json:"type,omitempty"`

	Properties *category `json:"properties,omitempty"`

	Required []string `json:"required,omitempty"`
}
