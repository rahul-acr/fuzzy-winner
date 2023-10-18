package usecase

import (
	"context"
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
)

type MatchPayload struct {
	MatchId  string `json:"matchId"`
	WinnerId int    `json:"winnerId"`
	LoserId  int    `json:"loserId"`
}

type MatchData struct {
	MatchId    any    `json:"matchId"`
	WinnerId   int    `json:"winnerId"`
	WinnerName string `json:"winnerName"`
	LoserId    int    `json:"loserId"`
	LoserName  string `json:"loserName"`
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
	winner, err := m.findPlayerById(matchPayload.WinnerId)
	if err != nil {
		return err
	}

	loser, err := m.findPlayerById(matchPayload.LoserId)
	if err != nil {
		return err
	}

	winner.WinAgainst(&loser)
	match := domain.Match{Winner: winner, Loser: loser}

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

func (m MatchManager) FetchMatchesOfPlayer(ctx context.Context, playerId int) ([]MatchData, error) {
	matches, err := m.MatchRepo.FindMatchesOf(ctx, domain.PlayerId(playerId))
	if err != nil {
		return nil, err
	}
	matchData := make([]MatchData, len(matches))
	for i, match := range matches {
		matchData[i] = MatchData{
			MatchId:    match.Id,
			WinnerId:   int(match.Winner.Id()),
			WinnerName: match.Winner.Name(),
			LoserId:    int(match.Loser.Id()),
			LoserName:  match.Loser.Name(),
		}
	}

	return matchData, nil
}
