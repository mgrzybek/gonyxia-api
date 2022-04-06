package core

// OrchestratorAdapter is the interface that represents an orchestrator used
// to host users’ workloads.
type OrchestratorAdapter interface {
	Health() error

	UserCreate() error
	UserDelete() error

	NamespaceCreate(name string, quota Quota, owner string) error
	NamespaceDelete(name string) error

	RoleCreate() error
	RoleDelete() error

	GetQuota(namespaceID string) (Quota, error)
	SetQuota(quota Quota, namespaceID string) error
}

// IdentityManagerAdaptor is the interface that represents the backoffice used to
// manage IDs (users, passwords, token validation…)
type IdentityManagerAdaptor interface {
	GetUserInfoFromBearerToken(token string) (OnyxiaUser, error)
}
