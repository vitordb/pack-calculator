package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"pack-calculator/internal/application"
)

var Service *application.CalculatorService

func CalculateHandler(c echo.Context) error {
	amountStr := c.QueryParam("amount")
	packSizesStr := c.QueryParam("packs")

	if amountStr == "" || packSizesStr == "" {
		return c.String(http.StatusBadRequest, "amount or packs not provided")
	}

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid amount")
	}

	packsSplit := strings.Split(packSizesStr, ",")
	var packSizes []int
	for _, p := range packsSplit {
		p = strings.TrimSpace(p)
		val, err := strconv.Atoi(p)
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid pack size")
		}
		packSizes = append(packSizes, val)
	}

	result, err := Service.CalculateAndSave(amount, packSizes)
	if err != nil {
		if err.Error() == "amount must be greater than zero" || err.Error() == "no pack sizes provided" ||
			err.Error() == "pack sizes must be greater than zero" {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result.Solution)
}
