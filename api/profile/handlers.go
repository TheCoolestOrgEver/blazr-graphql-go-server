package handlers

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
	"strconv"
	"fmt"
	"encoding/json"
	"../../services/profile"
	"../../models/location"
	profileTypes "../../models/profile"
)

func GetProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	id := ps.ByName("userID")
	fmt.Print(id)
	profile := profile.GetProfile( id )
	p, _ := json.Marshal(profile)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}

func GetProfiles( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	coordinates := ps.ByName("coordinates")
	radius := ps.ByName("radius")

	b := []byte(coordinates)
	var c location.Coordinates
	err := json.Unmarshal(b, &c)
	if(err!=nil) {
		panic(err)
	}
	rad, _ := strconv.ParseFloat(radius, 64)

	profile.GetProfiles( c, rad )
}

func CreateProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	// name := ps.ByName("name")
	// age, _ := strconv.Atoi(ps.ByName("age"))
	// bio := ps.ByName("bio")
	// imageURL := ps.ByName("imageURL")

	// newProfile := profileTypes.BlazrProfile {
	// 	Name: name,
	// 	Age: age,
	// 	Bio: bio,
	// 	ImageURL: imageURL,
	// }
	profile := profileTypes.BlazrProfile {}
	json.NewDecoder(r.Body).Decode(&profile)
	profile.CreateProfile( &profile )
	p := json.Marshal(profile)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", p)
}

func UpdateProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	// userID := ps.ByName("userID")
	// name := ps.ByName("name")
	// age, _ := strconv.Atoi(ps.ByName("age"))
	// bio := ps.ByName("bio")
	// imageURL := ps.ByName("imageURL")

	// toUpdate := profileTypes.BlazrProfile {
	// 	UserID: userID,
	// 	Name: name,
	// 	Age: age,
	// 	Bio: bio,
	// 	ImageURL: imageURL,
	// }
	p := profileTypes.BlazrProfile {}
	json.NewDecoder(r.Body).Decode(&p)
	profile.UpdateProfile( &toUpdate )
}

func DeleteProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	userID := ps.ByName("userID")

	profile.DeleteProfile( userID )
}