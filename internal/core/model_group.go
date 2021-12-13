package core

type Group struct {
	Id string `json:"id,omitempty"`

	Apps []Service `json:"apps,omitempty"`
}
