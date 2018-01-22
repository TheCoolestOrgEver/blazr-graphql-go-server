package schema

import (
	profileTypes "../models/profile"
	profileResolvers "./resolvers/profile"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
)

var (
	Schema 		graphql.Schema
	profileType *graphql.Object 
)

func init() {
	profileType = graphql.NewObject( graphql.ObjectConfig { 
		Name: "Profile",
		Description: "A blazr user profile.",
		Fields: graphql.Fields {
			"name": &graphql.Field {
				Type: graphql.String,
				Description: "The name of a profile", 
				Resolve: getProfileName,
			},
			"age": &graphql.Field {
				Type: graphql.Int, 
				Description: "The age of the user",
				Resolve: getProfileAge,
			},
			"bio": &graphql.Field {
				Type: graphql.String, 
				Description: "The bio of the user",
				Resolve: getProfileBio,
			}, 
			"imageURL": &graphql.Field {
				Type: graphql.String, 
				Description: "A URL to a profile pic",
				Resolve: getProfileImageURL,
			}, 
			"matchPool": &graphql.Field {
				Type: graphql.NewList(profileType), 
				Description: "A freshly updated pool of new matches",
				Resolve: getProfileMatchPool,
			},
		},
	} )

	// queryType := graphql.NewObject( graphql.ObjectConfig {
	// 	Name: "Query",
	// 	Fields: graphql.Fields {
	// 		"profile": &graphql.Field {
	// 			Type: ,
	// 			Args: ,
	// 			Resolve: ,
	// 		}, 
	// 		"profiles": &graphql.Field {
	// 			Type: ,
	// 			Args: ,
	// 			Resolve: ,
	// 		}
	// 	}
	// } )
}

func getProfileName( p graphql.ResolveParams ) graphql.String {

}

func getProfileAge( p graphql.ResolveParams ) graphql.Int {

}

func getProfileBio( p graphql.ResolveParams ) graphql.String {

}

func getProfileImageURL( p graphql.ResolveParams ) graphql.String {

}

func getProfileMatchPool( p graphql.ResolveParams ) *graphql.List {

}

