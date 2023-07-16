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
	OpponentId int       `json:"opponentId"`
	MatchTime  time.Time `json:"matchTime"`
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
	opponent := challenge.Opponent()
	opponent.Accept(challenge, accept.MatchTime)
	return nil
}
