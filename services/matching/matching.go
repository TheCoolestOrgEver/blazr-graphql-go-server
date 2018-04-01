package matching 

import (
	"log"
	matchType "../../models/match"
	matchDAO "../../daos/match"
	"github.com/rs/xid"
)

func SaveMatch( matcher string, matchee string ) matchType.Match {
	
	err, match := matchDAO.FindByUserID( matcher, matchee )

	if err != nil {
		// make new match
		matchID := xid.New()
		match = matchType.Match{
			MatchID: matchID.String(),
			UserA: matcher,
			UserB: matchee,
			Matched: false,
		}
		err, saved := matchDAO.Save( &match )
		if err!=nil {
			log.Fatal( err )
		}
		return saved
	}

	if !match.Matched && ( match.UserA == matchee ) && ( match.UserB == matcher ) {
		// update match to have matched true
		match.Matched = true
		err, match = matchDAO.Update( &match )
		if err!=nil {
			log.Fatal( err )
		}
	}

	//return the updated match
	return match
}

func CheckForMatch(matcher string, matchee string) bool {
	// do logic for checking if a match exists
	err, match := matchDAO.FindByUserID( matcher, matchee )
	if err == nil {
		// make a matcher object and save to mongo
		return match.Matched
	}
	return false
}