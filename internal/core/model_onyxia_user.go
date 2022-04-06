package core

// OnyxiaUser represents a user account and their groups/projects
type OnyxiaUser struct {
	Email      string    `json:"email,omitempty"`
	Idep       string    `json:"idep,omitempty"`
	NomComplet string    `json:"nomComplet,omitempty"`
	Password   string    `json:"password,omitempty"`
	IP         string    `json:"ip,omitempty"`
	Groups     []string  `json:"groups,omitempty"`
	Projects   []project `json:"projects,omitempty"`
}
