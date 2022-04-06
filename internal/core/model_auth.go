package core

// Auth is used to interact with an identity provider
type Auth struct {
	AuthType            string `json:"type"`
	Realm               string `json:"string:realm"`
	Resource            string `json:"string:resource"`
	AuthServerURL       string `json:"string:auth-server-url"`
	RedirectURL         string `json:"string:redirect-server-url"`
	ClientID            string `json:"string:client-id"`
	ClientSecret        string `json:"string:client-secret"`
	SSLRequired         string `json:"string:ssl-required"`
	PublicClient        bool   `json:"bool:public-client"`
	EnableBasicAuth     bool   `json:"bool:enable-basic-auth"`
	BearerOnly          bool   `json:"bool:bearer-only"`
	DisableTrustManager bool   `json:"bool:disable-trust-manager"`

	Driver IdentityManagerAdaptor
}
