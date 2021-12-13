package core

type Category struct {
	Properties map[string]Property `json:"properties,omitempty"`

	Description string `json:"description,omitempty"`

	Type_ string `json:"type,omitempty"`

	Required []string `json:"required,omitempty"`
}
