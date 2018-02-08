package geolocation

import (
	"math"
	"../../models/location"
)

const EarthRadiusMiles = 3959.0
const LatMin = -1*math.Pi/2
const LatMax = math.Pi/2
const LongMin = -1*math.Pi
const LongMax = math.Pi

// Gets bounding coordinates for a profile
func GetMinMaxBounds( coordinates location.Coordinates, radiusMiles float64 ) (location.Coordinates, location.Coordinates) {
	if radiusMiles < 0 {
		panic("Incorrect radius")
	}
	// perform calculations for getting profiles within a radius here
	angularRadius := radiusMiles/EarthRadiusMiles
	// calculate min and max latitude 
	latMin := coordinates.Lat - angularRadius
	latMax := coordinates.Lat + angularRadius
	// calculate min and max longitude
	var longMin, longMax float64
	if latMin > LatMin && latMax < LatMax {
		deltaLong := math.Asin(math.Sin(angularRadius)/math.Cos(coordinates.Lat))
		longMin = coordinates.Long - deltaLong
		if longMin < LongMin {
			longMin += 2*math.Pi
		}
		longMax = coordinates.Long + deltaLong
		if longMax > LongMax {
			longMax -= 2*math.Pi
		}
	} else {
		// pole stuff
		latMin = math.Max(latMin, LatMin)
		latMax = math.Min(latMax, LatMax)
		longMin = LongMin
		longMax = LongMax
	}
	minCoords := location.Coordinates {
		Lat: latMin,
		Long: longMin,
	}
	maxCoords := location.Coordinates {
		Lat: latMax,
		Long: longMax,
	}
	return minCoords, maxCoords
}

// gets the distance between two points
func GetDistance( pointA location.Coordinates, pointB location.Coordinates ) float64 {
	return math.Acos( math.Sin(pointA.Lat) * math.Sin(pointB.Lat) + 
		   math.Cos(pointA.Lat) * math.Cos(pointB.Lat) * 
	       math.Cos(pointA.Long - pointB.Long)) * EarthRadiusMiles
}