package rest

import (
	"net/http"
)

func userInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}
