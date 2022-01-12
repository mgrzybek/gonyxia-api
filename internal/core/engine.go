package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

// Engine is the main core object of the program.
type Engine struct {
	regions  []Region
	catalogs []Catalog
}

// NewEngine is the Engine’s constructor
func NewEngine(r []Region, c []Catalog) (Engine, error) {
	var err error

	if len(c) == 0 {
		err = fmt.Errorf("given catalogs are empty")
	}
	if len(r) == 0 {
		err = fmt.Errorf("given regions are empty")
	}

	return Engine{
		regions:  r,
		catalogs: c,
	}, err
}

// GetCatalogs returns an array of configured catalogs
func (e Engine) GetCatalogs() []Catalog {
	return e.catalogs
}

// GetCatalogByID returns the Catalog matching the given id.
// If no match is found, nil is returned.
func (e Engine) GetCatalogByID(id string) *Catalog {
	log.Trace("looking for catalog id ", id)
	for i := range e.catalogs {
		if e.catalogs[i].ID == id {
			return &e.catalogs[i]
		}
	}
	log.Trace("catalog id ", id, " not found")
	return nil
}

// GetRegions returns an array of the configured regions
func (e Engine) GetRegions() []Region {
	r := e.regions

	for i := range r {
		r[i].Services.Driver = nil
	}

	return r
}

// Health returns an array of errors is a region is unhealthy
func (e Engine) Health() (result []error) {
	for i := range e.regions {
		err := e.regions[i].Services.Driver.Health()

		if err != nil {
			result = append(result, err)
		}
	}

	return result
}
