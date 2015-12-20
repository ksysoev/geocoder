package geocoder

//GeoCoder - base structure to work with package. It help you to store configuration between requests.
type GeoCoder struct {
	apikey string
}

//New -  create new GeoCoder object.
func New() *GeoCoder {
	geoCoder := new(GeoCoder)
	return geoCoder
}

//FindOne send request to server and retern only one result.
func (g GeoCoder) FindOne(requestString string) GeoObject {
	return FindOne(requestString)
}

//FindOneReverse send reverse request to server and retern only one result.
func (g GeoCoder) FindOneReverse(N, E float64, kind string) GeoObject {
	return FindOneReverse(N, E, kind)
}

//FindOneFromScope send request to server and retern only one result placed in scope.
func (g GeoCoder) FindOneFromScope(requestString string, scope Scope) GeoObject {
	return FindOneFromScope(requestString, scope)
}

// func (g GeoCoder) Find(requestString string, scope Scope) хъGeoObject {
// 	return FindOneFromScope(requestString, scope)
// }
