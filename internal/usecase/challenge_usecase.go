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

type ChallengeInfo struct {
	Id         any        `json:"id"`
	Challenger int        `json:"challengerId"`
	Opponent   int        `json:"opponentId"`
	Winner     int        `json:"winnerId,omitempty"`
	IsAccepted bool       `json:"isAccepted"`
	MatchTime  *time.Time `json:"matchTime"`
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
	challengeRecords, err := c.ChallengeRepository.FindChallengesForPlayer(ctx, playerId)
	if err != nil {
		return nil, err
	}
	challenges := make([]ChallengeInfo, len(challengeRecords))
	for i, r := range challengeRecords {
		challenges[i] = ChallengeInfo{
			Id:         r.Id,
			Challenger: r.ChallengerId,
			Opponent:   r.OpponentId,
			Winner:     r.WinnerId,
			MatchTime:  r.Time,
			IsAccepted: r.IsAccepted,
		}
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
	opponent, err := c.PlayerManager.FindPlayer(ctx, accept.OpponentId)
	if err != nil {
		return err
	}
	return opponent.Accept(challenge, accept.MatchTime)
}

func (c ChallengeManager) loadChallenge(ctx context.Context, record db.ChallengeRecord) (*domain.Challenge, error) {
	opponent, err := c.PlayerManager.FindPlayer(ctx, record.OpponentId)
	if err != nil {
		return nil, err
	}
	challenger, err := c.PlayerManager.FindPlayer(ctx, record.ChallengerId)
	if err != nil {
		return nil, err
	}

	var winner domain.Player
	if record.WinnerId != 0 {
		winner, err = c.PlayerManager.FindPlayer(ctx, record.WinnerId)
		if err != nil {
			return nil, err
		}
	}

	challenge := domain.LoadChallenge(record.Id, challenger, opponent, winner, record.IsAccepted, record.Time)
	return challenge, nil
}
