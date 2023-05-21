package usecase

import (
	"time"
	"tv/quick-bat/internal/domain"
)

type Challenge struct {
	Challenger int `json:"challengerId"`
	Opponent   int `json:"opponentId"`
}

func CreateChallenge(c Challenge) {
	challenger := findPlayerById(c.Challenger)
	opponent := findPlayerById(c.Opponent)
	challenger.Challenge(opponent)
}

func AcceptChallenge(challengeId int, opponentId int, matchTime time.Time) {
	challenge := domain.LoadChallenge(challengeId)
	opponent := findPlayerById(opponentId)
	opponent.Accept(challenge, matchTime)
}
