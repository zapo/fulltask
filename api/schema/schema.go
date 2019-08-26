package schema

import (
	"github.com/graphql-go/graphql"
	"log"
)

var Schema graphql.Schema

func init() {
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: RootMutation,
	})

	if err != nil {
		log.Fatal("Couldnt create the schema", err)
	}
}
