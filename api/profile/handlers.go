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
	err, profile := profileService.GetProfile( id )
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	p, _ := json.Marshal(profile)
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)	
    fmt.Fprintf(w, "%s", p)
}

func GetProfiles( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	q := r.URL.Query()
	userID := q.Get("userID")
	radius, _ := strconv.ParseFloat(q.Get("radius"), 64)
	lat, _ := strconv.ParseFloat(q.Get("lat"), 64)
	long, _ := strconv.ParseFloat(q.Get("long"), 64)
	coordinates := location.Coordinates {
		lat,
		long,
	}
	fmt.Println("Fetching nearby profiles")
	err, profiles := profileService.GetProfiles( coordinates, radius )
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err2, matches := matchpoolService.GetMatchedIds( userID )
	var filtered []profileTypes.BlazrProfile
	//would probably use a hashset for this yada yada yada
	if err2 != nil {
		filtered = profiles
	} else {
		for i := 0; i < len(profiles); i++ {
			if contains(matches, profiles[i].UserID) == false {
				filtered = append(filtered, profiles[i])
			}
		}
	}

	p, _ := json.Marshal(filtered)
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
	err, created := profileService.CreateProfile( &profile )
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
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
	err, updated := profileService.UpdateProfile( &profile )
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
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
	err, deleted := profileService.DeleteProfile( userID )
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
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
	err, matches := matchpoolService.GetMatchedProfiles( id )
	if err!=nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//fmt.Print(err)
	p, _ := json.Marshal( matches )
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)	
    fmt.Fprintf(w, "%s", p)
}

func UpdateLocation( w http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
	q := r.URL.Query()
	userID := q.Get("userID")
	lat, _ := strconv.ParseFloat(q.Get("lat"), 64)
	long, _ := strconv.ParseFloat(q.Get("long"), 64)
	fmt.Println("Updating location")
	err, updated := profileService.UpdateLocation( userID, lat, long )
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	p, _ := json.Marshal(updated)
	fmt.Println(string(p))
	fmt.Print("\n")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", p)
}

func contains( matches []string, id string ) bool {
	for _, m := range matches {
        if m == id {
            return true
        }
    }
    return false
}