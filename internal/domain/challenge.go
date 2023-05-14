package domain

import "time"

type Challenge struct {
	challenger *Player
	opponent   *Player
	isAccepted bool
	time       time.Time
}

func (c *Challenge) wonBy(winner *Player) {
	var loser *Player
	if winner.id == c.challenger.id {
		loser = c.opponent
	} else {
		loser = c.challenger
	}
	winner.WinAgainst(loser)
}
