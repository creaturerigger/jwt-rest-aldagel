package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/routes"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
