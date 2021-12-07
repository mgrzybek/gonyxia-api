/*
 * Onyxia-api
 *
 * Swagger onyxia-api
 *
 * API version: snapshot
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package rest

import (
//	"encoding/json"
//	"fmt"
	"net/http"
	"strings"

//	log "github.com/sirupsen/logrus"
)

func Configuration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func GetCatalogById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func GetCatalogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func GetIP(w http.ResponseWriter, r *http.Request) {
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
	writeHttpResponse(w, http.StatusOK, message)
}

func GetPackage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func GetRegions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
/*
	result, err := json.Marshal(engine.Regions)

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "%s", result)
*/
	w.WriteHeader(http.StatusOK)
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}
