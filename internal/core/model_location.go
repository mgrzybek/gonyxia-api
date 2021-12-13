package core

type location struct {
	Lat float64 `json:"lat,omitempty"`

	Name string `json:"name,omitempty"`

	Long float64 `json:"long,omitempty"`
}
