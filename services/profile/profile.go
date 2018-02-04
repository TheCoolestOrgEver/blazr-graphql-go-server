package profile

import (
	profileTypes "../../models/profile"
	profileDAO "../../daos/profile"
	"math"
	"github.com/rs/xid"
)

const EarthRadiusMiles = 3959.0

func GetProfile( id string ) profileTypes.BlazrProfile {
	return profileDAO.FindOne(id)
}

func CreateProfile( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	userID := xid.New()
	profile.UserID = userID.String();
	return profileDAO.InsertOne(profile)
}

func GetProfiles( coordinates profileTypes.Coordinates, radiusMiles float64 ) []profileTypes.BlazrProfile {
	
	return profileDAO.FindAll("query")
}

func getMinMaxBounds( coordinates profileTypes.Coordinates, radiusMiles float64 ) (profileTypes.Coordinates, profileTypes.Coordinates) {
		// perform calculations for getting profiles within a radius here
	
		angularRadius := radiusMiles/EarthRadiusMiles
	
		// calculate min and max latitude (this is straightforward)
	
		latMin := coordinates.Lat - angularRadius
		latMax := coordinates.Long + angularRadius
	
		// calculate min and max longitude (complicated formula)

		deltaLong := math.Asin(math.Sin(angularRadius)/math.Cos(coordinates.Lat))

		lonMin := coordinates.Long- deltaLong
		lonMax := coordinates.Long + deltaLong

		minCoords := profileTypes.Coordinates {
			Lat: latMin,
			Long: lonMin,
		}

		maxCoords := profileTypes.Coordinates {
			Lat: latMax,
			Long: lonMax,
		}

		return minCoords, maxCoords
}