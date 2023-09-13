package main

import (
	"encoding/json"
	"fmt"
	"go_graphql/resolvers"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    resolvers.QueryType,
		Mutation: resolvers.MutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
