package geocoder

type GeoCoder struct {
	apikey string
}

func New() *GeoCoder {
	geoCoder := new(GeoCoder)
	return geoCoder
}

func (g GeoCoder) FindOne(requestString string) GeoObject {
	return FindOne(requestString)
}

func (g GeoCoder) FindOneReverse(N, E float64, kind string) GeoObject {
	return FindOneReverse(N, E, kind)
}

func (g GeoCoder) FindOneFromScope(requestString string, scope Scope) GeoObject {
	return FindOneFromScope(requestString, scope)
}

// func (g GeoCoder) Find(requestString string, scope Scope) хъGeoObject {
// 	return FindOneFromScope(requestString, scope)
// }
