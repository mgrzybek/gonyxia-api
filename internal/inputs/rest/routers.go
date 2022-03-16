package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"

	"github.com/mgrzybek/gonyxia-api/internal/core"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

// NewRouter is the webservices Router constructor
func NewRouter(e *core.Engine) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range restRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	engine = e
	return router
}

/*
 * TODO: deal with tracing
 * https://fernando-bandeira.medium.com/building-apis-with-go-part-3-instrumentation-and-error-handling-daba9385e3ec
 */
func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"method":               r.Method,
			"request":              r.RequestURI,
			"resource":             name,
			"process_time_seconds": time.Since(start),
		}).Info("access")
	})
}

func writeHTTPResponseFromString(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	result := make(map[string]string)
	result["message"] = message
	jmsg, _ := writeJSONResponse(result)
	fmt.Fprintf(w, "%s", jmsg)
}

func writeHTTPResponseFromMap(w http.ResponseWriter, status int, message map[string]string) {
	w.WriteHeader(status)
	jmsg, _ := writeJSONResponse(message)
	fmt.Fprintf(w, "%s", jmsg)
}

func validateAuthorizationHeader(w http.ResponseWriter, r *http.Request) bool {
	if len(r.Header.Get("Authorization")) < 7 {
		log.Warn("Given header is too short: size is ", len(r.Header.Get("Authorization")), " expected greater than 7.")
		writeHTTPResponseFromString(w, http.StatusUnauthorized, "Invalid token")
		return false
	}

	log.Debug("Authorization header: " + r.Header.Get("Authorization"))
	matched, err := regexp.MatchString(`^Bearer [A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`, r.Header.Get("Authorization"))

	if err == nil && matched {
		log.Debug("Given token is valid.")
		return true
	}

	log.Warn("Given token is invalid. Header: " + r.Header["Authorization"][0])
	writeHTTPResponseFromString(w, http.StatusUnauthorized, "Given token is invalid.")
	return false
}

func writeJSONResponse(resp map[string]string) ([]byte, error) {
	j, err := json.Marshal(resp)
	if err != nil {
		log.Error("Error happened in JSON marshal. Err:", err)
	}
	return j, err
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var engine *core.Engine
var restRoutes = routes{
	route{
		"Index",
		"GET",
		"/metrics",
		promhttp.Handler().ServeHTTP,
	},

	route{
		"Index",
		"GET",
		"/",
		index,
	},

	route{
		"GetCloudShellStatus",
		strings.ToUpper("Get"),
		"/cloudshell",
		getCloudShellStatus,
	},

	route{
		"ApplyQuota",
		strings.ToUpper("Post"),
		"/my-lab/quota",
		applyQuota,
	},

	route{
		"DestroyApp",
		strings.ToUpper("Delete"),
		"/my-lab/app",
		destroyApp,
	},

	route{
		"GetApp",
		strings.ToUpper("Get"),
		"/my-lab/app",
		getApp,
	},

	route{
		"GetLogs",
		strings.ToUpper("Get"),
		"/my-lab/app/logs",
		getLogs,
	},

	route{
		"GetMyServices",
		strings.ToUpper("Get"),
		"/my-lab/services",
		getMyServices,
	},

	route{
		"GetQuota",
		strings.ToUpper("Get"),
		"/my-lab/quota",
		getQuota,
	},

	route{
		"GetQuotas",
		strings.ToUpper("Get"),
		"/my-lab/quotas/",
		getQuota,
	},

	route{
		"GetQuotas",
		strings.ToUpper("Get"),
		"/my-lab/quotas/{projectID}",
		getQuota,
	},

	route{
		"GetQuota",
		strings.ToUpper("Get"),
		"/my-lab/quota/{regionID}/{projectID}",
		getQuota,
	},

	route{
		"PublishService",
		strings.ToUpper("Put"),
		"/my-lab/app",
		publishService,
	},

	route{
		"ResetQuota",
		strings.ToUpper("POST"),
		"/my-lab/quota/reset",
		resetQuota,
	},

	route{
		"Onboard",
		strings.ToUpper("Post"),
		"/onboarding",
		onboard,
	},

	route{
		"Configuration",
		strings.ToUpper("Get"),
		"/public/configuration",
		configuration,
	},

	route{
		"GetCatalogById",
		strings.ToUpper("Get"),
		"/public/catalog/{catalogId}",
		getCatalogByID,
	},

	route{
		"GetCatalogs",
		strings.ToUpper("Get"),
		"/public/catalog",
		getCatalogs,
	},

	route{
		"GetIP",
		strings.ToUpper("Get"),
		"/public/ip",
		getIP,
	},

	route{
		"GetPackage",
		strings.ToUpper("Get"),
		"/public/catalog/{catalogId}/{packageName}",
		getPackage,
	},

	route{
		"GetRegions",
		strings.ToUpper("Get"),
		"/public/regions",
		getRegions,
	},

	route{
		"Healthcheck",
		strings.ToUpper("Get"),
		"/public/healthcheck",
		healthcheck,
	},

	route{
		"UserInfo",
		strings.ToUpper("Get"),
		"/user/info",
		userInfo,
	},
}
