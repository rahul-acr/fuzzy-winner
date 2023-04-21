package usecase

import (
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
)

type Match struct {
	ThisPlayerId  int  `json:"thisPlayerId"`
	OtherPlayerId int  `json:"otherPlayerId"`
	Win           bool `json:"win"`
}

type PlayerDetails struct {
	Id     int
	Wins   int
	Losses int
	Rank   int
}

func AddMatch(match *Match) {
	thisPlayer := findPlayerById(match.ThisPlayerId)
	otherPlayer := findPlayerById(match.OtherPlayerId)

	if match.Win {
		thisPlayer.WinAgainst(otherPlayer)
	} else {
		otherPlayer.WinAgainst(thisPlayer)
	}

	db.UpdatePlayer(thisPlayer)
	db.UpdatePlayer(otherPlayer)
}

func GetPlayerDetails(playerId int) *PlayerDetails {
	player := findPlayerById(playerId)
	leaderBoard := domain.GetLeaderBoard()
	return &PlayerDetails{
		Id:     playerId,
		Wins:   player.Wins(),
		Losses: player.Losses(),
		Rank:   leaderBoard.GetRank(player),
	}
}

func findPlayerById(id int) *domain.Player {
	playerId := domain.PlayerId(id)
	leaderBoard := domain.GetLeaderBoard()
	player := leaderBoard.FindPlayer(playerId)
	return player
}
