package profile

import ( 
	"../location/"
)

// BlazrProfile represents a profile
type BlazrProfile struct {
	UserID    string	
	Name      string	
	Age       int		
	Bio       string	
	ImageURL  string	
	Location  location.Coordinates
	MatchPool []BlazrProfile
}
