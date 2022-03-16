package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
	"regexp"
)

/*
 * HTTP status codes: https://pkg.go.dev/net/http
 */

func applyQuota(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func destroyApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func getApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func getLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func getMyServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func getQuota(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}

	vars := mux.Vars(r)

	log.Trace("Requested projectID: ", vars["projectID"])
	log.Trace("Requested regionID: ", vars["regionID"])

	regionID, doesKeyExist := vars["regionID"]
	if !doesKeyExist || len(regionID) == 0 {
		// TODO: if regionID is nil, use user’s default project
		writeHTTPResponseFromString(w, http.StatusNotImplemented, "empty regionID is not supported yet.")
		return
	}

	projectID, doesKeyExist := vars["projectID"]
	if !doesKeyExist || len(projectID) == 0 {
		// TODO: if projectID is nil, use user’s default project
		writeHTTPResponseFromString(w, http.StatusNotImplemented, "empty projectID is not supported yet.")
		return
	}

	result, err := engine.GetQuota(projectID, regionID)
	if err != nil {
		isNotFound, _ := regexp.MatchString("not found", err.Error())
		if isNotFound {
			writeHTTPResponseFromString(w, http.StatusNotFound, err.Error())
		}
		writeHTTPResponseFromString(w, http.StatusInternalServerError, err.Error())
	}

	jResult, err := json.Marshal(result)
	if err != nil {
		writeHTTPResponseFromString(w, http.StatusInternalServerError, err.Error())
	}

	fmt.Fprintf(w, "%s", jResult)
}

func getQuotas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}

	vars := mux.Vars(r)

	log.Trace("Requested projectID: ", vars["projectID"])

	id, doesKeyExist := vars["projectID"]
	if !doesKeyExist || len(id) == 0 {
		// TODO: if projectID is nil, use user’s default project
		writeHTTPResponseFromString(w, http.StatusNotImplemented, "empty projectID is not supported yet.")
		return
	}

	result, err := engine.GetQuotas(id)
	if err != nil {
		writeHTTPResponseFromString(w, http.StatusInternalServerError, err.Error())
	}

	log.Trace("result size is:", len(result))
	if len(result) == 0 {
		log.Trace("result is empty, returning HTTP 404.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jResult, err := json.Marshal(result)
	if err != nil {
		writeHTTPResponseFromString(w, http.StatusInternalServerError, err.Error())
	}

	fmt.Fprintf(w, "%s", jResult)
}

func publishService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func resetQuota(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}
