package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"./api"

	"github.com/graphql-go/graphql"
)

func main() {
	http.HandleFunc( "/graphql",  func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params {
			Schema: schema.Schema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result);
	})
	fmt.Println("Starting server")
	http.ListenAndServe(":8080", nil)
}