package inputs

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/mgrzybek/gonyxia-api/internal/core"
	"github.com/mgrzybek/gonyxia-api/internal/inputs/rest"
)

type RestService struct {
	bind_addr    string
	publish_addr string
	engine       *core.Engine
}

func NewRestService(b,p string, e *core.Engine) RestService {
	return RestService{
		bind_addr: b,
		publish_addr: p,
		engine: e,
	}
}

func (r *RestService) Run() {
	log.Info("Server started on " + r.bind_addr)
	router := rest.NewRouter(r.engine)
	log.Fatal(http.ListenAndServe(r.bind_addr, router))
}
