package matchpool

import (
	matchType "../../models/match"
	matchpoolType "../../models/profile/matchpool"
	matchpoolDAO "../../daos/profile/matchpool"
)

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

