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

type ChallengeManager struct {
	ChallengeRepository db.ChallengeRepository
	PlayerManager       PlayerManager
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

func (c ChallengeManager) FindChallengsForPlayer(ctx context.Context, playerId int) ([]domain.Challenge, error) {
	challengeRecords, err := c.ChallengeRepository.FindChallengesForPlayer(ctx, playerId)
	if err != nil {
		return nil, err
	}
	challenges := make([]domain.Challenge, len(challengeRecords))
	for i, r := range challengeRecords {
		challenge , err := c.loadChallenge(ctx, r)
		if err != nil {
			return nil, err
		}
		challenges[i] = *challenge
	}
	return challenges, nil
}

func (c ChallengeManager) AcceptChallenge(ctx context.Context, challengeId any, accept ChallengeAcceptPayload) error {
	record, err := c.ChallengeRepository.FindChallenge(ctx, challengeId)
	if err != nil {
		return err
	}
	challenge, err := c.loadChallenge(ctx, record)
	if err != nil {
		return err
	}
	challenger := challenge.Challenger()
	return challenger.Accept(challenge, accept.MatchTime)
}

func (c ChallengeManager) loadChallenge(ctx context.Context, record db.ChallengeRecord) (*domain.Challenge, error) {
	opponent, err := c.PlayerManager.FindPlayer(ctx, record.OpponentId)
	if err != nil {
		return nil,  err
	}
	challenger, err := c.PlayerManager.FindPlayer(ctx, record.ChallengerId)
	if err != nil {
		return nil, err
	}

	var winner domain.Player
	if record.WinnerId != 0 {
		winner, err = c.PlayerManager.FindPlayer(ctx, record.WinnerId)
		if err != nil {
			return  nil, err
		}
	}

	challenge := domain.LoadChallenge(record.Id, challenger, opponent, winner, record.IsAccepted, record.Time)
	return challenge, nil
}
