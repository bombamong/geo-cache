package geofunc

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Geofuncs(t *testing.T) {
	Convey("Should work properly", t, func() {
		Convey("Haversine", func() {
			p1 := Point{
				Latitude:  0.0,
				Longitude: 0.0,
			}
			p2 := Point{
				Latitude:  5.0,
				Longitude: 5.0,
			}
			dist := Haversine(p1.Longitude, p1.Latitude, p2.Longitude, p2.Latitude)
			So(dist, ShouldAlmostEqual, 785.76833085717, 0.01)
		})
		Convey("Polygon contains point", func() {
			Convey("Should handle square", func() {
				points := []*Point{
					{Latitude: 1, Longitude: 2},
					{Latitude: 2, Longitude: 1},
					{Latitude: 0, Longitude: 0},
					{Latitude: -2, Longitude: -2},
				}

				polygon := Polygon{
					Vertices: []*Point{
						{0, 0},
						{10, 0},
						{10, 10},
						{0, 10},
					},
				}
				contains := polygon.ContainsPoint(*points[0])
				So(contains, ShouldBeTrue)

				contains = polygon.ContainsPoint(*points[1])
				So(contains, ShouldBeTrue)

				contains = polygon.ContainsPoint(*points[2])
				So(contains, ShouldBeTrue)

				contains = polygon.ContainsPoint(*points[3])
				So(contains, ShouldBeFalse)
			})
		})

		Convey("Should handle strange polygons", func() {
			points := []*Point{
				{Latitude: 1, Longitude: 2},
				{Latitude: 2, Longitude: 1},
				{Latitude: 0, Longitude: 0},
				{Latitude: -2, Longitude: -2},
				{Latitude: -2, Longitude: -2},
				{Latitude: -2, Longitude: -2},
				{Latitude: -2, Longitude: -2},
				{Latitude: -2, Longitude: -2},
			}

			polygon := Polygon{
				Vertices: []*Point{
					{0, 0},
					{10, 0},
					{10, 10},
					{0, 10},
				},
			}
			contains := polygon.ContainsPoint(*points[0])
			So(contains, ShouldBeTrue)

			contains = polygon.ContainsPoint(*points[1])
			So(contains, ShouldBeTrue)

			contains = polygon.ContainsPoint(*points[2])
			So(contains, ShouldBeTrue)

			contains = polygon.ContainsPoint(*points[3])
			So(contains, ShouldBeFalse)
		})

	})
	return
}
