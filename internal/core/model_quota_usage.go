package core

type quotaUsage struct {
	Spec *Quota `json:"spec,omitempty"`

	Usage *Quota `json:"usage,omitempty"`
}
