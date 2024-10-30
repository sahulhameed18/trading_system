package service

import (
	"fmt"
	"sort"

	"github.com/sahul/trading_system/models"
)

// Initialize the order queue
var OrderQueue = []models.Order{}

// Initialize the order book
var OrderBook = map[string]map[string][]models.Order{}

// Function to append Orders in queue
func EnqueueOrder(order models.Order) {
	OrderQueue = append(OrderQueue, order)
	fmt.Println("New order is added in the OrderQueue")
}

// Function Remove the first order form the order queue
func DequeueOrder() models.Order {
	order := OrderQueue[0]
	OrderQueue = OrderQueue[1:]
	fmt.Println("0 th index order is removed from the OrderQueue")
	return order
}

// Function to check if the order queue is empty
func IsOrderQueueEmpty() bool {
	return len(OrderQueue) == 0
}

func addOrderToOrderBook(order models.Order, orderSide string) {
	symbol := order.Symbol
	// Check if the symbol exists
	existingOrder, isSymbolExist := OrderBook[symbol]

	if isSymbolExist {
		// Check the symbol has orderSide
		ordersList := existingOrder[orderSide]
		ordersList = append(ordersList, order)

		// sort the order list in descending order	
		sort.Slice(ordersList, func(i, j int) bool {
			return ordersList[i].Price > ordersList[j].Price
		})

		existingOrder[orderSide] = ordersList
		OrderBook[symbol] = existingOrder

	} else {
		// Creating new orderSide for  the symbol
		newOrderSide := map[string][]models.Order{}

		newOrderSide[orderSide] = []models.Order{order}
		OrderBook[symbol] = newOrderSide

	}

	fmt.Println("New Order is added in the order book")
}

// Function to route the order in the order book
func ProcessOrder(order models.Order) {
	if order.Side == "buy" {
		addOrderToOrderBook(order, "buyOrders")
	} else {
		addOrderToOrderBook(order, "sellOrders")
	}
}
