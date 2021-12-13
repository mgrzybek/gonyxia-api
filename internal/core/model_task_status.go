package core

type TaskStatus struct {
	Status string `json:"status,omitempty"`

	Reason string `json:"reason,omitempty"`
}
