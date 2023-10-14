package usecase

import (
	"context"
	"tv/quick-bat/internal/domain"
)

type Match struct {
	ThisPlayerId  int  `json:"thisPlayerId"`
	OtherPlayerId int  `json:"otherPlayerId"`
	Win           bool `json:"win"`
}

type PlayerDetails struct {
	Id     int
	Name   string
	Wins   int
	Losses int
	Rank   int
}

func AddMatch(ctx context.Context, match *Match) error {

	thisPlayer, err := findPlayerById(match.ThisPlayerId)
	if err != nil {
		return err
	}

	otherPlayer, err := findPlayerById(match.OtherPlayerId)
	if err != nil {
		return err
	}

	if match.Win {
		thisPlayer.WinAgainst(&otherPlayer)
	} else {
		otherPlayer.WinAgainst(&thisPlayer)
	}

	return nil
}

func GetPlayerDetails(ctx context.Context, playerId int) (PlayerDetails, error) {
	player, err := findPlayerById(playerId)
	if err != nil {
		return PlayerDetails{}, err
	}
	leaderBoard := domain.GetLeaderBoard()
	return PlayerDetails{
		Id:     playerId,
		Name:   player.Name(),
		Wins:   player.Wins(),
		Losses: player.Losses(),
		Rank:   leaderBoard.GetRank(player),
	}, nil
}

func findPlayerById(id int) (domain.Player, error) {
	playerId := domain.PlayerId(id)
	leaderBoard := domain.GetLeaderBoard()
	return leaderBoard.FindPlayer(playerId)
}
