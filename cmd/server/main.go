package main

import (
	"net/http"
	"strconv"
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
	"tv/quick-bat/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db.CreateConnection()
	defer db.CloseConnection()

	domain.GetLeaderBoard().Init(db.FetchAllPlayers())
	domain.OnPlayerChange = func(p *domain.Player) {
		db.UpdatePlayer(p)
	}
	
	router := gin.Default()

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
