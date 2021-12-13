package core

type Config struct {
	Type_ string `json:"type,omitempty"`

	Properties *Category `json:"properties,omitempty"`

	Required []string `json:"required,omitempty"`
}
