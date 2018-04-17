package matchpool

import (
	matchType "../../models/match"
	matchpoolType "../../models/profile/matchpool"
	matchpoolDAO "../../daos/profile/matchpool"
)

func GetMatches( id string ) (error, matchpoolType.MatchPool) {
	err, matches := matchpoolDAO.FindOne(id)

	return err, matches
}

func CreateMatchPool( matchpool *matchpoolType.MatchPool ) (error, matchpoolType.MatchPool) {
	err, matches := matchpoolDAO.Save(matchpool)
	return err, matches
}

func DeleteProfile( id string ) (error, matchpoolType.MatchPool) {
	err, matches := matchpoolDAO.Remove(id) 
	return err, matches
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
	err, matches = matchpoolDAO.Update(&matches)
	return err, matches
}

