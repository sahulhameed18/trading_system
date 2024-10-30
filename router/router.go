package router

import (
	"github.com/gorilla/mux"
	"github.com/sahul/trading_system/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Route Handling the incoming new order
	router.HandleFunc("/api/order", handlers.SubmitOrder).Methods("POST")

	// Route handling the user details request
	router.HandleFunc("/user/{userId}", handlers.GetUserDetails).Methods("GET")

	// Route Handling the web socket connection
	router.HandleFunc("/ws/orderbook", handlers.HandleClientConnection)
	return router
}
