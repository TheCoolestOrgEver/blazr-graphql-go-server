package matchpool

import (
	matchType "../../models/match"
	matchpoolType "../../models/profile/matchpool"
	matchpoolDAO "../../daos/profile/matchpool"
	profileType "../../models/profile"
	profileDAO "../../daos/profile"
)

func GetMatchedProfiles( id string ) (error, []profileType.BlazrProfile) {
	//return matchpoolDAO.FindOne(id)
	// do some real shit here
	err, matches := matchpoolDAO.FindOne(id)
	if err != nil {
		panic(err)
	}
	var ids []string
	for i := 0; i < len(matches.Matches); i++ {
		if matches.Matches[i].UserA == id {
			ids = append( ids, matches.Matches[i].UserB )
		} else {
			ids = append( ids, matches.Matches[i].UserA )
		}	
	}

	// now do a find all of all these ids
	return profileDAO.FetchMatches(ids);
}

func GetMatches( id string ) (error, matchpoolType.MatchPool) {
	return matchpoolDAO.FindOne(id)
}

func CreateMatchPool( matchpool *matchpoolType.MatchPool ) (error, matchpoolType.MatchPool) {
	return matchpoolDAO.Save(matchpool)
}

func DeleteProfile( id string ) (error, matchpoolType.MatchPool) {
	return matchpoolDAO.Remove(id) 
}

func AddMatchToMatchPool( userID string, match *matchType.Match ) (error, matchpoolType.MatchPool) {

	err, matches := GetMatches( userID )
	if err!= nil {
		m := matchpoolType.MatchPool {
			UserID: userID,
			Matches: []matchType.Match{ *match },
		}
		CreateMatchPool( &m )
	}
	matches.Matches = append(matches.Matches, *match)
	return matchpoolDAO.Update(&matches)
}

