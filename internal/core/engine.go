package core

import "fmt"

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
	for i, _ := range e.catalogs {
		if e.catalogs[i].Id == id {
			return &e.catalogs[i]
		}
	}
	return nil
}

func (e Engine) GetRegions() []Region {
	r := e.regions

	for i, _ := range r {
		r[i].Services.Driver = nil
	}

	return r
}
