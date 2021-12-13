package rest

import (
	"net/http"
)

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
	w.WriteHeader(http.StatusNotImplemented)
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

func resetQuota1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func resetQuota2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func resetQuota3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func resetQuota4(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func resetQuota5(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}

func resetQuota6(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if !validateAuthorizationHeader(w, r) {
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}
