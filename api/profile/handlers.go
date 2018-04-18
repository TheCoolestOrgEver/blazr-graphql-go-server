package handlers

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
	"strconv"
	"fmt"
	//"log"
	"encoding/json"
	profileService "../../services/profile"
	matchpoolService "../../services/matchpool"
	"../../models/location"
	profileTypes "../../models/profile"
)

func GetProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	id := ps.ByName("userID")
	fmt.Print(id)
	fmt.Println("Getting profile")
	profile := profileService.GetProfile( id )
	p, _ := json.Marshal(profile)
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
	fmt.Println("Fetching nearby profiles")
	profiles := profileService.GetProfiles( coordinates, radius )
	p, _ := json.Marshal(profiles)
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)	
	fmt.Fprintf(w, "%s", p)
}

func CreateProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	profile := profileTypes.BlazrProfile {}
	json.NewDecoder(r.Body).Decode(&profile)
	fmt.Println("Creating profile")
	created := profileService.CreateProfile( &profile )
	p, _ := json.Marshal(created)
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}

func UpdateProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	profile := profileTypes.BlazrProfile {}
	json.NewDecoder(r.Body).Decode(&profile)
	fmt.Println("Updating profile")
	updated := profileService.UpdateProfile( &profile )
	p, _ := json.Marshal(updated)
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}

func DeleteProfile( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	userID := ps.ByName("userID")
	fmt.Println("Deleting profile")
	deleted := profileService.DeleteProfile( userID )
	p, _ := json.Marshal(deleted)
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}

func GetMatches( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	id := ps.ByName("userID")
	fmt.Print(id)
	fmt.Println("Getting matches")
	err, matches := matchpoolService.GetMatches( id )
	if err!=nil {
		panic(err)
	}
	//fmt.Print(err)
	p, _ := json.Marshal( matches )
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)	
    fmt.Fprintf(w, "%s", p)
}