package core

type S3 struct {
	Monitoring *Monitoring `json:"monitoring,omitempty"`

	URL string `json:"URL,omitempty"`
}
