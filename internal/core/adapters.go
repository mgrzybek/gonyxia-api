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
}
