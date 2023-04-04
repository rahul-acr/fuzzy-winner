package main

import (
	"net/http"
	"strconv"

	"tv/quick-bat/internal/domain"
	"tv/quick-bat/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	domain.GetLeaderBoard().Init([]*domain.Player{
		domain.CreatePlayer(1, 0, 0),
		domain.CreatePlayer(2, 0, 0),
	})

	router.GET("/players/:id", func(ctx *gin.Context) {
		playerId, _ := strconv.Atoi(ctx.Param("id"))
		playerDetails := usecase.GetPlayerDetails(playerId)
		ctx.JSON(http.StatusOK, &playerDetails)
	})

	router.POST("/matches", func(ctx *gin.Context) {
		var match usecase.Match
		ctx.BindJSON(&match)
		usecase.AddMatch(&match)
		ctx.Status(http.StatusCreated)
	})

	router.Run("localhost:8080")
}
