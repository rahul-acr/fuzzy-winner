package usecase

import "tv/quick-bat/internal/domain"

type Match struct {
	thisPlayerId  int
	otherPlayerId int
	win           bool
}

func AddMatch(match Match) {
	thisPlayer := findPlayerById(match.thisPlayerId)
	otherPlayer := findPlayerById(match.otherPlayerId)

	if match.win {
		thisPlayer.WinAgainst(otherPlayer)
	} else {
		otherPlayer.WinAgainst(thisPlayer)
	}
}

func findPlayerById(id int) *domain.Player {
	playerId := domain.PlayerId(id)
	player := domain.TtLeaderBoard.FindPlayer(playerId)
	return player
}
