package service

import (
	"fmt"

	"github.com/sahul/trading_system/models"
)

var Trades []models.Trade

// Function executing the trade
func ExecuteTrade(buyOrder models.Order, sellOrder models.Order) (string, int, string, int) {
	fmt.Println("Executing trade begins", buyOrder, sellOrder)

	// trade quantity is lowest quatity either from buy or sell side
	tradeQuantity := min(buyOrder.Quantity, sellOrder.Quantity)

	fmt.Println("The trade quantity", tradeQuantity)

	// trade price is set to ask price(seller asked price)
	tradePrice := sellOrder.Price

	// Creating a new trade(Executing the trade)
	newTrade := models.Trade{
		BuyOrderId:  buyOrder.UserId,
		SellOrderId: sellOrder.UserId,
		Symbol:      buyOrder.Symbol,
		Quantity:    tradeQuantity,
		Price:       tradePrice,
	}

	fmt.Println("The new executed trade", newTrade)

	// Remaining Quantity
	remainingBuyOrderQuantity := buyOrder.Quantity - tradeQuantity
	remainingSellOrderQuantity := sellOrder.Quantity - tradeQuantity

	var buyState string
	var sellState string

	// Buy State
	if remainingBuyOrderQuantity == 0 {
		buyState = "FILLED"
	} else {
		buyState = "PARTIALLY_FILLED"
	}

	// Sell State
	if remainingSellOrderQuantity == 0 {
		sellState = "FILLED"
	} else {
		sellState = "PARTIALLY_FILLED"
	}

	// Logging the trade
	Trades = append(Trades, newTrade)
	fmt.Println("The trades", Trades)

	return buyState, remainingBuyOrderQuantity, sellState, remainingSellOrderQuantity
}
