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
	client := db.CreateConnection()
	defer db.CloseConnection()

	database := client.Database("quickbat")
	playerRepo := db.NewPlayerRepository(database.Collection("players"))
	challengeRepo := db.NewChallengeRepository(database.Collection("challenges"), playerRepo)

	domain.MainLeaderBoard = domain.NewLeaderBoard(playerRepo.FetchAll())

	// TODO move this to appropriate packages
	domain.AddChallengeChangeListener(challengeRepo.Update)
	domain.AddPlayerUpdateListener(playerRepo.Update)
	domain.AddPlayerUpdateListener(domain.MainLeaderBoard.UpdatePlayer)

	playerManager := usecase.PlayerManager{PlayerRepository: *playerRepo}
	challengerManager := usecase.ChallengeManager{ChallengeRepository: *challengeRepo, PlayerManager: playerManager}
	router := gin.Default()

	router.GET("/players/:id", func(ctx *gin.Context) {
		playerId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		playerDetails, err := usecase.GetPlayerDetails(ctx, playerId)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, &playerDetails)
	})

	router.POST("/matches", func(ctx *gin.Context) {
		var match usecase.Match
		err := ctx.BindJSON(&match)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		err = usecase.AddMatch(ctx, &match)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.Status(http.StatusCreated)
	})

	router.POST("/challenges", func(ctx *gin.Context) {
		var challenge usecase.ChallengeCreatePayload
		err := ctx.BindJSON(&challenge)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		createdChallenge, err := challengerManager.CreateChallenge(ctx, challenge)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusCreated, usecase.NewChallengeInfo(createdChallenge))
	})

	router.GET("/players/:id/challenges", func(ctx *gin.Context) {
		playerId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		challenges, err := challengerManager.FindChallengsForPlayer(ctx, playerId)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusCreated, challenges)
	})

	router.POST("/challenges/:id/accept", func(ctx *gin.Context) {
		challengeId := ctx.Param("id")
		var challengeAccept usecase.ChallengeAcceptPayload
		err := ctx.BindJSON(&challengeAccept)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		err = challengerManager.AcceptChallenge(ctx, challengeId, challengeAccept)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	})

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
