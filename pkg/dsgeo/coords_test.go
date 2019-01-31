package dsgeo_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/pkg/dsgeo"
)

func TestNewQuadrantFromString(t *testing.T) {
	Convey("Given quadrant string", t, func() {
		q := "u15pmus9"

		Convey("Creating new Quadrant object from string", func() {
			quadrant := dsgeo.NewQuadrantFromString(q)

			Convey("Quadrant object should be created", func() {
				So(quadrant, ShouldNotBeNil)
			})

			Convey("Quadrant string representation should be correct", func() {
				So(quadrant.String(), ShouldEqual, q)
			})

			Convey("Quadrant boundary box coordinates should be correct", func() {
				So(fmt.Sprintf("%.6f", quadrant.Box().MinLat), ShouldEqual, fmt.Sprintf("%.6f", 51.924304962))
				So(fmt.Sprintf("%.6f", quadrant.Box().MaxLat), ShouldEqual, fmt.Sprintf("%.6f", 51.924476624))
				So(fmt.Sprintf("%.6f", quadrant.Box().MinLng), ShouldEqual, fmt.Sprintf("%.6f", 4.477615356))
				So(fmt.Sprintf("%.6f", quadrant.Box().MaxLng), ShouldEqual, fmt.Sprintf("%.6f", 4.477958679))
			})

			Convey("Quadrant precision should be correct", func() {
				So(quadrant.Precision(), ShouldEqual, 8)
			})
		})
	})
}

func TestNewQuadrantFromCoords(t *testing.T) {
	Convey("Given geo coordinates", t, func() {
		latitude := 51.92442
		longitude := 4.477736
		precision := uint(8)

		Convey("Creating new Quadrant object from coordinates", func() {
			quadrant := dsgeo.NewQuadrantFromCoords(latitude, longitude, precision)

			Convey("Quadrant object should be created", func() {
				So(quadrant, ShouldNotBeNil)
			})

			Convey("Quadrant string representation should be correct", func() {
				So(quadrant.String(), ShouldEqual, "u15pmus9")
			})

			Convey("Quadrant boundary box coordinates should be correct", func() {
				So(fmt.Sprintf("%.6f", quadrant.Box().MinLat), ShouldEqual, fmt.Sprintf("%.6f", 51.924304962))
				So(fmt.Sprintf("%.6f", quadrant.Box().MaxLat), ShouldEqual, fmt.Sprintf("%.6f", 51.924476624))
				So(fmt.Sprintf("%.6f", quadrant.Box().MinLng), ShouldEqual, fmt.Sprintf("%.6f", 4.477615356))
				So(fmt.Sprintf("%.6f", quadrant.Box().MaxLng), ShouldEqual, fmt.Sprintf("%.6f", 4.477958679))
			})

			Convey("Quadrant precision should be correct", func() {
				So(quadrant.Precision(), ShouldEqual, precision)
			})
		})
	})
}

func TestQuadrantConversion(t *testing.T) {
	tests := []struct {
		x         float64
		y         float64
		latitude  float64
		longitude float64
	}{
		{
			x:         0,
			y:         0,
			latitude:  51.924304963,
			longitude: 4.477615357,
		},
		{
			x:         78.11,
			y:         16.75,
			latitude:  51.924333714,
			longitude: 4.477883541,
		},
		{
			x:         50,
			y:         50,
			latitude:  51.924390793,
			longitude: 4.477787018,
		},
		{
			x:         100,
			y:         100,
			latitude:  51.924476623,
			longitude: 4.477958678,
		},
	}

	Convey("Given quadrant", t, func() {
		quadrant := dsgeo.NewQuadrantFromString("u15pmus9")

		Convey("Converting relative coordinates to absolute", func() {
			Convey("Coortidantes should matches", func() {
				for _, test := range tests {
					latitude, longitude, err := quadrant.RelToAbs(test.x, test.y)
					So(err, ShouldBeNil)
					So(fmt.Sprintf("%.6f", latitude), ShouldEqual, fmt.Sprintf("%.6f", test.latitude))
					So(fmt.Sprintf("%.6f", longitude), ShouldEqual, fmt.Sprintf("%.6f", test.longitude))
				}
			})
		})

		Convey("Converting absolute coordinates to relative", func() {
			Convey("Coortidantes should matches", func() {
				for _, test := range tests {
					x, y, err := quadrant.AbsToRel(test.latitude, test.longitude)
					So(err, ShouldBeNil)
					So(fmt.Sprintf("%.2f", x), ShouldEqual, fmt.Sprintf("%.2f", test.x))
					So(fmt.Sprintf("%.2f", y), ShouldEqual, fmt.Sprintf("%.2f", test.y))
				}
			})
		})
	})
}

func TestQuadrantContains(t *testing.T) {
	tests := []struct {
		latitude  float64
		longitude float64
		contains  bool
	}{
		{
			latitude:  51.924304963,
			longitude: 4.477615357,
			contains:  true,
		},
		{
			latitude:  51.924333714,
			longitude: 4.477883541,
			contains:  true,
		},
		{
			latitude:  51.924390793,
			longitude: 4.477787018,
			contains:  true,
		},
		{
			latitude:  51.924476623,
			longitude: 4.477958678,
			contains:  true,
		},
		{
			latitude:  51.914304963,
			longitude: 4.477615357,
			contains:  false,
		},
		{
			latitude:  51.924304963,
			longitude: 4.467615357,
			contains:  false,
		},
		{
			latitude:  51.934476623,
			longitude: 4.477958678,
			contains:  false,
		},
		{
			latitude:  51.924476623,
			longitude: 4.487958678,
			contains:  false,
		},
	}

	Convey("Given quadrant", t, func() {
		quadrant := dsgeo.NewQuadrantFromString("u15pmus9")

		Convey("Checking if coordinates belongs to quadrant boundary box", func() {
			Convey("Result should be correct", func() {
				for _, test := range tests {
					contains, err := quadrant.Contains(test.latitude, test.longitude)
					So(err, ShouldBeNil)
					So(contains, ShouldEqual, test.contains)
				}
			})
		})
	})
}

func TestQuadrantRelToAbsErrors(t *testing.T) {
	tests := []struct {
		x float64
		y float64
	}{
		{
			x: -1,
			y: 12,
		},
		{
			x: 78,
			y: -18,
		},
		{
			x: 100.1,
			y: 98,
		},
		{
			x: 2,
			y: 999,
		},
	}

	Convey("Given quadrant", t, func() {
		quadrant := dsgeo.NewQuadrantFromString("u15pmus9")

		Convey("Converting wrong relative coordinates to absolute", func() {
			Convey("Error should be returned", func() {
				for _, test := range tests {
					latitude, longitude, err := quadrant.RelToAbs(test.x, test.y)
					So(err, ShouldEqual, dsgeo.ErrBadCoordinate)
					So(latitude, ShouldEqual, 0)
					So(longitude, ShouldEqual, 0)
				}
			})
		})
	})
}

func TestQuadrantAbsToRelErrors(t *testing.T) {
	tests := []struct {
		latitude  float64
		longitude float64
		err       error
	}{
		{
			latitude:  51.914304963,
			longitude: 4.477615357,
			err:       dsgeo.ErrCoordinateOutOfRange,
		},
		{
			latitude:  51.924304963,
			longitude: 4.467615357,
			err:       dsgeo.ErrCoordinateOutOfRange,
		},
		{
			latitude:  51.934476623,
			longitude: 4.477958678,
			err:       dsgeo.ErrCoordinateOutOfRange,
		},
		{
			latitude:  51.924476623,
			longitude: 4.487958678,
			err:       dsgeo.ErrCoordinateOutOfRange,
		},
		{
			latitude:  -90.1,
			longitude: 0,
			err:       dsgeo.ErrBadLatitude,
		},
		{
			latitude:  90.1,
			longitude: 0,
			err:       dsgeo.ErrBadLatitude,
		},
		{
			latitude:  0,
			longitude: -180.1,
			err:       dsgeo.ErrBadLongitude,
		},
		{
			latitude:  0,
			longitude: 180.1,
			err:       dsgeo.ErrBadLongitude,
		},
	}

	Convey("Given quadrant", t, func() {
		quadrant := dsgeo.NewQuadrantFromString("u15pmus9")

		Convey("Converting wrong absolute coordinates to relative", func() {
			Convey("Error should be returned", func() {
				for _, test := range tests {
					x, y, err := quadrant.AbsToRel(test.latitude, test.longitude)
					So(err, ShouldEqual, test.err)
					So(x, ShouldEqual, 0)
					So(y, ShouldEqual, 0)
				}
			})
		})
	})
}

func TestQuadrantContainsErrors(t *testing.T) {
	tests := []struct {
		latitude  float64
		longitude float64
		err       error
	}{
		{
			latitude:  -90.1,
			longitude: 0,
			err:       dsgeo.ErrBadLatitude,
		},
		{
			latitude:  90.1,
			longitude: 0,
			err:       dsgeo.ErrBadLatitude,
		},
		{
			latitude:  0,
			longitude: -180.1,
			err:       dsgeo.ErrBadLongitude,
		},
		{
			latitude:  0,
			longitude: 180.1,
			err:       dsgeo.ErrBadLongitude,
		},
	}

	Convey("Given quadrant", t, func() {
		quadrant := dsgeo.NewQuadrantFromString("u15pmus9")

		Convey("Checking if invalid coordinates belongs to quadrant boundary box", func() {
			Convey("Error should be returned", func() {
				for _, test := range tests {
					contains, err := quadrant.Contains(test.latitude, test.longitude)
					So(err, ShouldEqual, test.err)
					So(contains, ShouldEqual, false)
				}
			})
		})
	})
}
