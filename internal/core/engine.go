package core

type Engine struct {
	regions []Region
	catalogs []Catalog
}

func NewEngine(r []Region, c []Catalog) Engine {
	return Engine{
		regions: r,
		catalogs: c,
	}
}

func (e Engine) GetCatalogs() []Catalog{
	return e.catalogs
}

func (e Engine) GetRegions() []Region {
	r := e.regions

	for i, _ := range r {
		r[i].Services.Driver = nil
	}

	return r
}