package usecase

import (
	"context"
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
)

type MatchPayload struct {
	ThisPlayerId  int  `json:"thisPlayerId"`
	OtherPlayerId int  `json:"otherPlayerId"`
	Win           bool `json:"win"`
}

type PlayerDetails struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Rank   int    `json:"rank"`
}

type MatchManager struct {
	MatchRepo db.MatchRepository
}

func (m MatchManager) AddMatch(ctx context.Context, matchPayload MatchPayload) error {
	thisPlayer, err := m.findPlayerById(matchPayload.ThisPlayerId)
	if err != nil {
		return err
	}

	otherPlayer, err := m.findPlayerById(matchPayload.OtherPlayerId)
	if err != nil {
		return err
	}

	var match domain.Match
	if matchPayload.Win {
		thisPlayer.WinAgainst(&otherPlayer)
		match = domain.Match{Winner: thisPlayer, Loser: otherPlayer}
	} else {
		otherPlayer.WinAgainst(&thisPlayer)
		match = domain.Match{Loser: thisPlayer, Winner: otherPlayer}
	}

	match, err = m.MatchRepo.Add(ctx, match)
	if err != nil {
		return err
	}

	return nil
}

func (m MatchManager) GetPlayerDetails(playerId int) (PlayerDetails, error) {
	player, err := m.findPlayerById(playerId)
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

func (m MatchManager) findPlayerById(id int) (domain.Player, error) {
	playerId := domain.PlayerId(id)
	leaderBoard := domain.GetLeaderBoard()
	return leaderBoard.FindPlayer(playerId)
}

func (m MatchManager) GetLeaderBoard() []PlayerDetails {
	topPlayers := domain.GetLeaderBoard().Players()
	playerDetails := make([]PlayerDetails, len(topPlayers))
	for rank, player := range topPlayers {
		playerDetails[rank] = PlayerDetails{
			Id:     int(player.Id()),
			Name:   player.Name(),
			Wins:   player.Wins(),
			Losses: player.Losses(),
			Rank:   rank + 1,
		}
	}
	return playerDetails
}
