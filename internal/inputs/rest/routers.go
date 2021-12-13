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

type ctxKey struct{}

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

func writeHTTPResponse(w http.ResponseWriter, status int, message map[string]string) {
	w.WriteHeader(status)
	jmsg, _ := writeJSONResponse(message)
	fmt.Fprintf(w, "%s", jmsg)
}

func validateAuthorizationHeader(w http.ResponseWriter, r *http.Request) bool {
	message := make(map[string]string)
	message["message"] = "Given token is invalid."

	if len(r.Header.Get("Authorization")) < 7 {
		log.Warn(message["message"])
		writeHTTPResponse(w, http.StatusUnauthorized, message)
		return false
	}

	log.Debug("Authorization header: " + r.Header.Get("Authorization"))
	matched, err := regexp.MatchString(`^Bearer [A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`, r.Header.Get("Authorization"))

	if err == nil && matched == true {
		log.Debug("Given token is valid")
		return true
	}

	log.Warn(message["message"] + " Header: " + r.Header["Authorization"][0])
	writeHTTPResponse(w, http.StatusUnauthorized, message)
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
		"PublishService",
		strings.ToUpper("Put"),
		"/my-lab/app",
		publishService,
	},

	route{
		"ResetQuota",
		strings.ToUpper("Options"),
		"/my-lab/quota/reset",
		resetQuota,
	},

	route{
		"ResetQuota1",
		strings.ToUpper("Put"),
		"/my-lab/quota/reset",
		resetQuota1,
	},

	route{
		"ResetQuota2",
		strings.ToUpper("Post"),
		"/my-lab/quota/reset",
		resetQuota2,
	},

	route{
		"ResetQuota3",
		strings.ToUpper("Delete"),
		"/my-lab/quota/reset",
		resetQuota3,
	},

	route{
		"ResetQuota4",
		strings.ToUpper("Get"),
		"/my-lab/quota/reset",
		resetQuota4,
	},

	route{
		"ResetQuota5",
		strings.ToUpper("Head"),
		"/my-lab/quota/reset",
		resetQuota5,
	},

	route{
		"ResetQuota6",
		strings.ToUpper("Patch"),
		"/my-lab/quota/reset",
		resetQuota6,
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
