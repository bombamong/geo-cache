package geofunc

import (
	"fmt"
	"math"

	"github.com/bombamong/geo-cache/pkg/cache"
)

// CompareFunc is a callback function to compare rawData
type CompareFunc func(rd cache.RawData) bool

// Haversine is a function that returns the distance between
// two coordinates on spherical body.
// From play.golang.org
func Haversine(lonFrom float64, latFrom float64, lonTo float64, latTo float64) (distance float64) {
	const earthRadius = 6371.001

	var deltaLat = (latTo - latFrom) * (math.Pi / 180)
	var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)

	var a = math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(latFrom*(math.Pi/180))*math.Cos(latTo*(math.Pi/180))*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance = earthRadius * c

	return
}

//ContainsPoint is a Polygon method to evaluate whether a point is inside a
//polygon
func (pol Polygon) ContainsPoint(np Point) bool {
	//TODO: max / min for x and y
	if np.Latitude < pol.MinLat() || np.Latitude > pol.MaxLat() || np.Longitude < pol.MinLong() || np.Longitude > pol.MaxLong() {
		fmt.Println("Off limits")
		return false
	}
	var in = false
	for i := 0; i < len(pol.Vertices); i++ {
		cV := pol.Vertices[i]
		cN := pol.Vertices[(i+1)%len(pol.Vertices)]
		if rayIntersectsSegment(np, cV, cN) {
			in = !in
		}
	}
	return in
}

func rayIntersectsSegment(point Point, vertOne, vertTwo *Point) bool {
	fmt.Printf("info, vert1 Lat: %f  vert1 Long: %f vert2 Lat: %f vert2 Long: %f\n", vertOne.Latitude, vertOne.Longitude, vertTwo.Latitude, vertTwo.Longitude)
	if vertOne.Latitude > vertTwo.Latitude {
		vertOne, vertTwo = vertTwo, vertOne
	}
	for point.Latitude == vertOne.Latitude || point.Latitude == vertTwo.Latitude {
		point.Latitude = math.Nextafter(point.Latitude, math.Inf(1))
	}
	if point.Latitude < vertOne.Latitude {
		fmt.Printf("out of Latitude bounds Point:%f smaller than VertOne: %f\n", point.Latitude, vertOne.Latitude)
		return false
	}
	if point.Latitude > vertTwo.Latitude {
		fmt.Printf("out of Latitude bounds Point:%f greater than  VertTwo: %f\n", point.Latitude, vertTwo.Latitude)
	}
	if vertOne.Longitude > vertTwo.Longitude {
		if point.Longitude > vertOne.Longitude {
			fmt.Println("out of Longitude bounds")
			return false
		}
		if point.Longitude < vertTwo.Longitude {
			return true
		}
	} else {
		if point.Longitude > vertTwo.Longitude {
			fmt.Println("out of Longitude bounds")
			return false
		}
		if point.Longitude < vertOne.Longitude {
			return true
		}
	}
	return (point.Latitude-vertOne.Latitude)/(point.Longitude-vertOne.Longitude) >= (vertTwo.Latitude-vertOne.Latitude)/(vertTwo.Longitude-vertOne.Longitude)
}

func (pol Polygon) MinLat() float64 {
	min := pol.Vertices[0].Latitude
	for _, p := range pol.Vertices {
		if p.Latitude < min {
			min = p.Latitude
		}
	}
	return min
}

func (pol Polygon) MaxLat() float64 {
	max := pol.Vertices[0].Latitude
	for _, p := range pol.Vertices {
		if p.Latitude > max {
			max = p.Latitude
		}
	}
	return max
}

func (pol Polygon) MinLong() float64 {
	min := pol.Vertices[0].Longitude
	for _, p := range pol.Vertices {
		if p.Longitude < min {
			min = p.Longitude
		}
	}
	return min
}

func (pol Polygon) MaxLong() float64 {
	max := pol.Vertices[0].Longitude
	for _, p := range pol.Vertices {
		if p.Longitude > max {
			max = p.Longitude
		}
	}
	return max
}

type Point struct {
	Latitude  float64
	Longitude float64
}

type Polygon struct {
	Vertices []*Point
}
