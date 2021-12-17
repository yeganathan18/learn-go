package fields

import (
	"github.com/graphql-go/graphql"
	qglt "gitlab.com/lyra/backend/user/graphql/types"
	"gitlab.com/lyra/backend/user/resolvers"
)

var GetAllUsernames = &graphql.Field{
		Type:    graphql.NewList(qglt.UserType),
		Args:    graphql.FieldConfigArgument{},
		Resolve: resolvers.GetAllUsernames,
}
