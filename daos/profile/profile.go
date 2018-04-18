package profile

import (
		"strconv"
		//"log"
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
		panic(err)
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
	//fmt.Println(query)
	if err != nil {
		panic(err)
	}
	return result
}

 func Save( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	
	err := c.Insert( profile )

	if err != nil {
		panic(err)
	}

	return *profile
 }

func Remove( id string ) profileTypes.BlazrProfile {
	toRemove := FindOne( id )
	err := c.Remove(bson.M{"userID": id})
	if err != nil {
		panic(err)
	}
	return toRemove
 }

func Update( profile *profileTypes.BlazrProfile ) profileTypes.BlazrProfile {
	change := bson.M { "$set": bson.M {"name": profile.Name, "age": profile.Age, "bio": profile.Bio, "imageURL": profile.ImageURL, "location": profile.Location } }
	err := c.Update(bson.M { "userID": profile.UserID }, change)
	if err != nil {
		panic(err)
	}
	return *profile
}

func floatToString( number float64 ) string {
	return strconv.FormatFloat( number, 'f', 7, 64 )
}

func createRadiusQuery( minCoordinates location.Coordinates, maxCoordinates location.Coordinates ) bson.M {
	query := bson.M { "$and": []interface{} {
		bson.M {
			"location.lat": bson.M {
				"$gte": minCoordinates.Lat, "$lte": maxCoordinates.Lat,
			}, 
		},
		bson.M {
			"location.long": bson.M {
				"$gte": minCoordinates.Long, "$lte": maxCoordinates.Long,
			},
		},
	},
	}

	return query
}