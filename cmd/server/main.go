package main

import (
	"time"
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
)

func main() {
	client := db.CreateConnection()
	defer db.CloseConnection()

	database := client.Database("quickbat")
	challengeRepo := db.NewChallengeRepository(database.Collection("challenges"))
	playerRepo := db.NewPlayerRepository(database.Collection("players"))

	domain.MainLeaderBoard = domain.NewLeaderBoard(playerRepo.FetchAll())
	domain.OnPlayerChange = playerRepo.Update

	domain.OnChallengeCreate = challengeRepo.Add
	domain.OnChallengeChange = challengeRepo.Update
	domain.LoadChallenge = challengeRepo.Find

	player1 := domain.GetLeaderBoard().FindPlayer(1)
	player2 := domain.GetLeaderBoard().FindPlayer(2)

	challenge := player1.Challenge(player2)

	now := time.Now()
	player2.Accept(challenge, now)

	challenge.WonBy(player2)

	//router := gin.Default()
	//
	//router.GET("/players/:id", func(ctx *gin.Context) {
	//	playerId, _ := strconv.Atoi(ctx.Param("id"))
	//	playerDetails := usecase.GetPlayerDetails(playerId)
	//	ctx.JSON(http.StatusOK, &playerDetails)
	//})
	//
	//router.POST("/matches", func(ctx *gin.Context) {
	//	var match usecase.Match
	//	ctx.BindJSON(&match)
	//	usecase.AddMatch(&match)
	//	ctx.Status(http.StatusCreated)
	//})
	//
	//router.Run("localhost:8080")
}
