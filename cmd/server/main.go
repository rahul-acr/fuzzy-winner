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
	playerRepo := db.NewPlayerRepository(database.Collection("players"))
	challengeRepo := db.NewChallengeRepository(database.Collection("challenges"), playerRepo)

	domain.MainLeaderBoard = domain.NewLeaderBoard(playerRepo.FetchAll())

	usecase.LoadChallenge = challengeRepo.Find

	domain.AddChallengeCreateListener(challengeRepo.Add)
	domain.AddChallengeChangeListener(challengeRepo.Update)
	domain.AddPlayerChangeListener(playerRepo.Update)

	router := gin.Default()

	router.GET("/players/:id", func(ctx *gin.Context) {
		playerId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		playerDetails := usecase.GetPlayerDetails(playerId)
		ctx.JSON(http.StatusOK, &playerDetails)
	})

	router.POST("/matches", func(ctx *gin.Context) {
		var match usecase.Match
		if err := ctx.BindJSON(&match); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		usecase.AddMatch(&match)
		ctx.Status(http.StatusCreated)
	})

	router.POST("/challenges", func(ctx *gin.Context) {
		var challenge usecase.Challenge
		if err := ctx.BindJSON(&challenge); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		usecase.CreateChallenge(challenge)
	})

	router.POST("/challenges/:id/accept", func(ctx *gin.Context) {
		challengeId := ctx.Param("id")
		var challengeAccept usecase.ChallengeAccept
		if err := ctx.BindJSON(&challengeAccept); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		if err := usecase.AcceptChallenge(challengeId, challengeAccept); err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	})

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
