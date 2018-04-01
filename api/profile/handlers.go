package handlers

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
	"strconv"
	"fmt"
	"encoding/json"
	profileService "../../services/profile"
	"../../models/location"
	profileTypes "../../models/profile"
)

func GetProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	id := ps.ByName("userID")
	fmt.Print(id)
	profile := profileService.GetProfile( id )
	p, _ := json.Marshal(profile)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)	
    fmt.Fprintf(w, "%s", p)
}

func GetProfiles( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	q := r.URL.Query()
	radius, _ := strconv.ParseFloat(q.Get("radius"), 64)
	lat, _ := strconv.ParseFloat(q.Get("lat"), 64)
	long, _ := strconv.ParseFloat(q.Get("long"), 64)
	
	coordinates := location.Coordinates {
		lat,
		long,
	}
	fmt.Print(coordinates.Lat)
	fmt.Print(coordinates.Long)
	fmt.Print(radius)
	profileService.GetProfiles( coordinates, radius )
}

func CreateProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	profile := profileTypes.BlazrProfile {}
	json.NewDecoder(r.Body).Decode(&profile)
	created := profileService.CreateProfile( &profile )
	p, _ := json.Marshal(created)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}

func UpdateProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	profile := profileTypes.BlazrProfile {}
	json.NewDecoder(r.Body).Decode(&profile)
	updated := profileService.UpdateProfile( &profile )
	p, _ := json.Marshal(updated)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}

func DeleteProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	userID := ps.ByName("userID")

	deleted := profileService.DeleteProfile( userID )
	p, _ := json.Marshal(deleted)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}