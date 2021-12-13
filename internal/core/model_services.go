package core

// Services represents an orchestrator endpoint (driver,
// default configurations…).
type Services struct {
	Server *server `json:"server"`

	Driver OrchestratorAdapter

	Type string `json:"type,omitempty"`

	SingleNamespace bool `json:"singleNamespace,omitempty"`

	NamespacePrefix string `json:"namespacePrefix,omitempty"`

	GroupNamespacePrefix string `json:"groupNamespacePrefix,omitempty"`

	UsernamePrefix string `json:"usernamePrefix,omitempty"`

	GroupPrefix string `json:"groupPrefix,omitempty"`

	AuthenticationMode string `json:"authenticationMode,omitempty"`

	Expose *expose `json:"expose,omitempty"`

	Monitoring *monitoring `json:"monitoring,omitempty"`

	Cloudshell *cloudshellConfiguration `json:"cloudshell,omitempty"`

	InitScript string `json:"initScript,omitempty"`

	Quotas *quotas `json:"quotas,omitempty"`

	DefaultConfiguration *defaultConfiguration `json:"defaultConfiguration,omitempty"`
}
