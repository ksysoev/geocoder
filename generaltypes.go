package geocoder

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type GeoPoint struct {
	latitude  float64
	longitude float64
}

func NewGeoPoint(posString string) (*GeoPoint, error) {
	posValues := strings.Split(posString, " ")

	Longitude, err := strconv.ParseFloat(posValues[0], 64)
	if err != nil {
		err = errors.New("Error parsing longitude in string: '" + posString + "'")
		return new(GeoPoint), err
	}

	Latitude, err := strconv.ParseFloat(posValues[1], 64)
	if err != nil {
		err = errors.New("Error parsing latitude in string: '" + posString + "'")
		return new(GeoPoint), err
	}

	GeoPoint := GeoPoint{latitude: Latitude, longitude: Longitude}
	return &GeoPoint, nil
}

func (g GeoPoint) Latitude() float64 {
	return g.latitude
}

func (g GeoPoint) Longitude() float64 {
	return g.longitude
}

func (g GeoPoint) String() string {
	Latitude := strconv.FormatFloat(g.latitude, 'f', -1, 64)
	Longitude := strconv.FormatFloat(g.longitude, 'f', -1, 64)
	return Longitude + " " + Latitude
}

func (g GeoPoint) stringToScopeRequest() string {
	Latitude := strconv.FormatFloat(g.latitude, 'f', -1, 64)
	Longitude := strconv.FormatFloat(g.longitude, 'f', -1, 64)
	return Latitude + "," + Longitude
}

type Scope struct {
	lowerCorner GeoPoint
	upperCorner GeoPoint
}

type ScopeSize struct {
	GeoPoint
}

func NewScope(LowerCorner, UpperCorner GeoPoint) *Scope {
	Scope := Scope{lowerCorner: LowerCorner, upperCorner: UpperCorner}
	return &Scope
}

func (s Scope) Center() *GeoPoint {
	DeltaLongitude := (s.upperCorner.Longitude() - s.lowerCorner.Longitude()) * math.Pi / 180

	LowerLatitude := s.lowerCorner.Latitude() * math.Pi / 180
	LowerLongitude := s.lowerCorner.Longitude() * math.Pi / 180
	UpperLatitude := s.upperCorner.Latitude() * math.Pi / 180

	Bx := math.Cos(UpperLatitude) * math.Cos(DeltaLongitude)
	By := math.Cos(UpperLatitude) * math.Sin(DeltaLongitude)

	MidleLatitude := math.Atan2(math.Sin(LowerLatitude)+math.Sin(UpperLatitude), math.Sqrt((math.Cos(LowerLatitude)+Bx)*(math.Cos(LowerLatitude)+Bx)+By*By))
	MidleLongitude := LowerLongitude + math.Atan2(By, math.Cos(LowerLatitude)+Bx)

	DegreesMidleLatitude := MidleLatitude * 180 / math.Pi
	DegreesMidleLongitude := MidleLongitude * 180 / math.Pi

	GeoPoint := GeoPoint{latitude: DegreesMidleLatitude, longitude: DegreesMidleLongitude}

	return &GeoPoint
}

func (s Scope) Size() *ScopeSize {
	lenLatitude := math.Abs(s.lowerCorner.Latitude() - s.upperCorner.Latitude())
	lenLongitude := math.Abs(s.lowerCorner.Longitude() - s.upperCorner.Longitude())
	return &ScopeSize{GeoPoint{latitude: lenLatitude, longitude: lenLongitude}}

}
