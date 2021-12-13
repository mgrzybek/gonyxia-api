package core

type servicesListing struct {
	Apps []service `json:"apps,omitempty"`

	Groups []group `json:"groups,omitempty"`
}
