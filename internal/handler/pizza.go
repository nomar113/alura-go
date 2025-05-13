package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPizzas(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"pizzas": data.Pizzas,
	})
}

func PostPizzas(ctx *gin.Context) {
	var newPizza models.Pizza
	if err := ctx.ShouldBindJSON(&newPizza); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	if err := service.ValidatePizzaPrice(&newPizza); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}
	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizza()
	ctx.JSON(http.StatusCreated, newPizza)
}

func GetPizzasById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for _, pizza := range data.Pizzas {
		if pizza.ID == id {
			ctx.JSON(http.StatusOK, pizza)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "Pizza not found",
	})
}

func DeletePizzasById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	for index, pizza := range data.Pizzas {
		if pizza.ID == id {
			data.Pizzas = append(data.Pizzas[:index], data.Pizzas[index+1:]...)
			data.SavePizza()
			ctx.JSON(http.StatusOK, gin.H{"message": "pizza deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}

func UpdatePizzasById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	var updatedPizza models.Pizza
	if err := ctx.ShouldBindBodyWithJSON(&updatedPizza); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	if err := service.ValidatePizzaPrice(&updatedPizza); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"erro": err.Error()})
		return
	}
	for index, pizza := range data.Pizzas {
		if pizza.ID == id {
			data.Pizzas[index] = updatedPizza
			data.Pizzas[index].ID = id
			data.SavePizza()
			ctx.JSON(http.StatusOK, data.Pizzas[index])
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "pizza not found"})
}
