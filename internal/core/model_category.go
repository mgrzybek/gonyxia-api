package core

type category struct {
	Properties map[string]property `json:"properties,omitempty"`

	Description string `json:"description,omitempty"`

	Type string `json:"type,omitempty"`

	Required []string `json:"required,omitempty"`
}
