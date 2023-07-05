package usecase

import (
	"time"
	"tv/quick-bat/internal/domain"
)

type Challenge struct {
	ChallengerId int `json:"challengerId"`
	OpponentId   int `json:"opponentId"`
}

type ChallengeAccept struct {
	OpponentId int    `json:"opponentId"`
	MatchTime  string `json:"matchTime"`
}

var LoadChallenge = func(challengeId interface{}) (*domain.Challenge, error) { panic("Hook is not linked") }

func CreateChallenge(c Challenge) {
	challenger := findPlayerById(c.ChallengerId)
	opponent := findPlayerById(c.OpponentId)
	challenger.Challenge(opponent)
}

func AcceptChallenge(challengeId interface{}, accept ChallengeAccept) error {
	challenge, err := LoadChallenge(challengeId)
	if err != nil {
		return err
	}
	matchTime, err := time.Parse(time.RFC3339, accept.MatchTime)
	if err != nil {
		return err
	}
	challenge.Opponent().Accept(challenge, matchTime)
	return nil
}
