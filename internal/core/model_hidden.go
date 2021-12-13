package core

type Hidden struct {
	Value *interface{} `json:"value,omitempty"`

	Path string `json:"path,omitempty"`
}
