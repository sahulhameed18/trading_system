package service

import (
	"fmt"

	"github.com/sahul/trading_system/models"
)

// Finding the mathching order for given buyOrders and selLOrders
func findOrderMatching(buyOrders []models.Order, sellOrders []models.Order, symbol string) {
	fmt.Println("Finding the matching orders")

	fmt.Println("Buy orders", buyOrders, "Sell Orders", sellOrders)

	//FIXED Solution for channel blocking is infinite loop without condition: Error observed is due to inifine while loop withouy any break condition Now added if condition if while loop and use len() not to keep in seperate variable

	// While loop for matching order and executing the trade
	for len(buyOrders) > 0 && len(sellOrders) > 0 {

		totalSellOrders := len(sellOrders)

		// Matching order condition
		if buyOrders[0].Price >= sellOrders[totalSellOrders-1].Price {
			bestBuyOrder := buyOrders[0]
			bestSellOrder := sellOrders[totalSellOrders-1]

			fmt.Println("Matching order found", bestBuyOrder, bestSellOrder)

			// Function to execute the trade
			buyState, remainingBuyOrderQuantity, sellState, remainingSellOrderQuantity := ExecuteTrade(bestBuyOrder, bestSellOrder)

			// Handling the buy orders after trade execution
			if buyState == "FILLED" {
				buyOrders = buyOrders[1:]
			} else {
				buyOrders[0].Quantity = remainingBuyOrderQuantity
			}

			// Handling the sell orders after trade execution
			if sellState == "FILLED" {
				sellOrders = sellOrders[:totalSellOrders-1]
			} else {
				sellOrders[totalSellOrders-1].Quantity = remainingSellOrderQuantity
			}

			fmt.Println("Updating the order book after trading")
			// Updating the orderBook
			OrderBook[symbol]["buyOrders"] = buyOrders
			OrderBook[symbol]["sellOrders"] = sellOrders

			fmt.Println("Updated order book", OrderBook)

			// Update the user's balance and user's open position
			updateUserBalacnceAndOpenPosition(Trades[len(Trades)-1])

		} else {
			break
		}

	}
}

func OrderMatch() {
	fmt.Println("Matching the order starts")

	// Loop through all the symbols buy and sell orders for finding the order match
	for symbol, orders := range OrderBook {

		buyOrders := orders["buyOrders"]
		sellOrders := orders["sellOrders"]

		// Finding the matching orders
		fmt.Println("Buy orders", buyOrders, "Sell Orders", sellOrders)
		findOrderMatching(buyOrders, sellOrders, symbol)
	}

}

func updateUserBalacnceAndOpenPosition(trade models.Trade) {
	fmt.Println("Updating the user's balance after trade")

	fmt.Println(trade.BuyOrderId, trade.SellOrderId)

	buyUser, _ := getUserById(trade.BuyOrderId)
	sellUser, _ := getUserById(trade.SellOrderId)

	for userIndex, user := range models.Users {
		if (user.Id) == buyUser.Id {
			// updating the user's balance for buy side
			models.Users[userIndex].Balance -= trade.Price * float64(trade.Quantity)

			// updating the user's positions for buy side
			openPosIndex := getUserPositionIndex(user.Id, trade.Symbol)
			if openPosIndex != -1 {
				models.Users[userIndex].OpenPosition[openPosIndex].Quantity += trade.Quantity
				models.Users[userIndex].OpenPosition[openPosIndex].EntryPrice = trade.Price
			} else {
				models.Users[userIndex].OpenPosition = append(models.Users[userIndex].OpenPosition,
					models.OpenPosition{Symbol: trade.Symbol,
						Quantity:   trade.Quantity,
						EntryPrice: trade.Price})
			}

		}

		if (user.Id) == sellUser.Id {
			// updating the user's balance for sell side
			models.Users[userIndex].Balance += trade.Price * float64(trade.Quantity)

			// updating the user's positions for sell side
			openPosIndex := getUserPositionIndex(user.Id, trade.Symbol)
			if openPosIndex != -1 {
				models.Users[userIndex].OpenPosition[openPosIndex].Quantity -= trade.Quantity
				models.Users[userIndex].OpenPosition[openPosIndex].EntryPrice = trade.Price
			}

		}
	}

}
