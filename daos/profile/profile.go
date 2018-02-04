package profile

import (
		"gopkg.in/mgo.v2"
		profileTypes "../../models/profile"
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

 func InsertOne( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	return Kevin
 }

 func FindAll( query string ) []profileTypes.BlazrProfile {
	return []profileTypes.BlazrProfile{ Kevin, Kevin }
 }