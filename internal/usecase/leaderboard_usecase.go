package usecase

import "tv/quick-bat/internal/domain"

type Match struct {
	thisPlayerId  int
	otherPlayerId int
	win           bool
}

type PlayerDetails struct {
	Id     int
	Wins   int
	Losses int
	Rank   int
}

func AddMatch(match *Match) {
	thisPlayer := findPlayerById(match.thisPlayerId)
	otherPlayer := findPlayerById(match.otherPlayerId)

	if match.win {
		thisPlayer.WinAgainst(otherPlayer)
	} else {
		otherPlayer.WinAgainst(thisPlayer)
	}
}

func GetPlayerDetails(playerId int) *PlayerDetails {
	player := findPlayerById(playerId)
	return &PlayerDetails{
		Id:     playerId,
		Wins:   player.Wins(),
		Losses: player.Losses(),
		Rank:   domain.TtLeaderBoard.GetRank(player),
	}
}

func findPlayerById(id int) *domain.Player {
	playerId := domain.PlayerId(id)
	player := domain.TtLeaderBoard.FindPlayer(playerId)
	return player
}
