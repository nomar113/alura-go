package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

func getPizzas(ctx *gin.Context) {
	var pizzas = []models.Pizza{
		{ID: 1, Name: "Toscana", Price: 49.5},
		{ID: 2, Name: "Marguerita", Price: 79.5},
		{ID: 3, Name: "Frango com requeijão", Price: 69.5},
	}
	ctx.JSON(200, gin.H{
		"message": pizzas,
	})
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.Run()
}
