package profile

// BlazrProfile represents a profile
type BlazrProfile struct {
	ID        string
	Name      string
	Age       string
	Bio       string
	ImageURL  string
	MatchPool []BlazrProfile
}
