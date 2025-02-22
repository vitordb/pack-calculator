package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"pack-calculator/internal/application"
)

// Service is the CalculatorService used to perform the business logic of calculating pack sizes
var Service *application.CalculatorService

// CalculateHandler handles the /calculate route
// It processes the query parameters and performs the pack calculation logic
func CalculateHandler(c echo.Context) error {
	amountStr := c.QueryParam("amount")
	packSizesStr := c.QueryParam("packs")

	// Check if either 'amount' or 'packs' is missing in the request
	if amountStr == "" || packSizesStr == "" {
		return c.String(http.StatusBadRequest, "amount or packs not provided")
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid amount")
	}

	// Split the 'packs' parameter by commas and convert each pack size to an integer
	packsSplit := strings.Split(packSizesStr, ",")
	var packSizes []int
	for _, p := range packsSplit {
		// Trim any spaces around the pack size string
		p = strings.TrimSpace(p)
		// Convert the pack size to an integer
		val, err := strconv.Atoi(p)
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid pack size")
		}
		// Add the valid pack size to the slice of pack sizes
		packSizes = append(packSizes, val)
	}

	// Call the Service method to calculate the packs and save the result
	result, err := Service.CalculateAndSave(amount, packSizes)
	if err != nil {
		if err.Error() == "amount must be greater than zero" || err.Error() == "no pack sizes provided" ||
			err.Error() == "pack sizes must be greater than zero" {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Return the calculation result as a JSON response with the status OK
	return c.JSON(http.StatusOK, result.Solution)
}
