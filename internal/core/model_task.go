package core

type task struct {
	ID string `json:"id,omitempty"`

	Status *taskStatus `json:"status,omitempty"`

	Containers []container `json:"containers,omitempty"`
}
