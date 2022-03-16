package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
)

// Engine is the main core object of the program.
type Engine struct {
	regions  Regions
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
		regions: func(region []Region) Regions {
			result := Regions{}
			for i := range region {
				result[region[i].ID] = region[i]
			}
			return result
		}(r),
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
func (e Engine) GetRegions() (result []Region) {
	r := e.regions

	for _, v := range r {
		result = append(result, v)
		result[len(result)-1].Services.Driver = nil
	}

	return result
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

// GetQuotas provides the Quota objects for each region
func (e Engine) GetQuotas(projectID string) (result QuotaPerRegion, err error) {
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
			result[e.regions[i].Name] = r
		}
	}

	return result, err
}

// GetQuota provides the Quota object for the given project in a region
func (e Engine) GetQuota(projectID, regionID string) (result Quota, err error) {
	if len(projectID) == 0 {
		return result, fmt.Errorf("projectID is empty")
	}

	if len(regionID) == 0 {
		return result, fmt.Errorf("regionID is empty")
	}

	r, doesKeyExist := e.regions[regionID]
	if !doesKeyExist {
		return result, fmt.Errorf("regionID not found")
	}

	return r.Services.Driver.GetQuota(projectID)
}

// SetQuota sets the given quota in all regions that provides the project
func (e Engine) SetQuota(quota Quota, projectID string) (err error) {
	if len(projectID) == 0 {
		return fmt.Errorf("projectID is empty")
	}

	existingQuotas, err := e.GetQuotas(projectID)
	if err != nil {
		return err
	}

	for r, q := range existingQuotas {
		if q != quota {
			err := e.regions[r].Services.Driver.SetQuota(quota, projectID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
