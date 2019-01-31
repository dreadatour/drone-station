package storage_test

import (
	"testing"

	"github.com/dreadatour/drone-station/pkg/dsgeo"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/dreadatour/drone-station/model"
	"github.com/dreadatour/drone-station/storage"
)

func TestNewDroneStorage(t *testing.T) {
	Convey("Creating new drone storage", t, func() {
		storage := storage.NewDronesStorage()

		Convey("Storage object should be created", func() {
			So(storage, ShouldNotBeNil)
		})
	})
}

func TestDroneStorageAdd(t *testing.T) {
	Convey("Given empty drone storage", t, func() {
		storage := storage.NewDronesStorage()

		Convey("Adding new drone to storage", func() {
			drone := storage.Add(model.Drone{
				Latitude:  51.924333714,
				Longitude: 4.477883541,
			})

			Convey("New drone should be created", func() {
				So(drone.ID, ShouldNotBeEmpty)

				Convey("Drones list should have this drone", func() {
					drones := storage.List()

					So(len(drones), ShouldEqual, 1)
					So(drones[0], ShouldResemble, *drone)
				})
			})
		})
	})
}

func TestDroneStorageListWithinQuadrant(t *testing.T) {
	Convey("Given drone storage with some drones inside", t, func() {
		storage := storage.NewDronesStorage()
		storage.Add(model.Drone{Latitude: 18, Longitude: 56})
		storage.Add(model.Drone{Latitude: -64, Longitude: 71})
		storage.Add(model.Drone{Latitude: 0, Longitude: 1})
		storage.Add(model.Drone{Latitude: 65, Longitude: -11})

		Convey("Storage must contains all these drones", func() {
			So(len(storage.List()), ShouldEqual, 4)
		})

		Convey("Adding new drone to storage", func() {
			drone := storage.Add(model.Drone{
				Latitude:  51.924333714,
				Longitude: 4.477883541,
			})

			Convey("New drone should be created", func() {
				So(drone.ID, ShouldNotBeEmpty)

				Convey("Drones list should have this drone in quadrant", func() {
					quadrant := dsgeo.NewQuadrantFromString("u15pmus9")
					drones := storage.ListWithinQuadrant(quadrant)

					So(len(drones), ShouldEqual, 1)
					So(drones[0], ShouldResemble, *drone)
				})

				Convey("Drones list should not have this drone in neighbour quadrant", func() {
					quadrant := dsgeo.NewQuadrantFromString("u15pmus8")
					drones := storage.ListWithinQuadrant(quadrant)

					So(len(drones), ShouldEqual, 0)
				})
			})
		})

		Convey("Adding new drone with incorrect coords", func() {
			drone := storage.Add(model.Drone{
				Latitude:  90.1,
				Longitude: 4.477883541,
			})

			Convey("New drone should be created", func() {
				So(drone.ID, ShouldNotBeEmpty)

				Convey("Drones list should panics while getting this drone", func() {
					quadrant := dsgeo.NewQuadrantFromString("u15pmus8")
					So(func() { storage.ListWithinQuadrant(quadrant) }, ShouldPanic)
				})
			})
		})
	})
}

func TestDroneStorageRemove(t *testing.T) {
	Convey("Given empty drone storage", t, func() {
		storage := storage.NewDronesStorage()

		Convey("Adding new drone to storage", func() {
			drone := storage.Add(model.Drone{
				Latitude:  51.924333714,
				Longitude: 4.477883541,
			})

			Convey("New drone should be created", func() {
				So(drone.ID, ShouldNotBeEmpty)

				Convey("Drones list should have this drone", func() {
					drones := storage.List()

					So(len(drones), ShouldEqual, 1)
					So(drones[0], ShouldResemble, *drone)
				})

				Convey("Removing this drone from list should be successful", func() {
					ok := storage.Remove(drone.ID)

					So(ok, ShouldBeTrue)
				})

				Convey("Removing nonexistent ID from list should not be successful", func() {
					ok := storage.Remove("foobar")

					So(ok, ShouldBeFalse)
				})
			})
		})
	})
}
