package core

type Services struct {
	Server *Server `json:"server"`

	Driver OrchestratorAdapter

	Type_ string `json:"type,omitempty"`

	SingleNamespace bool `json:"singleNamespace,omitempty"`

	NamespacePrefix string `json:"namespacePrefix,omitempty"`

	GroupNamespacePrefix string `json:"groupNamespacePrefix,omitempty"`

	UsernamePrefix string `json:"usernamePrefix,omitempty"`

	GroupPrefix string `json:"groupPrefix,omitempty"`

	AuthenticationMode string `json:"authenticationMode,omitempty"`

	Expose *Expose `json:"expose,omitempty"`

	Monitoring *Monitoring `json:"monitoring,omitempty"`

	Cloudshell *CloudshellConfiguration `json:"cloudshell,omitempty"`

	InitScript string `json:"initScript,omitempty"`

	Quotas *Quotas `json:"quotas,omitempty"`

	DefaultConfiguration *DefaultConfiguration `json:"defaultConfiguration,omitempty"`
}
