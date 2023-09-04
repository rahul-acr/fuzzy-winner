package usecase

import (
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

func (c ChallengeManager) CreateChallenge(challenge Challenge) (domain.Challenge, error) {
	challenger, err := c.PlayerManager.FindPlayer(challenge.ChallengerId)
	if err != nil {
		return domain.Challenge{}, err
	}
	opponent, err := c.PlayerManager.FindPlayer(challenge.OpponentId)
	if err != nil {
		return domain.Challenge{}, err
	}
	return c.ChallengeRepository.Add(challenger.Challenge(opponent))
}

func (c ChallengeManager) AcceptChallenge(challengeId interface{}, accept ChallengeAccept) error {
	record, err := c.ChallengeRepository.FindChallenge(challengeId)
	if err != nil {
		return err
	}
	opponent, err := c.PlayerManager.FindPlayer(record.OpponentId)
	if err != nil {
		return err
	}
	challenger, err := c.PlayerManager.FindPlayer(record.ChallengerId)
	if err != nil {
		return err
	}

	var winner domain.Player
	if record.WinnerId != 0 {
		winner, err = c.PlayerManager.FindPlayer(record.WinnerId)
		if err != nil {
			return err
		}
	}

	challenge := domain.LoadChallenge(record.Id, challenger, opponent, winner, record.IsAccepted, record.Time)
	challenger.Accept(challenge, accept.MatchTime)
	return nil
}
