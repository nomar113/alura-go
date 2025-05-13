package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReview(ctx *gin.Context) {
	pizzaId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	var newReview models.Review
	if err := ctx.ShouldBindBodyWithJSON(&newReview); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	if err := service.ValidateReviewRating(&newReview); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	for index, pizza := range data.Pizzas {
		if pizza.ID == pizzaId {
			pizza.Review = append(pizza.Review, newReview)
			data.Pizzas[index] = pizza
			data.SavePizza()
			ctx.JSON(http.StatusCreated, gin.H{"message": pizza})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"erro": "pizza not found"})
}
