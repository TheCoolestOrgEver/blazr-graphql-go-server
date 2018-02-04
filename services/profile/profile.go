package profile

import (
	profileTypes "../../models/profile"
	"../../models/location"
	profileDAO "../../daos/profile"
	"../geolocation"
	"github.com/rs/xid"
)

func GetProfile( id string ) profileTypes.BlazrProfile {
	return profileDAO.FindOne(id)
}

func CreateProfile( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	userID := xid.New()
	profile.UserID = userID.String();
	return profileDAO.Save(profile)
}

func GetProfiles( coordinates location.Coordinates, radiusMiles float64 ) []profileTypes.BlazrProfile {
	
	// use geolocation package to create bounds around our profile
	minCoordinates, maxCoordinates := geolocation.GetMinMaxBounds(coordinates, radiusMiles)

	return profileDAO.FindByCoordinatesBetween( minCoordinates, maxCoordinates )
}