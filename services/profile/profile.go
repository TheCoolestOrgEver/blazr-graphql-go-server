package profile

import (
	profileTypes "../../models/profile"
	profileDAO "../../daos/profile"
)

func GetProfile( id string ) profileTypes.BlazrProfile {
	return profileDAO.FindOne(id)
}

