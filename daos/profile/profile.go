package profile

import (
		"strconv"
		"log"
		profileTypes "../../models/profile"
		"../../models/location"
		"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
)

/*


  |		|	|
  |		|	| |
  |		|	| |	
________|________
		|		
  |	|	| |	
  |	|	| |	____
  |	|	| |


*/

var (
	Kevin profileTypes.BlazrProfile
	url string
	database string
	collection string
	c *mgo.Collection
)

func init() {

	url = "localhost:27017"
	database = "blazr"
	collection = "profiles"

	Kevin = profileTypes.BlazrProfile {
		UserID: "1", 
		Name: "Kevin", 
		Age: 22,
		Bio: "No butt stuff", 
		MatchPool: []profileTypes.BlazrProfile{},
	}

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	c = session.DB(database).C(collection)
}

func FindOne( id string ) profileTypes.BlazrProfile {

	// make query here
	var result profileTypes.BlazrProfile
	err := c.Find( bson.M{ "userID": id } ).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func FindAll( query string ) []profileTypes.BlazrProfile {
	return []profileTypes.BlazrProfile{ Kevin, Kevin }
 }

func FindByCoordinatesBetween( minCoordinates location.Coordinates, maxCoordinates location.Coordinates ) []profileTypes.BlazrProfile {
	// generate query
	query := createRadiusQuery( minCoordinates, maxCoordinates )
	var result []profileTypes.BlazrProfile
	err := c.Find(query).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}

 func Save( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	return Kevin
 }

func floatToString( number float64 ) string {
	return strconv.FormatFloat( number, 'f', 7, 64 )
}

func createRadiusQuery( minCoordinates location.Coordinates, maxCoordinates location.Coordinates ) bson.M {

	lat := "{ $gte:" + floatToString( minCoordinates.Lat ) + ", $lte: " + floatToString( maxCoordinates.Lat ) + "}"
	long := "{ $gte: " + floatToString( minCoordinates.Long ) + ", $lte: " + floatToString( maxCoordinates.Long ) + " }"

	return bson.M { "coordinates.lat": lat, "coordinates.long": long }
}