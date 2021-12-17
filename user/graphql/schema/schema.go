package schema

import (
	"gitlab.com/lyra/backend/user/graphql/mutations"
	"gitlab.com/lyra/backend/user/graphql/queries"
	qglt "gitlab.com/lyra/backend/user/graphql/types"
	"log"

	"github.com/graphql-go/graphql"
)

func InitSchema() graphql.Schema {
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queries.UserQueries,
		Mutation: mutations.UserMutations,
		Types: []graphql.Type{qglt.ID},
	})
	if err != nil {
		log.Fatal(err)
	}
	return graphqlSchema

}

