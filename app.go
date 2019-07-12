package main

import (
	. "github.com/davdwhyte87/smh-api/dao"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

import routers "github.com/davdwhyte87/smh-api/routers"

var dao = DatabaseDAO{}


// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	Server, _ := os.LookupEnv("SERVER")
	database, _ := os.LookupEnv("DATABASE")
	dao.Server = Server
	dao.Database = database
	dao.Connect()
}


func main() {
	r := mux.NewRouter()
	v1Router := r.PathPrefix("/api/v1").Subrouter()

	// v1Router.HandleFunc("/user/signup", CreateUser).Methods("POST")
	routers.GetUserRoutes(v1Router)
	
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}