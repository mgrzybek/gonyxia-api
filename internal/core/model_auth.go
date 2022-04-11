package core

// Auth is used to interact with an identity provider
type Auth struct {
	AuthType            string `json:"type"`
	Realm               string `json:"realm"`
	Resource            string `json:"resource"`
	AuthServerURL       string `json:"auth-server-url"`
	RedirectURL         string `json:"redirect-server-url"`
	ClientID            string `json:"client-id"`
	ClientSecret        string `json:"client-secret"`
	SSLRequired         string `json:"ssl-required"`
	PublicClient        bool   `json:"public-client"`
	EnableBasicAuth     bool   `json:"enable-basic-auth"`
	BearerOnly          bool   `json:"bearer-only"`
	DisableTrustManager bool   `json:"disable-trust-manager"`

	Driver IdentityManagerAdaptor
}
