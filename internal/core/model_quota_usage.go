package core

type QuotaUsage struct {
	Spec *Quota `json:"spec,omitempty"`

	Usage *Quota `json:"usage,omitempty"`
}
