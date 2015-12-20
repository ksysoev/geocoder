package geocoder

import "testing"

func TestNewGeoPoint(t *testing.T) {
	geoPoint, err := NewGeoPoint("59.939095 30.315868")
	if err != nil {
		t.Fatal("Error parsing geo coordinates")
	}

	if geoPoint.longitude != 59.939095 {
		t.Fatal("Error parsing longitude")
	}

	if geoPoint.latitude != 30.315868 {
		t.Fatal("Error parsing latitude")
	}
}

func TestLongitude(t *testing.T) {
	geoPoint, _ := NewGeoPoint("59.939095 30.315868")
	if geoPoint.Longitude() != 59.939095 {
		t.Fatal("Error return longitude")
	}
}

func TestLatitude(t *testing.T) {
	geoPoint, _ := NewGeoPoint("59.939095 30.315868")
	if geoPoint.Latitude() != 30.315868 {
		t.Fatal("Error return latitude")
	}
}

func TestString(t *testing.T) {
	geoPoint, _ := NewGeoPoint("59.939095 30.315868")
	if geoPoint.String() != "59.939095 30.315868" {
		t.Fatal("Error return coordinates string")
	}
}
