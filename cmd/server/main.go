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
	matchRepo := db.NewMatchRepository(database.Collection("matches"))

	domain.MainLeaderBoard = domain.NewLeaderBoard(playerRepo.FetchAll())

	// TODO move this to appropriate packages
	domain.AddChallengeChangeListener(challengeRepo.Update)
	domain.AddPlayerUpdateListener(playerRepo.Update)
	domain.AddPlayerUpdateListener(domain.MainLeaderBoard.UpdatePlayer)

	playerManager := usecase.PlayerManager{PlayerRepository: *playerRepo}
	challengerManager := usecase.ChallengeManager{ChallengeRepository: *challengeRepo, PlayerManager: playerManager}
	matchManager := usecase.MatchManager{MatchRepo: matchRepo}

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/leaderboard", func(ctx *gin.Context) {
		playerDetails := matchManager.GetLeaderBoard()
		ctx.JSON(http.StatusOK, playerDetails)
	})

	router.GET("/players/:id", func(ctx *gin.Context) {
		playerId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		playerDetails, err := matchManager.GetPlayerDetails(playerId)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, &playerDetails)
	})

	router.POST("/matches", func(ctx *gin.Context) {
		var matchPayload usecase.MatchPayload
		err := ctx.BindJSON(&matchPayload)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		err = matchManager.AddMatch(ctx, matchPayload)
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
		ctx.JSON(http.StatusOK, challenges)
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

	router.POST("/challenges/:id/result", func(ctx *gin.Context) {
		challengeId := ctx.Param("id")
		var challegeResult usecase.ChallengeResult
		err := ctx.BindJSON(&challegeResult)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		err = challengerManager.AddChallengeResult(ctx, challengeId, challegeResult)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
		}
	})

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
