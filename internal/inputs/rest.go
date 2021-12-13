package inputs

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/mgrzybek/gonyxia-api/internal/core"
	"github.com/mgrzybek/gonyxia-api/internal/inputs/rest"
)

// RestService refers to the webservice object
type RestService struct {
	bindAddr    string
	publishAddr string
	engine      *core.Engine
}

// NewRestService is the RestService’s constructor
func NewRestService(b, p string, e *core.Engine) RestService {
	return RestService{
		bindAddr:    b,
		publishAddr: p,
		engine:      e,
	}
}

// Run starts the REST service on the given socket
func (r *RestService) Run() {
	log.Info("Server started on " + r.bindAddr)
	router := rest.NewRouter(r.engine)
	log.Fatal(http.ListenAndServe(r.bindAddr, router))
}
