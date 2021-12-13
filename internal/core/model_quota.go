package core

type Quota struct {
	RequestsMemory string `json:"requests.memory,omitempty"`

	RequestsCpu string `json:"requests.cpu,omitempty"`

	LimitsMemory string `json:"limits.memory,omitempty"`

	LimitsCpu string `json:"limits.cpu,omitempty"`

	RequestsStorage string `json:"requests.storage,omitempty"`

	Countpods int32 `json:"count/pods,omitempty"`
}
