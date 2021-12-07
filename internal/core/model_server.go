package core

type Server struct {
	url string `json:"url"`
	Auth *Credentials `json:"auth"`
}