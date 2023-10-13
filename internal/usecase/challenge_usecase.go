package usecase

import (
	"context"
	"time"
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
)

type ChallengeCreatePayload struct {
	ChallengerId int `json:"challengerId"`
	OpponentId   int `json:"opponentId"`
}

type ChallengeAcceptPayload struct {
	OpponentId int       `json:"opponentId"`
	MatchTime  time.Time `json:"matchTime"`
}

type ChallengeResult struct {
	WinnerId int `json:"winnerId"`
}

type ChallengeInfo struct {
	Id         any        `json:"id"`
	Challenger int        `json:"challengerId"`
	Opponent   int        `json:"opponentId"`
	Winner     int        `json:"winnerId,omitempty"`
	IsAccepted bool       `json:"isAccepted"`
	MatchTime  *time.Time `json:"matchTime"`
}

type ChallengeManager struct {
	ChallengeRepository db.ChallengeRepository
	PlayerManager       PlayerManager
}

func NewChallengeInfo(c domain.Challenge) ChallengeInfo {
	challenger := c.Challenger()
	opponent := c.Opponent()

	cInfo := ChallengeInfo{
		Id:         c.Id,
		Challenger: int(challenger.Id()),
		Opponent:   int(opponent.Id()),
	}

	winner := c.Winner()
	if (winner != domain.Player{}) {
		cInfo.Winner = int(winner.Id())
	}

	cInfo.MatchTime = c.Time()
	cInfo.IsAccepted = c.IsAccepted()

	return cInfo
}

func (c ChallengeManager) CreateChallenge(ctx context.Context, challenge ChallengeCreatePayload) (domain.Challenge, error) {
	challenger, err := c.PlayerManager.FindPlayer(ctx, challenge.ChallengerId)
	if err != nil {
		return domain.Challenge{}, err
	}
	opponent, err := c.PlayerManager.FindPlayer(ctx, challenge.OpponentId)
	if err != nil {
		return domain.Challenge{}, err
	}
	return c.ChallengeRepository.Add(ctx, challenger.Challenge(opponent))
}

func (c ChallengeManager) FindChallengsForPlayer(ctx context.Context, playerId int) ([]ChallengeInfo, error) {
	challenges, err := c.ChallengeRepository.FindChallengesForPlayer(ctx, playerId)
	if err != nil {
		return nil, err
	}
	challengeInfos := make([]ChallengeInfo, len(challenges))
	for i, r := range challenges {
		challengeInfos[i] = NewChallengeInfo(r)
	}
	return challengeInfos, nil
}

func (c ChallengeManager) AcceptChallenge(ctx context.Context, challengeId any, accept ChallengeAcceptPayload) error {
	challenge, err := c.ChallengeRepository.FindChallenge(ctx, challengeId)
	if err != nil {
		return err
	}
	opponent, err := c.PlayerManager.FindPlayer(ctx, accept.OpponentId)
	if err != nil {
		return err
	}
	return opponent.Accept(&challenge, accept.MatchTime)
}

func (c ChallengeManager) AddChallengeResult(ctx context.Context, challengeId string, result ChallengeResult) error {
	challenge, err := c.ChallengeRepository.FindChallenge(ctx, challengeId)
	if err != nil {
		return err
	}
	winner, err := c.PlayerManager.FindPlayer(ctx, result.WinnerId)
	if err != nil {
		return err
	}
	return winner.Win(&challenge)
}
