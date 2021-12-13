package core

type appInfo struct {
	Build *buildInfo `json:"build,omitempty"`

	Regions []Region `json:"regions,omitempty"`
}
