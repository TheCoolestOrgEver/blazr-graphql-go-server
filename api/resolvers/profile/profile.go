package profile

import (
	profileTypes "../../../models/profile"
	profileService "../../../services/profile"
)

var (
	Kevin profileTypes.BlazrProfile
)

func GetProfile( id string ) profileTypes.BlazrProfile {
	return profileService.GetProfile(id)
}
