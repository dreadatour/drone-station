package dsgeo

import (
	"fmt"

	"github.com/mmcloughlin/geohash"
)

var (
	// ErrBadCoordinate is "Wrong coordinate, should be from 0 to 100" error
	ErrBadCoordinate = fmt.Errorf("Wrong coordinate, should be from 0 to 100")
	// ErrCoordinateOutOfRange is "Coordinate does not belongs to quadrant" error
	ErrCoordinateOutOfRange = fmt.Errorf("Coordinate does not belongs to quadrant")
	// ErrBadLatitude is "Wrong latitude, should be from -90 to 90" error
	ErrBadLatitude = fmt.Errorf("Wrong latitude, should be from -90 to 90")
	// ErrBadLongitude is "Wrong longitude, should be from -180 to 180" error
	ErrBadLongitude = fmt.Errorf("Wrong longitude, should be from -180 to 180")
)

// Quadrant is map area with bounding box coordinates and geohash string
type Quadrant interface {
	String() string
	Box() geohash.Box
	Precision() uint
	RelToAbs(x, y float64) (float64, float64, error)
	AbsToRel(latitude, longitude float64) (float64, float64, error)
	Contains(latitude, longitude float64) (bool, error)
}

// NewQuadrantFromString returns new quadrant from string
func NewQuadrantFromString(q string) Quadrant {
	return quadrant{
		q: q,
		b: geohash.BoundingBox(q),
	}
}

// NewQuadrantFromCoords returns new quadrant from latitude, longitude and precision
func NewQuadrantFromCoords(latitude, longitude float64, precision uint) Quadrant {
	q := geohash.EncodeWithPrecision(latitude, longitude, precision)
	return NewQuadrantFromString(q)
}

type quadrant struct {
	q string
	b geohash.Box
}

// String returns quadrant geohash
func (q quadrant) String() string {
	return q.q
}

// Box returns quadrant bounding box
func (q quadrant) Box() geohash.Box {
	return q.b
}

// Precision is quadrant geohash length
func (q quadrant) Precision() uint {
	return uint(len(q.q))
}

// RelToAbs converts relative coordinates within quadrant to absolute latitude and longitude
func (q quadrant) RelToAbs(x, y float64) (float64, float64, error) {
	if x < 0 || x > 100 || y < 0 || y > 100 {
		return 0, 0, ErrBadCoordinate
	}

	latitude := q.b.MinLat + (q.b.MaxLat-q.b.MinLat)*y/100
	longitude := q.b.MinLng + (q.b.MaxLng-q.b.MinLng)*x/100

	return latitude, longitude, nil
}

// AbsToRel converts latitude and longitude to quadrant relative coordinates
func (q quadrant) AbsToRel(latitude, longitude float64) (float64, float64, error) {
	contains, err := q.Contains(latitude, longitude)
	if err != nil {
		return 0, 0, err
	}
	if !contains {
		return 0, 0, ErrCoordinateOutOfRange
	}

	x := (longitude - q.b.MinLng) * 100 / (q.b.MaxLng - q.b.MinLng)
	y := (latitude - q.b.MinLat) * 100 / (q.b.MaxLat - q.b.MinLat)

	return x, y, nil
}

// Contains returns true if latitude, longitude coordinates are within quadrant
func (q quadrant) Contains(latitude, longitude float64) (bool, error) {
	if latitude < -90 || latitude > 90 {
		return false, ErrBadLatitude
	}
	if longitude < -180 || longitude > 180 {
		return false, ErrBadLongitude
	}

	return q.b.Contains(latitude, longitude), nil
}
