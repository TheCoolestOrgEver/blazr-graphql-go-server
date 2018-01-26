package profile

import (
		//import "gopkg.in/mgo.v2"
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
)

func init() {
	Kevin = profileTypes.BlazrProfile {
		ID: "1", 
		Name: "Kevin", 
		Age: 22,
		Bio: "No butt stuff", 
	}
}

func FindOne(id string) profileTypes.BlazrProfile {
	return Kevin
}