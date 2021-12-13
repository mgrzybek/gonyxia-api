package core

type Service struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Instances int32 `json:"instances,omitempty"`

	Cpus float64 `json:"cpus,omitempty"`

	Mem float64 `json:"mem,omitempty"`

	Status string `json:"status,omitempty"`

	Type string `json:"type,omitempty"`

	Urls []string `json:"urls,omitempty"`

	InternalUrls []string `json:"internalUrls,omitempty"`

	Logo string `json:"logo,omitempty"`

	Env map[string]string `json:"env,omitempty"`

	Tasks []Task `json:"tasks,omitempty"`

	Events []Event `json:"events,omitempty"`

	Subtitle string `json:"subtitle,omitempty"`

	Monitoring *Monitoring `json:"monitoring,omitempty"`

	PostInstallInstructions string `json:"postInstallInstructions,omitempty"`

	StartedAt int64 `json:"startedAt,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`
}
