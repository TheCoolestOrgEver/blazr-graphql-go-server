package profile

import (
	profileTypes "../../models/profile"
	"../../models/location"
	profileDAO "../../daos/profile"
	"../geolocation"
	"github.com/rs/xid"
	"fmt"
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
	fmt.Print("min coords : %d, %d ", minCoordinates.Lat, minCoordinates.Long)
	fmt.Print("max coords : %d, %d ", maxCoordinates.Lat, maxCoordinates.Long)

	return profileDAO.FindByCoordinatesBetween( minCoordinates, maxCoordinates )
}

func DeleteProfile( id string ) profileTypes.BlazrProfile {
	return profileDAO.Remove(id)
}

func UpdateProfile( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	return profileDAO.Update( profile )
}