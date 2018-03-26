package match

import (
	"strconv"
	"log"
	matchType "../../models/match"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {

	url = "localhost:27017"
	database = "blazr"
	collection = "matches"

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	c = session.DB(database).C(collection)

}

func FindOne( id string ) matchType.Match {

	// make query here
	var result matchType.Match
	err := c.Find( bson.M{ "matchID": id } ).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func FindByUserID( userA string, userB string ) matchType.Match {
	// look for the tuple forwards or reverse

	query := bson.M { "$or": [ { "$and" : ["userA" : userA }, { "userB" : userB }]}, "$and" : ["userA" : userB }, { "userB" : userA }]} ] }
	var user = matchType.Match{}
	err := c.Find( query ).One(&user)
	if err != nil {
		panic(err)
	}
	return user 
}

func Save( match *matchType.Match ) matchType.Match {
	
	err := c.Insert( match )

	if err != nil {
		panic(err)
	}
	return *match
}

func Remove( id string ) matchType.Match {
	toRemove := FindOne( id )
	err := c.Remove(bson.M{"matchID": id})
	if err != nil {
		panic(err)
	}
	return toRemove
 }

func Update( match *matchType.Match ) matchType.Match {
	change := bson.M { "$set": bson.M {"userA": match.UserA, "userB": match.UserB, "matched": match.Matched } }
	err := c.Update(bson.M { "matchID": profile.UserID }, change)
	if err != nil {
		panic(err)
	}
	return *match
}

func floatToString( number float64 ) string {
	return strconv.FormatFloat( number, 'f', 7, 64 )
}