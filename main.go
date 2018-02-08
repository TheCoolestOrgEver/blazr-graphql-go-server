package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"context"
	"log"

	"./api"

	"github.com/graphql-go/graphql"
)

func graphqlHandler( w http.ResponseWriter, r *http.Request ) {
	user := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{1, "cool user"}
	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: r.URL.Query().Get("query"),
		Context:       context.WithValue(context.Background(), "currentUser", user),
	})
	if len(result.Errors) > 0 {
		log.Printf("wrong result, unexpected errors: %v", result.Errors)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc( "/graphql", graphqlHandler)
	fmt.Println("Starting server")
	http.ListenAndServe(":8080", nil)
}