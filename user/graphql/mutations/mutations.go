package mutations

import (
	"github.com/graphql-go/graphql"
	mutationfields "gitlab.com/lyra/backend/user/graphql/mutations/fields"
)

var UserMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"addUsername": mutationfields.AddUsername,
	},
})