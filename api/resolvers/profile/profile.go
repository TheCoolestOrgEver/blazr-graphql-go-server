package profile

import (
	profileTypes "../../../models/profile"
	profileService "../../../services/profile"
)

func GetProfile( id string ) profileTypes.BlazrProfile {
	return profileService.GetProfile(id)
}

func CreateProfile( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	return profileService.CreateProfile(profile)
}

func GetProfiles( coordinates profileTypes.Coordinates ) []profileTypes.BlazrProfile {
	return profileService.GetProfiles(coordinates)
}
