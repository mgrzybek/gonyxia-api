package rest

import (
	"testing"

	"bytes"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"

	"github.com/mgrzybek/gonyxia-api/internal/backoffice"
	"github.com/mgrzybek/gonyxia-api/internal/core"

	log "github.com/sirupsen/logrus"
)

const (
	validAuthHeader = "Bearer am9uaC5kb2UK.am9uaC5kb2UK.am9uaC5kb2UK"
)

func generateEngine() *core.Engine {
	log.Info("Starting generateEngine()…")
	var regions []core.Region
	kubeConfig := "/home/coder/.kube/config"
	log.Info("Creating the k8s driver…")
	k, err := backoffice.NewKubernetes(
		&kubeConfig,
	)
	if err != nil {
		log.Panic(err)
	}
	region := core.Region{
		ID:          "kub",
		Name:        "Kubernetes (in-cluster)",
		Description: "The in-cluster Kubernetes region.",
		Services:    &core.Services{},
	}

	log.Info("Assigning k8s driver to region…")
	region.Services.Driver = k

	regions = append(regions, region)

	var catalogs []core.Catalog
	catalog := core.Catalog{
		ID:          "inseefrlab-helm-charts-datascience",
		Name:        "Inseefrlab datascience",
		Description: "Services for datascientists. https://github.com/InseeFrLab/helm-charts-datascience",
		Maintainer:  "innovation@insee.fr",
		Location:    "https://inseefrlab.github.io/helm-charts-datascience",
		Status:      "PROD",
	}
	catalogs = append(catalogs, catalog)

	log.Info("Creating the engine…")
	result, _ := core.NewEngine(regions, catalogs)
	log.Info("Starting health()…")
	result.Health()

	log.Info("generateEngine() done.")

	return &result
}

func createRequest(t *testing.T, method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", validAuthHeader)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	return req
}

func TestGetQuotaWithoutNamespace(t *testing.T) {
	engine = generateEngine()
	w := httptest.NewRecorder()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/my-lab/quota/{projectID}", getQuota)
	http.Handle("/my-lab/quota/", router)

	req := createRequest(t, http.MethodGet, "/my-lab/quota/")

	router.ServeHTTP(w, req)
	if w.Result().StatusCode != 404 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(w.Result().Body)
		newStr := buf.String()

		t.Fatalf("HTTP code should be 404, got %d ; message: %s", w.Result().StatusCode, newStr)
	}
}

func TestGetQuotaWithUnknownNamespace(t *testing.T) {
	engine = generateEngine()
	w := httptest.NewRecorder()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/my-lab/quota/{projectID}", getQuota)
	//http.Handle("/my-lab/quota/", router)

	req := createRequest(t, http.MethodGet, "/my-lab/quota/unknown")

	router.ServeHTTP(w, req)
	if w.Result().StatusCode != 404 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(w.Result().Body)
		newStr := buf.String()

		t.Fatalf("HTTP code should be 404, got %d ; message: %s", w.Result().StatusCode, newStr)
	}
}
