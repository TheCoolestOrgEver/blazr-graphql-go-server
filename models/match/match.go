package match

type Match struct {
	MatchID string `json:"matchID" bson:"matchID"`
	UserA string   `json:"userA" bson:"userA"`
	UserB string   `json:"userB" bson:"userB"`
	Matched bool   `json:"matched" bson:"matched"`
}

