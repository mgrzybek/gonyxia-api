package core

type AppInfo struct {
	Build *BuildInfo `json:"build,omitempty"`

	Regions []Region `json:"regions,omitempty"`
}
