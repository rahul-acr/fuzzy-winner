package usecase

import (
	"time"
	"tv/quick-bat/internal/domain"
)

type Challenge struct {
	ChallengerId int `json:"challengerId"`
	OpponentId   int `json:"opponentId"`
}

var LoadChallenge = func(challengeId interface{}) *domain.Challenge { panic("Hook is not linked") }

func CreateChallenge(c Challenge) {
	challenger := findPlayerById(c.ChallengerId)
	opponent := findPlayerById(c.OpponentId)
	challenger.Challenge(opponent)
}

func AcceptChallenge(challengeId interface{}, opponentId int, matchTime time.Time) {
	challenge := LoadChallenge(challengeId)
	opponent := findPlayerById(opponentId)
	opponent.Accept(challenge, matchTime)
}
