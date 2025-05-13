package main

import (
	"pizzaria/internal/data"
	"pizzaria/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()
	router := gin.Default()
	router.GET("/pizzas", handler.GetPizzas)
	router.POST("/pizzas", handler.PostPizzas)
	router.GET("/pizzas/:id", handler.GetPizzasById)
	router.DELETE("/pizzas/:id", handler.DeletePizzasById)
	router.PUT("/pizzas/:id", handler.UpdatePizzasById)
	router.POST("/pizzas/:id/reviews", handler.PostReview)
	router.Run()
}
