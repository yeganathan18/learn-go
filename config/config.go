package config

type configModel struct {
	MongoUri    string
	MongoDb     string
	TokenSecret string
	TokenExp    string
	ServeUri    string
}

var Config = configModel{
	MongoUri:    "mongodb://localhost:27017/animals?authSource=admin", // Mongo Uri example: database://admin:123456@localhost:37812/react-recipes
	MongoDb:     "animals",                                                     // DB name
	TokenSecret: "secret",                                                      // Secret to use in Tokens
	TokenExp:    "1h",                                                          // Expiration of Token
	ServeUri:    ":8080",                                                       // Serve
}

