package core

type Region struct {
	id, name, description string
	location              Location
	services              Services
	onyxiaAPI             OnyxiaAPI
	data                  Data
}

type Service struct {
}

type Services struct {
	stype                       Service.ServiceType
	singleNamespace             bool   `default:true`
	namespacePrefix             string `default:"user-"`
	groupNamespacePrefix        string `default:"projet-"`
	usernamePrefix, groupPrefix string
	authenticationMode          AuthenticationMode `default:AuthenticationMode.IMPERSONATE`
	expose                      Expose
	server                      Server
	monitoring                  Monitoring
	cloudshell                  CloudshellConfiguration
	initScript                  string
	quotas                      Quotas
	defaultConfiguration        DefaultConfiguration
}

type DefaultConfiguration struct {
	ipProtection, networkPolicy bool `default:true`
}

type Quota struct {
	enabled               bool `default:false`
	allowUserModification bool `default:true`
	defaultQuota          Quota
}

type Monitoring struct {
	urlPattern string
}

type Data struct {
	s3 S3
}

type Expose struct {
	domain string
}

type Server struct {
	url  string
	auth Auth
}

type OnyxiaAPI struct {
	baseURL string
}

type Location struct {
	latitude, longitude float64
	name                string
}

type CloudshellConfiguration struct {
	catalogID, packageName string
}

type Auth struct {
	token, username, password string
}
