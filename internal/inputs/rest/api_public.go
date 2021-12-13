package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func configuration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func getCatalogByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := mux.Vars(r)["catalogId"]
	result, err := json.Marshal(engine.GetCatalogByID(id))
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if string(result) != "null" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", result)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func getCatalogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	result, err := json.Marshal(engine.GetCatalogs())
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", result)

	w.WriteHeader(http.StatusOK)
}

func getIP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	message := make(map[string]string)
	ip := r.Header.Get("X-Real-Ip")

	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}

	message["ip"] = strings.Split(ip, ":")[0]
	writeHTTPResponse(w, http.StatusOK, message)
}

func getPackage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func getRegions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	result, err := json.Marshal(engine.GetRegions())

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "%s", result)

	w.WriteHeader(http.StatusOK)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}
