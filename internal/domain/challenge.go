package domain

import "time"

type Challenge struct {
	Id         interface{}
	challenger *Player
	opponent   *Player
	winner     *Player
	isAccepted bool
	time       *time.Time
}

func NewChallenge(challenger *Player, opponent *Player) *Challenge {
	challenge := &Challenge{challenger: challenger, opponent: opponent}
	publishChallengeCreate(challenge)
	return challenge
}

func (c *Challenge) WonBy(winner *Player) {
	var loser *Player
	if winner.id == c.challenger.id {
		loser = c.opponent
	} else {
		loser = c.challenger
	}
	winner.WinAgainst(loser)
	c.winner = winner
	publishChallengeUpdate(c)
}

func (c *Challenge) acceptBy(acceptedBy *Player, agreedTime time.Time) {
	if acceptedBy.id != c.opponent.id {
		panic("challenge can not be accepted by someone other than opponent")
	}
	c.isAccepted = true
	c.time = &agreedTime
	publishChallengeUpdate(c)
}

func (c *Challenge) Challenger() *Player {
	return c.challenger
}

func (c *Challenge) Opponent() *Player {
	return c.opponent
}

func (c *Challenge) Winner() *Player {
	return c.winner
}

func (c *Challenge) IsAccepted() bool {
	return c.isAccepted
}

func (c *Challenge) Time() *time.Time {
	return c.time
}

func LoadChallenge(id interface{}, challenger *Player, opponent *Player, winner *Player, isAccepted bool, time *time.Time) *Challenge {
	return &Challenge{
		Id:         id,
		challenger: challenger,
		opponent:   opponent,
		winner:     winner,
		isAccepted: isAccepted,
		time:       time,
	}
}
