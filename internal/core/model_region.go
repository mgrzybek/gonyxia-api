package core

type Region struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Location *Location `json:"location,omitempty"`

	Services *Services `json:"services,omitempty"`

	OnyxiaAPI *OnyxiaAPI `json:"onyxiaAPI,omitempty"`

	Data *Data `json:"data,omitempty"`
}
