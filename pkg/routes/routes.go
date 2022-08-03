package routes

import (
	"github.com/gorilla/mux"
	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/orders", controllers.AddOrder).Methods("POST")
	router.HandleFunc("/users/orders", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/users/friends", controllers.GetFriends).Methods("GET")
	router.HandleFunc("/users/friends", controllers.AddFriend).Methods("POST")
}
