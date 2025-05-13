package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidatePizzaPrice(pizza *models.Pizza) error {
	if pizza.Price < 0 {
		return errors.New("the price don't be less than zero")
	}
	return nil
}
