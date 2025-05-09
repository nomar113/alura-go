package main

import (
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Name: "Toscana", Price: 49.5},
	{ID: 2, Name: "Marguerita", Price: 79.5},
	{ID: 3, Name: "Frango com requeijão", Price: 69.5},
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasById)
	router.Run()
}

func getPizzas(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": pizzas,
	})
}

func postPizzas(ctx *gin.Context) {
	var newPizza models.Pizza
	if err := ctx.ShouldBindJSON(&newPizza); err != nil {
		ctx.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	pizzas = append(pizzas, newPizza)
}

func getPizzasById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	for _, pizza := range pizzas {
		if pizza.ID == id {
			ctx.JSON(200, pizza)
			return
		}
	}
	ctx.JSON(404, gin.H{
		"message": "Pizza not found",
	})
}
