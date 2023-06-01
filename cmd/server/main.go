package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
	"tv/quick-bat/internal/usecase"
)

func main() {
	client := db.CreateConnection()
	defer db.CloseConnection()

	database := client.Database("quickbat")
	challengeRepo := db.NewChallengeRepository(database.Collection("challenges"))
	playerRepo := db.NewPlayerRepository(database.Collection("players"))

	domain.MainLeaderBoard = domain.NewLeaderBoard(playerRepo.FetchAll())

	usecase.LoadChallenge = challengeRepo.Find

	domain.AddChallengeCreateListener(challengeRepo.Add)
	domain.AddChallengeChangeListener(challengeRepo.Update)
	domain.AddPlayerChangeListener(playerRepo.Update)

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
