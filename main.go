package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/josepnapitupulu/API_Jadwal_Minggu/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterJadwalsRoutes(r)
	http.Handle("/", r)

	fmt.Print("Starting Server localhost:7077")
	log.Fatal(http.ListenAndServe("localhost:7077", r))
}