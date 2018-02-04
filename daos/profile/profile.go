package profile

import (
		"strconv"
		profileTypes "../../models/profile"
		"../../models/location"
		"gopkg.in/mgo.v2"
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

	url = "google.com"
	database = "database"
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
		panic(err)
	}
	c = session.DB(database).C(collection)
}

func FindOne( id string ) profileTypes.BlazrProfile {

	// make query here
	query := "query string"
	var result profileTypes.BlazrProfile
	err := c.Find(query).One(&result)
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
	query := createRadiusQueryString( minCoordinates, maxCoordinates )
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

func createRadiusQueryString( minCoordinates location.Coordinates, maxCoordinates location.Coordinates ) string {
	return "{ coordinates.lat: { $gte: " + floatToString( minCoordinates.Lat ) + ", $lte: " + floatToString( maxCoordinates.Lat ) + 
		   " }, coordinates.long: { $gte: " + floatToString( minCoordinates.Long ) + ", $lte: " + floatToString( maxCoordinates.Long ) + " }"
}