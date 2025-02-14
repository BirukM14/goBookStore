package main

import (
	"log"
	"net/http"

	"github.com.gorilla/mux"
	"github.com/BirukM14/goBookStore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm.dialects/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("r",r)
	log.Fatal(http.ListenAndServe("Localhost:9010", r)) 
}