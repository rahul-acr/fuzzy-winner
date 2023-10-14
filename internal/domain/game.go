package domain

import (
	"time"
)

type PlayerId int

type Player struct {
	id     PlayerId
	name   string
	wins   int
	losses int
}

func (player *Player) WinAgainst(loser *Player) {
	player.wins += 1
	loser.losses += 1

	publishPlayerUpdate(*player)
	publishPlayerUpdate(*loser)
}

func NewPlayer(id PlayerId, name string, wins, looses int) Player {
	return Player{id, name, wins, looses}
}

func (player *Player) Name() string {
	return player.name
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

func (player *Player) Challenge(opponent Player) Challenge {
	return newChallenge(*player, opponent)
}

func (player *Player) Accept(challenge *Challenge, agreedTime time.Time) error {
	return challenge.acceptBy(*player, agreedTime)
}

func (player *Player) Win(challenge *Challenge) error {
	player.WinAgainst(&challenge.opponent)
	return challenge.winBy(*player)
}
