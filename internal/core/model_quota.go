package core

// Quota refers to the limits to apply against a namespace
type Quota struct {
	RequestsMemory string `json:"requests.memory,omitempty"`

	RequestsCPU string `json:"requests.cpu,omitempty"`

	LimitsMemory string `json:"limits.memory,omitempty"`

	LimitsCPU string `json:"limits.cpu,omitempty"`

	RequestsStorage string `json:"requests.storage,omitempty"`

	CountPods int64 `json:"count/pods,omitempty"`
}
