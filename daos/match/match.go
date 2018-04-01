package match

import (
	"strconv"
	"log"
	matchType "../../models/match"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	url string
	database string
	collection string
	c *mgo.Collection
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

func FindOne( id string ) ( error, matchType.Match ) {

	// make query here
	var result matchType.Match
	err := c.Find( bson.M{ "matchID": id } ).One(&result)
	// if err != nil {
	// 	panic(err)
	// }
	return err, result
}

func FindByUserID( userA string, userB string ) ( error, matchType.Match ) {
	// look for the tuple forwards or reverse

	query := bson.M { "$or": []interface{} {
		bson.M {"$and" : []interface{} { 
			bson.M {"userA" : userA },
			bson.M {"userB" : userB }}},
		bson.M {"$and" : []interface{} {
			bson.M {"userA" : userB }, 
			bson.M {"userB" : userA }}},
		},
	}
	var user = matchType.Match{}
	err := c.Find( query ).One(&user)
	// if err != nil {
	// 	panic(err)
	// }
	return err, user 
}

func Save( match *matchType.Match ) ( error, matchType.Match ) {
	
	err := c.Insert( match )

	// if err != nil {
	// 	panic(err)
	// }
	return err, *match
}

func Remove( id string ) ( error, matchType.Match ) {
	_, toRemove := FindOne( id )
	err := c.Remove(bson.M{"matchID": id})
	// if err != nil {
	// 	panic(err)
	// }
	return err, toRemove
 }

func Update( match *matchType.Match ) ( error, matchType.Match ) {
	change := bson.M { "$set": bson.M {"userA": match.UserA, "userB": match.UserB, "matched": match.Matched } }
	err := c.Update(bson.M { "matchID": match.MatchID }, change)
	// if err != nil {
	// 	panic(err)
	// }
	return err, *match
}

func floatToString( number float64 ) string {
	return strconv.FormatFloat( number, 'f', 7, 64 )
}