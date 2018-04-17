package matchpool

import (
	"../../match"
)

type MatchPool struct {
	UserID  string `json:"userID" bson:"userID"` 
	Matches	[]match.Match
}