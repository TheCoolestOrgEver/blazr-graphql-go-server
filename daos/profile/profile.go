package profile

import (
		"strconv"
		//"log"
		profileTypes "../../models/profile"
		"../../models/location"
		"gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
		"errors"
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

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	c = session.DB(database).C(collection)
}

func FindOne( id string ) (error, profileTypes.BlazrProfile) {

	// make query here
	var result profileTypes.BlazrProfile
	err := c.Find( bson.M{ "userID": id } ).One(&result)

	return err, result
}

func FindByCoordinatesBetween( minCoordinates location.Coordinates, maxCoordinates location.Coordinates ) (error, []profileTypes.BlazrProfile) {
	// generate query
	query := createRadiusQuery( minCoordinates, maxCoordinates )
	var result []profileTypes.BlazrProfile
	err := c.Find(query).All(&result)
	//fmt.Println(query)
	if len(result) == 0 {
		err = errors.New("Not found")
	}
	return err, result
}

 func Save( profile *profileTypes.BlazrProfile ) (error, profileTypes.BlazrProfile) {
	
	err := c.Insert( profile )

	return err, *profile
 }

func Remove( id string ) (error, profileTypes.BlazrProfile) {
	err, toRemove := FindOne( id )
	err = c.Remove(bson.M{"userID": id})

	return err, toRemove
 }

func Update( profile *profileTypes.BlazrProfile ) (error, profileTypes.BlazrProfile) {
	change := bson.M { "$set": bson.M {"name": profile.Name, "age": profile.Age, "bio": profile.Bio, "imageURL": profile.ImageURL, "location": profile.Location } }
	err := c.Update(bson.M { "userID": profile.UserID }, change)

	return err, *profile
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

func FetchMatches(ids []string) (error, []profileTypes.BlazrProfile) {
	var result []profileTypes.BlazrProfile
	err := c.Find(bson.M{"userID": bson.M{"$in": ids}}).All(&result);
	return err, result
}