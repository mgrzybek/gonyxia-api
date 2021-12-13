package core

type event struct {
	Message string `json:"message,omitempty"`

	Timestamp int64 `json:"timestamp,omitempty"`
}
