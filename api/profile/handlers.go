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
	fmt.Println("Getting profile")
	fmt.Println(string(p))
	fmt.Print("\n")
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
	profiles := profileService.GetProfiles( coordinates, radius )
	p, _ := json.Marshal(profiles)
	fmt.Println("Fetching nearby profiles")
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)	
	fmt.Fprintf(w, "%s", p)
}

func CreateProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	profile := profileTypes.BlazrProfile {}
	json.NewDecoder(r.Body).Decode(&profile)
	created := profileService.CreateProfile( &profile )
	p, _ := json.Marshal(created)
	fmt.Println("Creating profile")
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}

func UpdateProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	profile := profileTypes.BlazrProfile {}
	json.NewDecoder(r.Body).Decode(&profile)
	updated := profileService.UpdateProfile( &profile )
	p, _ := json.Marshal(updated)
	fmt.Println("Updating profile")
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}

func DeleteProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	userID := ps.ByName("userID")

	deleted := profileService.DeleteProfile( userID )
	p, _ := json.Marshal(deleted)
	fmt.Println("Deleting profile")
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}