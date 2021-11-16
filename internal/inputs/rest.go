package inputs

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.con/mgrzybek/gonyxia-api/internal/core"
	"github.con/mgrzybek/gonyxia-api/internal/inputs/rest"
)

type RestService struct {
	Bind_addr    string
	Publish_addr string
	Engine       core.Engine
}

func (r *RestService) Run() {
	log.Info("Server started on " + r.Bind_addr)
	router := rest.NewRouter()
	log.Fatal(http.ListenAndServe(r.Bind_addr, router))
}
