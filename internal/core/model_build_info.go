package core

type buildInfo struct {
	Version string `json:"version,omitempty"`

	Timestamp int64 `json:"timestamp,omitempty"`
}
