package schema

import (
	profileTypes "../models/profile"
	profileResolvers "./resolvers/profile"

	"strconv"

	"github.com/graphql-go/graphql"
	//"github.com/graphql-go/graphql/language/ast"
	//"github.com/graphql-go/graphql/language/parser"
)

var profileType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Profile",
	Description: "A blazr user profile.",
	Fields: graphql.Fields{
		"userID": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The id of a profile",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if profile, ok := p.Source.(profileTypes.BlazrProfile); ok {
					return profile.UserID, nil
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
		// 		if profile, ok := p.Source.(profileTypes.BlazrProfile); ok && len(profile.MatchPool) > 0 {
		// 			return profile.MatchPool, nil
		// 		}
		// 		return []interface{}{}, nil
		// 	},
		// },
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"profile": &graphql.Field{
			Type: profileType,
			Args: graphql.FieldConfigArgument{
				"userID": &graphql.ArgumentConfig{
					Description: "id of the profile",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["userID"].(string)
				return GetProfile(id), nil
			},
		},
		"profiles": &graphql.Field {
			Type: []profileType,
			Args: graphql.FieldConfigArgument{
				"geolocation": &graphql.ArgumentConfig{
					Description: "location of the user to query nearby profiles",
					Type:		 graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// parse JSON here 
				//return GetProfiles()
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation", 
	Fields: graphql.Fields {
		"createProfile": &graphql.Field {
			Type: profileType, 
			Description: "Create new profile", 
			Args: graphql.FieldConfigArgument {
				"name": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
				}, 
				"age": &graphql.ArgumentConfig {
					Type: graphql.NewNonNull(graphql.String),
				}, 
				"bio": &graphql.ArgumentConfig {
					Type: graphql.String,
				},
				"imageURL": &graphql.ArgumentConfig {
					Type: graphql.String, 
				}, 
			}, 
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name, _ := p.Args["name"].(string)
				age, _ := strconv.Atoi(p.Args["age"].(string))
				bio, _ := p.Args["bio"].(string)
				imageURL, _ := p.Args["imageURL"].(string)

				newProfile := profileTypes.BlazrProfile {
					Name: name,
					Age: age,
					Bio: bio,
					ImageURL: imageURL,
				}
				
				return createProfile(&newProfile), nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})


func GetProfile(id string) profileTypes.BlazrProfile {
	return profileResolvers.GetProfile(id)
}

func GetProfiles(coordinates profileTypes.Coordinates) []profileTypes.BlazrProfile {
	return profileResolvers.GetProfiles(coordinates)
}

func createProfile(profile *profileTypes.BlazrProfile) profileTypes.BlazrProfile {
	return profileResolvers.CreateProfile(profile)
}
