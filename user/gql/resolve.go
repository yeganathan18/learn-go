package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/learn/database"
	"github.com/learn/user/utils"
)

var db database.MongoDB

// Queries

func GetAllUsernames(_ graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	results, err = db.ConnectDB().GetAllUsernames()
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Mutations

func AddUsername(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var results interface{}

	username := p.Args["username"].(string)

	if err := utils.IsRequired(map[string]string{
		"username": username}); err != nil {
		return nil, err
	}
	results, err = db.ConnectDB().AddUsername(username)
	if err != nil {
		return nil, err
	}
	return results, nil
}
