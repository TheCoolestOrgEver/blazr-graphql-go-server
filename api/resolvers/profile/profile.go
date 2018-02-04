package profile

import (
	profileTypes "../../../models/profile"
	"../../../models/location"
	profileService "../../../services/profile"
)

func GetProfile( id string ) profileTypes.BlazrProfile {
	return profileService.GetProfile( id )
}

func CreateProfile( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	return profileService.CreateProfile( profile ) 
}

func GetProfiles( coordinates location.Coordinates, radiusMiles float64 ) []profileTypes.BlazrProfile {
	return profileService.GetProfiles( coordinates, radiusMiles )
}
