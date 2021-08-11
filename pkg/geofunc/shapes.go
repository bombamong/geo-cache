package geofunc

import (
	"math"

	"github.com/bombamong/geo-cache/pkg/cache"
)

type CompareFunc func(rd cache.RawData) bool

// Haversine is a function that returns the distance between
// two coordinates on spherical body.
// From play.golang.org
func Haversine(lonFrom float64, latFrom float64, lonTo float64, latTo float64) (distance float64) {
	const earthRadius = float64(6371)

	var deltaLat = (latTo - latFrom) * (math.Pi / 180)
	var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)

	var a = math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(latFrom*(math.Pi/180))*math.Cos(latTo*(math.Pi/180))*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance = earthRadius * c

	return
}

func (pol Polygon) ContainsPoint(p Point) {
	//TODO: max / min for x and y
	//TODO: find bound from all connecting Points of Polygon
	//TODO: raycasting from p to find all collisions

}

type Point struct {
	Latitude  float64
	Longitude float64
}

type Polygon struct {
	Points []*Point
}
