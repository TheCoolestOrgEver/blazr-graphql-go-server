package profile

// coordinates represents coordinates
type Coordinates struct {
	Lat float64
	Long float64
}

// BlazrProfile represents a profile
type BlazrProfile struct {
	UserID    string
	Name      string
	Age       int
	Bio       string
	ImageURL  string
	Location  Coordinates 
	MatchPool []BlazrProfile
}
