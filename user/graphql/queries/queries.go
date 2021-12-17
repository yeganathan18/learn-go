package queries

import (
	"github.com/graphql-go/graphql"
	queryfields "gitlab.com/lyra/backend/user/graphql/queries/fields"
)

var UserQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getAllUsernames": queryfields.GetAllUsernames,
	},
})