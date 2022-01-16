package user

import (
	"log"

	"github.com/graphql-go/graphql"
)

func InitSchema() graphql.Schema {
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    UserQueries,
		Mutation: UserMutations,
		Types:    []graphql.Type{ID},
	})
	if err != nil {
		log.Fatal(err)
	}
	return graphqlSchema

}

