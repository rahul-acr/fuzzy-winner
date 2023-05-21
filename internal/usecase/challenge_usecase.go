package usecase

import (
	"time"
	"tv/quick-bat/internal/domain"
)

type Challenge struct {
	ChallengerId int `json:"challengerId"`
	OpponentId   int `json:"opponentId"`
}

func CreateChallenge(c Challenge) {
	challenger := findPlayerById(c.ChallengerId)
	opponent := findPlayerById(c.OpponentId)
	challenger.Challenge(opponent)
}

func AcceptChallenge(challengeId interface{}, opponentId int, matchTime time.Time) {
	challenge := domain.LoadChallenge(challengeId)
	opponent := findPlayerById(opponentId)
	opponent.Accept(challenge, matchTime)
}
