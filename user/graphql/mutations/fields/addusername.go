package fields

import (
	"github.com/graphql-go/graphql"
	qglt "gitlab.com/lyra/backend/user/graphql/types"
	"gitlab.com/lyra/backend/user/resolvers"
)

var AddUsername = &graphql.Field{
	Type: qglt.UserType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.AddUsername,
}
