package domain

import (
	"time"
)

type PlayerId int

type Player struct {
	id     PlayerId
	wins   int
	losses int
}

func (player *Player) WinAgainst(loser *Player) {
	player.wins += 1
	loser.losses += 1

	publishPlayerUpdate(player)
	publishPlayerUpdate(loser)
	MainLeaderBoard.refresh()
}

func NewPlayer(id PlayerId, wins, looses int) *Player {
	return &Player{id, wins, looses}
}

func (player *Player) Wins() int {
	return player.wins
}

func (player *Player) Losses() int {
	return player.losses
}

func (player *Player) Id() PlayerId {
	return player.id
}

func (player *Player) Challenge(opponent *Player) *Challenge {
	return NewChallenge(player, opponent)
}

func (player *Player) Accept(challenge *Challenge, agreedTime time.Time) {
	challenge.acceptBy(player, agreedTime)
}
