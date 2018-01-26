package schema

import (
	profileTypes "../models/profile"
	profileResolvers "./resolvers/profile"

	"github.com/graphql-go/graphql"
	//"github.com/graphql-go/graphql/language/ast"
	//"github.com/graphql-go/graphql/language/parser"
)

var (
	Schema      graphql.Schema
	profileType *graphql.Object
	Kevin       profileTypes.BlazrProfile
)

func init() {
	profileType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Profile",
		Description: "A blazr user profile.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of a profile",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of a profile",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.Name, nil
					}
					return nil, nil
				},
			},
			"age": &graphql.Field{
				Type:        graphql.Int,
				Description: "The age of the user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.Age, nil
					}
					return nil, nil
				},
			},
			"bio": &graphql.Field{
				Type:        graphql.String,
				Description: "The bio of the user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.Bio, nil
					}
					return nil, nil
				},
			},
			"imageURL": &graphql.Field{
				Type:        graphql.String,
				Description: "A URL to a profile pic",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
						return profile.ImageURL, nil
					}
					return nil, nil
				},
			},
			// "matchPool": &graphql.Field {
			// 	Type: 		 graphql.NewList(profileType),
			// 	Description: "A freshly updated pool of new matches",
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
			// 			return profile.MatchPool, nil
			// 		}
			// 		return []interface{}{}, nil
			// 	},
			// },
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"profile": &graphql.Field{
				Type: profileType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the profile",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					return GetProfile(id), nil
				},
			},
			// "profiles": &graphql.Field {
			// 	Type: ,
			// 	Args: ,
			// 	Resolve: ,
			// }
		},
	})
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

func GetProfile(id string) profileTypes.BlazrProfile {
	return profileResolvers.GetProfile(id)
}

func GetProfiles() {

}
