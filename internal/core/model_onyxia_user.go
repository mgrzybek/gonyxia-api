package core

type onyxiaUser struct {
	Email string `json:"email,omitempty"`

	Idep string `json:"idep,omitempty"`

	NomComplet string `json:"nomComplet,omitempty"`

	Password string `json:"password,omitempty"`

	IP string `json:"ip,omitempty"`

	Groups []string `json:"groups,omitempty"`

	Projects []project `json:"projects,omitempty"`
}
