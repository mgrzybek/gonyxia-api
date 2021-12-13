package core

type group struct {
	ID string `json:"id,omitempty"`

	Apps []service `json:"apps,omitempty"`
}
