package usecase

import (
	"context"
	"time"
	"tv/quick-bat/internal/db"
	"tv/quick-bat/internal/domain"
)

type Challenge struct {
	ChallengerId int `json:"challengerId"`
	OpponentId   int `json:"opponentId"`
}

type ChallengeAccept struct {
	OpponentId int       `json:"opponentId"`
	MatchTime  time.Time `json:"matchTime"`
}

type ChallengeManager struct {
	ChallengeRepository db.ChallengeRepository
	PlayerManager       PlayerManager
}

func (c ChallengeManager) CreateChallenge(ctx context.Context, challenge Challenge) (domain.Challenge, error) {
	challenger, err := c.PlayerManager.FindPlayer(ctx, challenge.ChallengerId)
	if err != nil {
		return domain.Challenge{}, err
	}
	opponent, err := c.PlayerManager.FindPlayer(ctx, challenge.OpponentId)
	if err != nil {
		return domain.Challenge{}, err
	}
	return c.ChallengeRepository.Add(challenger.Challenge(opponent))
}

func (c ChallengeManager) AcceptChallenge(ctx context.Context, challengeId any, accept ChallengeAccept) error {
	record, err := c.ChallengeRepository.FindChallenge(challengeId)
	if err != nil {
		return err
	}
	opponent, err := c.PlayerManager.FindPlayer(ctx, record.OpponentId)
	if err != nil {
		return err
	}
	challenger, err := c.PlayerManager.FindPlayer(ctx, record.ChallengerId)
	if err != nil {
		return err
	}

	var winner domain.Player
	if record.WinnerId != 0 {
		winner, err = c.PlayerManager.FindPlayer(ctx, record.WinnerId)
		if err != nil {
			return err
		}
	}

	challenge := domain.LoadChallenge(record.Id, challenger, opponent, winner, record.IsAccepted, record.Time)
	return challenger.Accept(challenge, accept.MatchTime)
}
