package core

type Task struct {
	Id string `json:"id,omitempty"`

	Status *TaskStatus `json:"status,omitempty"`

	Containers []Container `json:"containers,omitempty"`
}
