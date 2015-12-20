package geocoder

import "testing"

func TestNewScope(t *testing.T) {
	lowerCorner, _ := NewGeoPoint("59.92 30.30")
	upperCorner, _ := NewGeoPoint("59.94 30.32")
	scope := NewScope(*lowerCorner, *upperCorner)
	if scope.lowerCorner.longitude != 59.92 || scope.lowerCorner.latitude != 30.30 || scope.upperCorner.longitude != 59.94 || scope.upperCorner.latitude != 30.32 {
		t.Fatal("Error create new scope")
	}
}

func TestCenter(t *testing.T) {
	lowerCorner, _ := NewGeoPoint("59.92 30.30")
	upperCorner, _ := NewGeoPoint("59.94 30.32")
	scope := NewScope(*lowerCorner, *upperCorner)
	scopeCenter := scope.Center()
	if scopeCenter.Longitude() != 59.929998979703235 || scopeCenter.Latitude() != 30.310000380213477 {
		t.Fatal("Error get center of scope")
	}
}

func TestSize(t *testing.T) {
	lowerCorner, _ := NewGeoPoint("59.92 30.30")
	upperCorner, _ := NewGeoPoint("59.94 30.32")
	scope := NewScope(*lowerCorner, *upperCorner)
	scopeSize := scope.Size()
	if scopeSize.Longitude() != 0.01999999999999602 || scopeSize.Latitude() != 0.019999999999999574 {
		t.Fatal("Error get size of scope")
	}
}
