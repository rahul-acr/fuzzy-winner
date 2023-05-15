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
	onChallengeChange(c)
}

func (c *Challenge) acceptBy(acceptedBy *Player, agreedTime time.Time) {
	if acceptedBy.id != c.opponent.id {
		panic("challenge can not be accepted by someone other than opponent")
	}
	c.isAccepted = true
	c.time = agreedTime
	onChallengeChange(c)
}

var onChallengeCreate = func(c *Challenge) {}
var onChallengeChange = func(c *Challenge) {}
