package profile

import (
	profileTypes "../../models/profile"
	profileDAO "../../daos/profile"

	"github.com/rs/xid"
)

func GetProfile( id string ) profileTypes.BlazrProfile {
	return profileDAO.FindOne(id)
}

func CreateProfile( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	userID := xid.New()
	profile.UserID = userID.String();
	return profileDAO.InsertOne(profile)
}

func GetProfiles( coordinates profileTypes.Coordinates ) []profileTypes.BlazrProfile {
	// perform calculations for getting profiles within a radius here
	return profileDAO.FindAll("query")
}