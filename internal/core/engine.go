package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
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
		if e.regions[i].Services == nil {
			err := fmt.Errorf("Driver is nil")
			result = append(result, err)
			break
		}
		err := e.regions[i].Services.Driver.Health()

		if err != nil {
			result = append(result, err)
		}
	}

	return result
}

// GetQuota provides the Quota objects for each region
func (e Engine) GetQuota(projectID string) (result []Quota, err error) {
	if len(projectID) == 0 {
		return result, fmt.Errorf("projectID is empty")
	}

	for i := range e.regions {
		log.Trace("Looking for project ’", projectID, "’ in region ’", e.regions[i].Name, "’")
		r, err := e.regions[i].Services.Driver.GetQuota(projectID)

		log.Debug(r)
		if err != nil {
			isNotFound, _ := regexp.MatchString("not found", err.Error())
			if !isNotFound {
				log.Error(err)
				return result, err
			}
			log.Trace("project not found, skipping error…")
			err = nil
		} else {
			result = append(result, r)
		}
	}

	return result, err
}
