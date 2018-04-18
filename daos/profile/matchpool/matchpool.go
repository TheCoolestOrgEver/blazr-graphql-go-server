package matchpool

import (
		//"log"
		"../../../models/profile/matchpool"
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
	collection = "matchpools"

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	c = session.DB(database).C(collection)
}

func FindOne( id string ) (error, matchpool.MatchPool) {

	// make query here
	var result matchpool.MatchPool
	err := c.Find( bson.M{ "userID": id } ).One(&result)

	return err, result
}

func Save( match *matchpool.MatchPool ) (error, matchpool.MatchPool) {
	
	err := c.Insert( match )

	return err, *match
 }

func Remove( id string ) (error, matchpool.MatchPool) {
	err, toRemove := FindOne( id )
	if err != nil {
		panic(err)
	}
	err = c.Remove(bson.M{"userID": id})

	return err, toRemove
 }

func Update( matchpool *matchpool.MatchPool ) (error, matchpool.MatchPool) {
	change := bson.M { "$set": bson.M { "matches": matchpool.Matches } }
	err := c.Update(bson.M { "userID": matchpool.UserID }, change)

	return err, *matchpool
}