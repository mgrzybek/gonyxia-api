package core

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Engine struct {
	regions  []Region
	catalogs []Catalog
}

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

func (e Engine) GetCatalogs() []Catalog {
	return e.catalogs
}

func (e Engine) GetCatalogById(id string) *Catalog {
	log.Trace("looking for catalog id ", id)
	for i, _ := range e.catalogs {
		if e.catalogs[i].Id == id {
			return &e.catalogs[i]
		}
	}
	log.Trace("catalog id ", id, " not found")
	return nil
}

func (e Engine) GetRegions() []Region {
	r := e.regions

	for i, _ := range r {
		r[i].Services.Driver = nil
	}

	return r
}
