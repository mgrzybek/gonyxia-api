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
