package core

type ServicesListing struct {
	Apps []Service `json:"apps,omitempty"`

	Groups []Group `json:"groups,omitempty"`
}
