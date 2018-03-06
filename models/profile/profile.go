package profile

import ( 
	"../location/"
)

// BlazrProfile represents a profile
type BlazrProfile struct {
	UserID    string				`json:"userID" bson:"userID"` 
	Name      string				`json:"name" bson:"name"`
	Age       int					`json:"age" bson:"age"`
	Bio       string				`json:"bio" bson:"bio"`
	ImageURL  string				`json:"imageURL" bson:"imageURL"`
	Location  location.Coordinates	`json:"location" bson:"location"`
	MatchPool []BlazrProfile		`json:"matchPool" bson:"matchPool`
}
