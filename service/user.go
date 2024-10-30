package service

import (
	"errors"
	"fmt"

	"github.com/sahul/trading_system/models"
)

// Retrive the user by userId
func getUserById(userId string) (models.User, error) {
	for _, user := range models.Users {
		if userId == user.Id {
			return user, nil
		}
	}

	userNotFound := errors.New("user not found")

	return models.User{}, userNotFound
}

// Check the user is valid
func isValidUser(userId string) bool {
	_, err := getUserById(userId)
	if err != nil {
		fmt.Println("Error while retrival of user:", err)
		return false
	}
	return true
}

// Balance Verification
func isSufficientBalance(userId string, orderPrice float64) bool {
	user, _ := getUserById(userId)

	return user.Balance >= orderPrice
}

// Margin Availability

// Get position size and position value for a user
func getNetPositionSizeAndValue(userId string) (int, float64) {
	user, _ := getUserById(userId)

	var totalPositionSize int = 0
	var totalPositionValue float64 = 0.0

	for _, pos := range user.OpenPosition {
		totalPositionSize += pos.Quantity
		totalPositionValue += float64(pos.Quantity) * pos.EntryPrice
	}

	return totalPositionSize, totalPositionValue
}

func calculateTotalCollateral(userId string) float64 {
	user, _ := getUserById(userId)

	totalCollateral := user.Balance

	for _, pos := range user.OpenPosition {
		// In the real world the position is calculated with the current market price
		totalCollateral += float64(pos.Quantity) * pos.EntryPrice
	}

	return totalCollateral
}

func getUserPositionIndex(userId string, symbol string) int {
	user, _ := getUserById(userId)

	for openPosIndex, openPos := range user.OpenPosition {
		if openPos.Symbol == symbol {
			return openPosIndex
		}
	}

	return -1
}

// Margin Computation
// Here by default for all the stocks the margin requirement is set to 30% in real world it changes for each stock that is decided the trading plaform

// Function to calculate the required margin
func calculateRequiredMargin(userId string, newOrderQty int, newOrderPrice float64) int {
	// Setting a default margin requirement for all stocks
	const marginReqPercent float64 = 0.30

	_, currentPositionValue := getNetPositionSizeAndValue(userId)

	totalPositionValue := currentPositionValue + float64(newOrderQty)*newOrderPrice // Changed this line

	requiredMargin := totalPositionValue * marginReqPercent

	return int(requiredMargin)

}

// Function to calculate the available margin
func calculateAvailableMargin(userId string, newOrdeQty int, newOrderPrice float64) float64 {
	totalColletral := calculateTotalCollateral(userId)

	requiredMargin := calculateRequiredMargin(userId, newOrdeQty, newOrderPrice)

	availableMargin := totalColletral - float64(requiredMargin)

	return availableMargin
}

// Function to check the sufficient margin
func checkSufficientMargin(userId string, newOrdeQty int, newOrderPrice float64) bool {
	requiredMargin := calculateRequiredMargin(userId, newOrdeQty, newOrderPrice)
	availableMargin := calculateAvailableMargin(userId, newOrdeQty, newOrderPrice)

	return availableMargin >= float64(requiredMargin)
}

// TODO: Handle the sell side if user does not have the symbol opened

// Function to check the trading permission
func hasTradingPermission(userId string, symbol string) bool {
	user, _ := getUserById(userId)

	for _, tradingSymbol := range user.TradingPermission {
		if tradingSymbol == symbol {
			return true
		}
	}

	return false
}

// Risk Management check

// Function to Check Position Limit
func isWithinPositionLimit(userId string, newOrderQty int) bool {
	user, _ := getUserById(userId)

	totoalPositionSize, _ := getNetPositionSizeAndValue(userId)

	return (totoalPositionSize + newOrderQty) <= user.PositionLimit
}

// Function to Check Exposure Limit
func isWithinExposureLimit(userId string, newOrdeQty int, newOrderPrice float64) bool {
	user, _ := getUserById(userId)

	_, totalPositionValue := getNetPositionSizeAndValue(userId)

	newOrderValue := newOrderPrice * float64(newOrdeQty)

	return (totalPositionValue + newOrderValue) <= float64(user.ExposureLimit)
}

// Function to Check the risk tolerence
func isWithinRiskTolerance(userId string, newOrderQty int, newOrderPirce float64) bool {
	user, _ := getUserById(userId)

	newOrderValue := newOrderPirce * float64(newOrderQty)

	switch user.RiskTolerance {
	case "low":
		return newOrderValue < 1000
	case "medium":
		return newOrderValue < 5000
	case "high":
		return newOrderValue < 10000
	default:
		return false
	}
}

// Function to Consolidate Risk Management Checks
func passesRiskManagement(userId string, newOrderQty int, newOrderPrice float64) (bool, string) {
	if !isWithinPositionLimit(userId, newOrderQty) {
		return false, "POSITION_LIMIT_EXCEEDED"
	}
	if !isWithinExposureLimit(userId, newOrderQty, newOrderPrice) {
		return false, "EXPOSURE_LIMIT_EXCEEDED"
	}
	if !isWithinRiskTolerance(userId, newOrderQty, newOrderPrice) {
		return false, "RISK_TOLERANCE_EXCEEDED"
	}

	return true, "Risk Management Validated"
}

// Function to check the account status
func CheckAccountStatus(newOrder models.Order) (bool, string) {

	if !isValidUser(newOrder.UserId) {
		return false, "INVALID_USER"
	}

	if !isSufficientBalance(newOrder.UserId, newOrder.Price) {
		return false, "INSUFFICIENT_BALANCE"
	}

	if !checkSufficientMargin(newOrder.UserId, newOrder.Quantity, newOrder.Price) {
		return false, "INSUFFICIENT_MARGIN"
	}

	if newOrder.Side == "sell" && getUserPositionIndex(newOrder.UserId, newOrder.Symbol) == -1 {
		return false, "NO_OPEN_POSITION_FOR_SYMBOL"
	}

	if !hasTradingPermission(newOrder.UserId, newOrder.Symbol) {
		return false, "TRADING_PERMISSION_DENIED"
	}

	riskManagementStatus, statusReason := passesRiskManagement(newOrder.UserId, newOrder.Quantity, newOrder.Price)
	if !riskManagementStatus {
		return false, statusReason
	}

	return true, "VALID_USER"
}
