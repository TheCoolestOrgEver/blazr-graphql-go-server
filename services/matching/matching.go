package matching 

import (
	matchType "../../models/match"
	"../../models/location"
	matchDAO "../../daos/match"
	"github.com/rs/xid"
)

func SaveMatch( matcher string, matchee string ) matchType.Match {
	match := matchDAO.FindByUserID( matcher, matchee )

	if match == nil {
		// make new match
		matchID := xid.New()
		match := matchType.Match{
			MatchID: matchID.String(),
			UserA: matcher,
			UserB: matchee,
			Matched: false,
		}
		saved := matchDAO.Save( match )
		return saved
	}

	if !match.Matched && ( match.UserA == matchee ) && ( match.UserB == matcher ) {
		// update match to have matched true
		match.Matched = true
		match = matchDAO.Update( match )
	}

	//return the updated match
	return match
}

func CheckForMatch(matcher string, matchee string) bool {
	// do logic for checking if a match exists
	match := matchDAO.FindByUserID( matcher, matchee )
	if match != nil {
		// make a matcher object and save to mongo
		return match.Matched
	}
	return false
}