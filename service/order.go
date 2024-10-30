package service

import (
	"strings"
	"time"

	"github.com/sahul/trading_system/models"
)

// Global Symbols
var Symbols = []string{
	"AAPL",
	"GOOG",
	"MSFT",
	"AMZN",
	"TSLA",
	"FB",
	"NFLX",
	"NVDA",
	"INTC",
	"AMD",
}

// Initial Market Status
var MarketStatus = map[string]string{}

// Validation

// Check the symbol is valid
func isValidSymbol(orderSymbol string) bool {
	for _, symbol := range Symbols {
		if symbol == orderSymbol {
			return true
		}
	}
	return false
}

// Check the order is valid
func isValidOrderType(orderType string) bool {
	validOrderType := map[string]bool{
		"market":   true,
		"limit":    true,
		"stoploss": true,
	}

	return validOrderType[strings.ToLower(orderType)]
}

// Check the side is valid
func isValidSide(side string) bool {
	validSide := map[string]bool{
		"buy":  true,
		"sell": true,
	}

	return validSide[strings.ToLower(side)]
}

// Check the quantity is valid
func isValidQuantity(quantity int) bool {
	return quantity > 0
}

// Check the price is valid
func isValidPrice(price float64) bool {
	return price > 0 && price < 100
}

// Market status service

// Initialize market status
func initializeMarketStatus() {
	for _, symbol := range Symbols {
		MarketStatus[symbol] = "OPEN"
	}
}

// Check the symbol tradable
func isSymbolTradable(symbol string) bool {
	return MarketStatus[symbol] == "OPEN"
}

// Halt the trading for the given symbol
func HaltTrading(symbol string) {
	MarketStatus[symbol] = "HALTED"
}

// Resume the halted trading
func ResumeTrading(symbol string) {
	MarketStatus[symbol] = "OPEN"
}

// Check if the market is open
func isMarketOpen() bool {
	var marketStartHour = 6
	var marketStartMinute = 15
	var marketEndHour = 23
	var marketEndMinute = 30

	currentTime := time.Now()
	hour, minute, _ := currentTime.Clock()

	if hour < marketStartHour || hour > marketEndHour {
		return false
	}

	if hour == marketStartHour && minute < marketStartMinute {
		return false
	}

	if hour == marketEndHour && minute > marketEndMinute {
		return false
	}

	return true
}

// Function checks the market status 
func checkMarketStatus(symbol string) (bool, string) {
	// Initializing the market status
	initializeMarketStatus()

	if !isSymbolTradable(symbol) {
		return false, "SYMBOL_CLOSED"
	}

	if !isMarketOpen() {
		return false, "MARKET_CLOSED"
	}

	return true, "Market status verified"
}

// Function Validating an order 
func ValidateOrder(newOrder models.Order) (bool, string) {

	if !isValidSymbol(newOrder.Symbol) {
		return false, "INVALID_SYMBOL"
	}

	if !isValidOrderType(newOrder.Type) {
		return false, "INVALID_TYPE"
	}

	if !isValidSide(newOrder.Side) {
		return false, "INVALID_SIDE"
	}

	if !isValidQuantity(newOrder.Quantity) {
		return false, "INVALID_QUANTITY"
	}

	if !isValidPrice(newOrder.Price) {
		return false, "INVALID_PRICE"
	}

	marketStatus, statusReason := checkMarketStatus(newOrder.Symbol)
	if !marketStatus {
		return marketStatus, statusReason
	}

	return true, "Valid Order"
}
