package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bvalkanov/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRouters(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:9010", r))

	fmt.Println("Listening on port 9010")
}
