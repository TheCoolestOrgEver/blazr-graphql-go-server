package profile

import (
	profileTypes "../../models/profile"
	"../../models/location"
	profileDAO "../../daos/profile"
	"../geolocation"
	//"github.com/rs/xid"
	"fmt"
	"math/rand"
)

func GetProfile( id string ) (error, profileTypes.BlazrProfile) {
	return profileDAO.FindOne(id)
}

func CreateProfile( profile *profileTypes.BlazrProfile ) (error, profileTypes.BlazrProfile) {
	//userID := xid.New()
	//profile.UserID = userID.String();
	return profileDAO.Save(profile)
}

func GetProfiles( coordinates location.Coordinates, radiusMiles float64 ) (error, []profileTypes.BlazrProfile) {
	
	// use geolocation package to create bounds around our profile
	minCoordinates, maxCoordinates := geolocation.GetMinMaxBounds(coordinates, radiusMiles)
	fmt.Print("min coords : %d, %d ", minCoordinates.Lat, minCoordinates.Long)
	fmt.Print("max coords : %d, %d ", maxCoordinates.Lat, maxCoordinates.Long)

	err, profiles := profileDAO.FindByCoordinatesBetween( minCoordinates, maxCoordinates )

	ShuffleProfiles(profiles)

	return err, profiles
}

func ShuffleProfiles( profiles []profileTypes.BlazrProfile ) {
	for i := 0; i < len(profiles); i++ {
		j := rand.Intn(i + 1)
		profiles[i], profiles[j] = profiles[j], profiles[i]
	}
}

func DeleteProfile( id string ) (error, profileTypes.BlazrProfile) {
	return profileDAO.Remove(id)
}

func UpdateProfile( profile *profileTypes.BlazrProfile ) (error, profileTypes.BlazrProfile) {
	return profileDAO.Update( profile )
}

func UpdateLocation( userID string, lat float64, long float64 ) (error, profileTypes.BlazrProfile) {
	err, toUpdate := profileDAO.FindOne( userID )
	if err != nil {
		fmt.Println(err)
	}
	coordinates := location.Coordinates {
		lat,
		long,
	}
	toUpdate.Location = coordinates
	return profileDAO.Update( &toUpdate )
}