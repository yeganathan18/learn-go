package main

import (
	"fmt"
	"github.com/graphql-go/handler"
	"github.com/learn/config"
	db2 "github.com/learn/db"
	"github.com/learn/gql"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go lang backend is working Yahoooo!")
}

var db db2.MongoDB

func main() {
	fmt.Println("Hello World!")

	// MongoDB
	db = db2.ConnectDB()
	defer db.CloseDB()

	// GraphQL
	schema := gql.InitSchema()
	h := handler.New(&handler.Config{
		Schema:   &schema,
		GraphiQL: true, // GraphiQL interface
	})

	// Serve
	http.HandleFunc("/", helloWorld)
	//http.Handle("/", http.FileServer(http.Dir("./public"))) // Serve the frontend in /public
	http.Handle("/graphql", h)
	err := http.ListenAndServe(config.Config.ServeUri, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
