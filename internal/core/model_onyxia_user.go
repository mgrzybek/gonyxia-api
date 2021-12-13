package core

type OnyxiaUser struct {
	Email string `json:"email,omitempty"`

	Idep string `json:"idep,omitempty"`

	NomComplet string `json:"nomComplet,omitempty"`

	Password string `json:"password,omitempty"`

	Ip string `json:"ip,omitempty"`

	Groups []string `json:"groups,omitempty"`

	Projects []Project `json:"projects,omitempty"`
}
