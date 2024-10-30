package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sahul/trading_system/models"
	"github.com/sahul/trading_system/service"
)

var clients = make(map[*websocket.Conn]bool)
var Broadcasts = make(chan map[string]map[string][]models.Order)

var OrderAdded = make(chan bool)

// Upgrading the http connection to socket connection
var upgrade = websocket.Upgrader{
	// Allow socket connection for any origin
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler for incoming new order
func SubmitOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Initializing the submit order
	var submitOrder models.Order

	// Decoding the new order and sroting in submitOrder
	_ = json.NewDecoder(r.Body).Decode(&submitOrder)

	// Need to process the order
	// Check the order is valid
	isOrderValidated, orderValidatedReason := service.ValidateOrder(submitOrder)

	// Check the user is valid
	isUserValid, userValidatedReason := service.CheckAccountStatus(submitOrder)

	// Initializing the order response
	var orderResponse models.OrderResponse

	orderResponse.Id = submitOrder.Id

	// Invalid order or Invalid user response
	if !isOrderValidated {
		orderResponse.ResponseCode = orderValidatedReason
	} else if !isUserValid {
		orderResponse.ResponseCode = userValidatedReason
	}

	// Valid Order and Valid User response
	if isOrderValidated && isUserValid {
		orderResponse.ResponseCode = orderValidatedReason
		fmt.Println("Adding The New Order In OrderQueue")
		service.EnqueueOrder(submitOrder)
	}

	json.NewEncoder(w).Encode(orderResponse)

	fmt.Printf("The New Order Validation Status %t\n", isOrderValidated)
	fmt.Printf("The Status Reason %s\n", orderValidatedReason)
	fmt.Println("Submit Order", submitOrder)
}

// Handler function for background worker
func OrderBookBackgroundWorker() {
	for {
		if !service.IsOrderQueueEmpty() {
			// Removes the order from order queue(FIFO)
			order := service.DequeueOrder()

			// Process the removed order to the order book
			service.ProcessOrder(order)

			// Sending the orderBook to broadcasts
			Broadcasts <- service.OrderBook
			fmt.Println("order book is sent to broadcast")

			// Notify the order match
			OrderAdded <- true
			fmt.Println("Order match notification is sent")

			time.Sleep(100 * time.Millisecond)
		}

	}
}

// Handling the incoming socket connection
func HandleClientConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrading the http to socket connection
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err, "Error during updrading the connection")
	}
	defer ws.Close()

	// Adding the new web socket connection in ws channel
	clients[ws] = true
	fmt.Printf("New client socket is added %p", ws)
	fmt.Println("The clients", clients)

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err, "Error during reading  the  message from client")
			delete(clients, ws)
			break
		}
	}
}

// Function sends the order to client
func HandleSendOrderBook() {
	for {
		// Broadcasts send the orderBook to all connected clients
		// Receiving the orders from Broadcast channels
		orderMsg := <-Broadcasts
		fmt.Println("order book is received from broadcast", orderMsg)
		for client := range clients {
			// Writing the order message on client side
			err := client.WriteJSON(orderMsg)

			fmt.Println("Order is written to client")
			if err != nil {
				fmt.Println(err, "Error during writing the  message to clients")
				client.Close()
				delete(clients, client)
			}
		}

	}
}

// OrderMatching background worker
func OrderMatchingBackgroundWorker() {
	for {
		fmt.Println("Order Matching Background is working")
		// Receiving the orderadded notificaiton
		<-OrderAdded
		fmt.Println("Order match notification is received")

		// Invoking order match
		service.OrderMatch()

		// Sending the new order book to broadcasts channel
		Broadcasts <- service.OrderBook
		fmt.Println("Order match processed and broadcasted")

	}
}

func GetUserDetails(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Getting the user Details")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, user := range models.Users {
		if user.Id == params["userId"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	json.NewEncoder(w).Encode("No User Found")

}
