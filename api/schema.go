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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.Name, nil
					}
					return nil, nil
				},
			},
			"age": &graphql.Field {
				Type: graphql.Int, 
				Description: "The age of the user",
				Resolve: Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.Age, nil
					}
					return nil, nil
				},
			},
			"bio": &graphql.Field {
				Type: graphql.String, 
				Description: "The bio of the user",
				Resolve: Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.Bio, nil
					}
					return nil, nil
				},
			}, 
			"imageURL": &graphql.Field {
				Type: graphql.String, 
				Description: "A URL to a profile pic",
				Resolve: Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.ImageURL, nil
					}
					return nil, nil
				},
			}, 
			"matchPool": &graphql.Field {
				Type: graphql.NewList(profileType), 
				Description: "A freshly updated pool of new matches",
				Resolve: Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.MatchPool, nil
					}
					return []interface{}{}, nil
				},
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

func GetProfile(){

}

func GetProfiles(){
	
}

