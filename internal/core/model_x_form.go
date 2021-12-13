package core

type XForm struct {
	Hidden bool `json:"hidden,omitempty"`

	Readonly bool `json:"readonly,omitempty"`

	Value string `json:"value,omitempty"`
}
