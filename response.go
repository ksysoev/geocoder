package geocoder

type Ymaps struct {
	GeoObjectCollection geoObjectCollection `xml:"GeoObjectCollection"`
}

type geoObjectCollection struct {
	FeatureMembers           []featureMemberType          `xml:"featureMember"`
	GeocoderResponseMetaData geocoderResponseMetaDataType `xml:"metaDataProperty"`
}

type geocoderResponseMetaDataType struct {
	GeocoderResponseMetaData geocoderResponseMetaData `xml:"GeocoderResponseMetaData"`
}

type geocoderResponseMetaData struct {
	Request   string        `xml:"request"`
	Kind      string        `xml:"kind"`
	Found     int           `xml:"found"`
	Results   int           `xml:"results"`
	Skip      int           `xml:"skip"`
	BoundedBy boundedByType `xml:"boundedBy"`
}
type featureMemberType struct {
	GeoObject GeoObject `xml:"GeoObject"`
}
