package core

type Event struct {
	Message string `json:"message,omitempty"`

	Timestamp int64 `json:"timestamp,omitempty"`
}
