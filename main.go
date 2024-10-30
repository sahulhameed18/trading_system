package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/sahul/trading_system/handlers"
	"github.com/sahul/trading_system/router"
)

func main() {
	fmt.Println("The Brokrage System")

	fmt.Println("Starting the server")

	// Inittalizing the wait groups 
	var wg sync.WaitGroup

	// 4 Routines to be waited 
	wg.Add(4)

	// Order Book Background worker
	go func() {
		defer wg.Done()
		handlers.OrderBookBackgroundWorker()
	}()

	// Sending the order to order book client 
	go func() {
		defer wg.Done()
		handlers.HandleSendOrderBook()
	}()

	// Order Matching background worker
	go func() {
		defer wg.Done()
		handlers.OrderMatchingBackgroundWorker()
	}()

	go func() {
		// Router 
		r := router.Router()

		if err := http.ListenAndServe(":7000", r); err != nil {
			fmt.Println("Failed to start server:", err)
		}
	}()

	wg.Wait()

}


