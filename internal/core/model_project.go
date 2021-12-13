package core

type project struct {
	ID string `json:"id,omitempty"`

	Group string `json:"group,omitempty"`

	Bucket string `json:"bucket,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	Name string `json:"name,omitempty"`
}
