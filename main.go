package main

import (
	"fmt"
	"github.com/graphql-go/handler"
	"github.com/learn/config"
	db2 "github.com/learn/database"
	gql "github.com/learn/user"
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
	db = db.ConnectDB()
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
	http.Handle("/api/graphiql", disableCors(h))
	err := http.ListenAndServe(config.Config.ServeUri, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

