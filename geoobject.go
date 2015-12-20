package geocoder

import "regexp"

type GeoObject struct {
	MetaDataProperty metaDataPropertyType `xml:"metaDataProperty"`
	Description      string               `xml:"description"`
	Name             string               `xml:"name"`
	BoundedBy        boundedByType        `xml:"boundedBy"`
	Point            pointType            `xml:"Point"`
}

type metaDataPropertyType struct {
	GeocoderMetaData geocoderMetaDataType `xml:"GeocoderMetaData"`
}

type geocoderMetaDataType struct {
	Kind           string             `xml:"kind"`
	Text           string             `xml:"text"`
	Precision      string             `xml:"precision"`
	AddressDetails addressDetailsType `xml:"AddressDetails"`
}

type addressDetailsType struct {
	Country countryType `xml:"Country"`
}

type countryType struct {
	AddressLine        string                 `xml:"AddressLine"`
	CountryNameCode    string                 `xml:"CountryNameCode"`
	CountryName        string                 `xml:"CountryName"`
	AdministrativeArea administrativeAreaType `xml:"AdministrativeArea"`
}

type administrativeAreaType struct {
	AdministrativeAreaName string                    `xml:"AdministrativeAreaName"`
	SubAdministrativeArea  subAdministrativeAreaType `xml:"SubAdministrativeArea"`
}

type subAdministrativeAreaType struct {
	SubAdministrativeAreaName string       `xml:"SubAdministrativeAreaName"`
	Locality                  localityType `xml:"Locality"`
}

type localityType struct {
	LocalityName string           `xml:"LocalityName"`
	Thoroughfare thoroughfareType `xml:"Thoroughfare"`
}

type thoroughfareType struct {
	ThoroughfareName string      `xml:"ThoroughfareName"`
	Premise          PremiseType `xml:"Premise"`
}
type PremiseType struct {
	PremiseNumber string `xml:"PremiseNumber"`
}

type boundedByType struct {
	Envelope envelopeType `xml:"Envelope"`
}

type envelopeType struct {
	LowerCorner string `xml:"lowerCorner"`
	UpperCorner string `xml:"upperCorner"`
}

type pointType struct {
	Pos string `xml:"pos"`
}

//Kind return kind of GeoObject
func (g GeoObject) Kind() string {
	return g.MetaDataProperty.GeocoderMetaData.Kind
}

//GeoPoint return coordinates of GeoObject
func (g GeoObject) GeoPoint() (*GeoPoint, error) {
	geoPoint, err := NewGeoPoint(g.Point.Pos)
	if err != nil {
		return new(GeoPoint), err
	}
	return geoPoint, nil
}

//Address return address string of GeoObject
func (g GeoObject) Address() string {
	return g.MetaDataProperty.GeocoderMetaData.Text
}

//Country return country name of GeoObject
func (g GeoObject) Country() string {
	return g.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.CountryName
}

//CountryCode return country code of GeoObject
func (g GeoObject) CountryCode() string {
	return g.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.CountryNameCode
}

//AdministrativeArea return administrative area of GeoObject
func (g GeoObject) AdministrativeArea() string {
	return g.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.AdministrativeAreaName
}

//SubAdministrativeArea return subadministrative area of GeoObject
func (g GeoObject) SubAdministrativeArea() string {
	return g.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.SubAdministrativeArea.SubAdministrativeAreaName
}

//Locality return locality of GeoObject
func (g GeoObject) Locality() string {
	return g.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.SubAdministrativeArea.Locality.LocalityName
}

//Thoroughfare return thoroughfare of GeoObject
func (g GeoObject) Thoroughfare() string {
	return g.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.SubAdministrativeArea.Locality.Thoroughfare.ThoroughfareName
}

//Premise return premise of GeoObject
func (g GeoObject) Premise() string {
	return g.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.SubAdministrativeArea.Locality.Thoroughfare.Premise.PremiseNumber
}

//Scope return scope of GeoObject
func (g GeoObject) Scope() *Scope {
	changeOrder := regexp.MustCompile(`^(\S+)\s(\S+)$`)
	lowerCorner, _ := NewGeoPoint(changeOrder.ReplaceAllString(g.BoundedBy.Envelope.LowerCorner, "$2 $1"))
	upperCorner, _ := NewGeoPoint(changeOrder.ReplaceAllString(g.BoundedBy.Envelope.UpperCorner, "$2 $1"))
	scope := NewScope(*lowerCorner, *upperCorner)
	return scope
}
