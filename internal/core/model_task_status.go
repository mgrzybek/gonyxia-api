package core

type taskStatus struct {
	Status string `json:"status,omitempty"`

	Reason string `json:"reason,omitempty"`
}
