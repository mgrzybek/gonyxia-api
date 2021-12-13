package core

type hidden struct {
	Value *interface{} `json:"value,omitempty"`

	Path string `json:"path,omitempty"`
}
