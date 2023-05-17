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

	domain.GetLeaderBoard().Init(playerRepo.FetchAll())
	domain.OnPlayerChange = func(p *domain.Player) {
		playerRepo.Update(p)
	}

	domain.OnChallengeCreate = func(c *domain.Challenge) {
		challengeRepo.Add(c)
	}

	domain.OnChallengeChange = func(c *domain.Challenge) {
		challengeRepo.Update(c)
	}

	//player1 := domain.GetLeaderBoard().FindPlayer(1)
	//player2 := domain.GetLeaderBoard().FindPlayer(2)
	//
	//challenge := player1.Challenge(player2)
	//
	//player2.Accept(challenge, time.Now())
	//
	//challenge.WonBy(player2)

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
