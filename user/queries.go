package user

import (
	"github.com/graphql-go/graphql"
)

var UserQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"getAllUsernames": &graphql.Field{
			Type:    graphql.NewList(UserType),
			Args:    graphql.FieldConfigArgument{},
			Resolve: GetAllUsernames,
		},
	},
})