package geolocation

import (
	"testing"
	"math"
	"../../models/location"
)

var (
	testTable []struct{}
	pControl location.Coordinates
	pNorthPole location.Coordinates
	pSouthPole location.Coordinates
	pMeridian location.Coordinates
	resultControlMin location.Coordinates
	resultControlMax location.Coordinates
	resultNorthPoleMin location.Coordinates
	resultNorthPoleMax location.Coordinates
	resultSouthPoleMin location.Coordinates
	resultSouthPoleMax location.Coordinates
	resultMeridianMin location.Coordinates
	resultMeridianMax location.Coordinates
)

func init() {

	// an average set of coordinates
	pControl = location.Coordinates {
		Lat: 1.3963,
		Long: -0.6981,
	}	

	pNorthPole = location.Coordinates {
		Lat: 1.5888,
		Long: 0,
	}

	pSouthPole = location.Coordinates {
		Lat: -1.5888,
		Long: 0,
	}

	pMeridian = location.Coordinates {
		Lat: 0,
		Long: 3.16,
	}

	resultControlMin = location.Coordinates {
		Lat: 1.2393,
		Long: -1.8186, 
	}

	resultControlMax = location.Coordinates {
		Lat: 1.5533,
		Long: 0.4224,
	}

	resultNorthPoleMin = location.Coordinates {
		Lat: 1.5812,
		Long: truncate(LongMin, 4), 
	}

	resultNorthPoleMax = location.Coordinates {
		Lat: 1.5708,
		Long: truncate(LongMax, 4), 
	}

	resultSouthPoleMin = location.Coordinates {
		Lat: -1.5708,
		Long: truncate(LongMin, 4), 
	}

	resultSouthPoleMax = location.Coordinates {
		Lat: -1.5812,
		Long: truncate(LongMax, 4),
	}

	resultMeridianMin = location.Coordinates {
		Lat: -0.0076,
		Long: 3.1524,
	}

	resultMeridianMax = location.Coordinates {
		Lat: 0.0076, 
		Long: -3.1156,
	}
}

func TestMinBoundsCorrect( t *testing.T ) {
	controlMin, _ := GetMinMaxBounds(pControl, 621.371)
	lat := truncate(controlMin.Lat, 4)
	long := truncate(controlMin.Long, 4)
	if lat != resultControlMin.Lat || long != resultControlMin.Long {
		t.Errorf("Min bounds calculation was incorrect got: %d, %d want: %d, %d", 
			lat, long, resultControlMin.Lat, resultControlMin.Long )
	}
}

func TestMinBoundsAtPoles( t *testing.T ) {
	nPoleMin, _ := GetMinMaxBounds(pNorthPole, 30.0)
	lat := truncate(nPoleMin.Lat, 4)
	long := truncate(nPoleMin.Long, 4)
	if lat != resultNorthPoleMin.Lat || long != resultNorthPoleMin.Long {
		t.Errorf("Min bounds calculation for North Pole was incorrect got: %d, %d want: %d, %d",
			lat, long, resultNorthPoleMin.Lat, resultNorthPoleMin.Long )
	}

	sPoleMin, _ := GetMinMaxBounds(pSouthPole, 30.0)
	lat = truncate(sPoleMin.Lat, 4)
	long = truncate(sPoleMin.Long, 4)
	if lat != resultSouthPoleMin.Lat || long != resultSouthPoleMin.Long {
		t.Errorf("Min bounds calculation for South Pole was incorrect got: %d, %d want: %d, %d",
			lat, long, resultSouthPoleMin.Lat, resultSouthPoleMin.Long )
	}
}

func TestMinBoundsAtMeridian( t *testing.T ) {
	mMin, _ := GetMinMaxBounds(pMeridian, 30.0)
	lat := truncate(mMin.Lat, 4)
	long := truncate(mMin.Long, 4)
	if lat != resultMeridianMin.Lat || long != resultMeridianMin.Long {
		t.Errorf("Min bounds calculation for 180th meridian was incorrect got: %d, %d want: %d, %d",
			lat, long, resultMeridianMin.Lat, resultMeridianMin.Long)
	}
}

func TestMaxBoundsCorrect( t *testing.T ) {
	_, controlMax := GetMinMaxBounds(pControl, 621.371)
	lat := truncate(controlMax.Lat, 4)
	long := truncate(controlMax.Long, 4)
	if lat != resultControlMax.Lat || long != resultControlMax.Long {
		t.Errorf("Max bounds calculation was incorrect got: %d, %d want: %d, %d", 
			lat, long, resultControlMax.Lat, resultControlMax.Long )
	}
}

func TestMaxBoundsAtPoles( t *testing.T ) {
	_, nPoleMax := GetMinMaxBounds(pNorthPole, 30.0)
	lat := truncate(nPoleMax.Lat, 4)
	long := truncate(nPoleMax.Long, 4)
	if lat != resultNorthPoleMax.Lat || long != resultNorthPoleMax.Long {
		t.Errorf("Max bounds calculation for North Pole was incorrect got: %d, %d want: %d, %d",
			lat, long, resultNorthPoleMax.Lat, resultNorthPoleMax.Long )
	}

	_, sPoleMax := GetMinMaxBounds(pSouthPole, 30.0)
	lat = truncate(sPoleMax.Lat, 4)
	long = truncate(sPoleMax.Long, 4)
	if lat != resultSouthPoleMax.Lat || long != resultSouthPoleMax.Long {
		t.Errorf("Max bounds calculation for South Pole was incorrect got: %d, %d want: %d, %d",
			lat, long, resultSouthPoleMax.Lat, resultSouthPoleMax.Long )
	}
}

func TestMaxBoundsAtMeridian( t *testing.T ) {
	_, mMax := GetMinMaxBounds(pMeridian, 30.0)
	lat := truncate(mMax.Lat, 4)
	long := truncate(mMax.Long, 4)
	if lat != resultMeridianMax.Lat || long != resultMeridianMax.Long {
		t.Errorf("Min bounds calculation for 180th meridian was incorrect got: %d, %d want: %d, %d",
			lat, long, resultMeridianMax.Lat, resultMeridianMax.Long)
	}
}

func TestIncorrectRadius( t *testing.T ) {
	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
	}()
	
	GetMinMaxBounds(pControl, -30.0)
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func truncate(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}