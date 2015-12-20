package geocoder

import "testing"

var TestGeoObject = GeoObject{
	MetaDataProperty: metaDataPropertyType{
		GeocoderMetaData: geocoderMetaDataType{
			Kind:      "house",
			Text:      "Россия, Москва, Тверская улица, 7",
			Precision: "exact",
			AddressDetails: addressDetailsType{
				Country: countryType{
					AddressLine:     "Москва, Тверская улица, 7",
					CountryNameCode: "RU",
					CountryName:     "Россия",
					AdministrativeArea: administrativeAreaType{
						AdministrativeAreaName: "Центральный федеральный округ",
						SubAdministrativeArea: subAdministrativeAreaType{
							SubAdministrativeAreaName: "Москва",
							Locality: localityType{
								LocalityName: "Москва",
								Thoroughfare: thoroughfareType{
									ThoroughfareName: "Тверская улица",
									Premise: PremiseType{
										PremiseNumber: "7",
									},
								},
							},
						},
					},
				},
			},
		},
	},
	Description: "Москва, Россия",
	Name:        "Тверская улица, 7",
	BoundedBy: boundedByType{
		Envelope: envelopeType{
			LowerCorner: "37.602777 55.753321",
			UpperCorner: "37.619234 55.762601",
		},
	},
	Point: pointType{
		Pos: "37.611006 55.757962",
	},
}

func TestKind(t *testing.T) {
	if TestGeoObject.Kind() != TestGeoObject.MetaDataProperty.GeocoderMetaData.Kind {
		t.Fatal("Error get kind from geo object")
	}
}

func TestGeoPoint(t *testing.T) {
	if TestGeoObject.Kind() != TestGeoObject.MetaDataProperty.GeocoderMetaData.Kind {
		t.Fatal("Error get kind from geo object")
	}
}

func TestAddress(t *testing.T) {
	if TestGeoObject.Address() != TestGeoObject.MetaDataProperty.GeocoderMetaData.Text {
		t.Fatal("Error get address from geo object")
	}
}

func TestCountry(t *testing.T) {
	if TestGeoObject.Country() != TestGeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.CountryName {
		t.Fatal("Error get country from geo object")
	}
}

func TestCountryCode(t *testing.T) {
	if TestGeoObject.CountryCode() != TestGeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.CountryNameCode {
		t.Fatal("Error get country code from geo object")
	}
}

func TestAdministrativeArea(t *testing.T) {
	if TestGeoObject.AdministrativeArea() != TestGeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.AdministrativeAreaName {
		t.Fatal("Error get administrative area from geo object")
	}
}

func TestSubAdministrativeArea(t *testing.T) {
	if TestGeoObject.SubAdministrativeArea() != TestGeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.SubAdministrativeArea.SubAdministrativeAreaName {
		t.Fatal("Error get subadministrative area from geo object")
	}
}

func TestLocality(t *testing.T) {
	if TestGeoObject.Locality() != TestGeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.SubAdministrativeArea.Locality.LocalityName {
		t.Fatal("Error get locality from geo object")
	}
}

func TestThoroughfare(t *testing.T) {
	if TestGeoObject.Thoroughfare() != TestGeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.SubAdministrativeArea.Locality.Thoroughfare.ThoroughfareName {
		t.Fatal("Error get thoroughfare from geo object")
	}
}

func TestPremise(t *testing.T) {
	if TestGeoObject.Premise() != TestGeoObject.MetaDataProperty.GeocoderMetaData.AddressDetails.Country.AdministrativeArea.SubAdministrativeArea.Locality.Thoroughfare.Premise.PremiseNumber {
		t.Fatal("Error get kind from geo object")
	}
}

func TestScope(t *testing.T) {
	TestScope := TestGeoObject.Scope()
	if TestScope.lowerCorner.longitude != 55.753321 ||
		TestScope.lowerCorner.latitude != 37.602777 ||
		TestScope.upperCorner.longitude != 55.762601 ||
		TestScope.upperCorner.latitude != 37.619234 {
		t.Fatal("Error get scope from geo object")
	}
}
