package main

import (
	"pizzaria/models"

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
