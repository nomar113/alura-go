package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main() {
	loadPizzas()
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasById)
	router.Run()
}

func getPizzas(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"pizzas": pizzas,
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
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	ctx.JSON(201, newPizza)
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

func loadPizzas() {
	file, err := os.Open("data/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func savePizza() {
	file, err := os.Create("data/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON: ", err)
	}
}
