package domain

import (
	"errors"
	"time"
)

type Challenge struct {
	id         any
	challenger Player
	opponent   Player
	winner     Player
	isAccepted bool
	matchTime  *time.Time
}

func (c *Challenge) SetId(value any) {
	c.id = value
}

func (c *Challenge) GetId() any {
	return c.id
}

func (c *Challenge) Time() *time.Time {
	return c.matchTime
}

func (c *Challenge) IsAccepted() bool {
	return c.isAccepted
}

func (c *Challenge) Opponent() Player {
	return c.opponent
}

func (c *Challenge) Challenger() Player {
	return c.challenger
}

func (c *Challenge) Winner() Player {
	return c.winner
}

func newChallenge(challenger Player, opponent Player) Challenge {
	challenge := Challenge{challenger: challenger, opponent: opponent}
	return challenge
}

func (c *Challenge) winBy(winner Player) {
	c.winner = winner
	publishChallengeUpdate(*c)
}

func (c *Challenge) acceptBy(acceptedBy Player, agreedTime time.Time) error {
	if acceptedBy.id != c.opponent.id {
		return errors.New("challenge can not be accepted by someone other than opponent")
	}
	c.isAccepted = true
	c.matchTime = &agreedTime
	publishChallengeUpdate(*c)
	return nil
}

func LoadChallenge(id any, challenger Player, opponent Player, winner Player, isAccepted bool, matchTime *time.Time) *Challenge {
	return &Challenge{
		id:         id,
		challenger: challenger,
		opponent:   opponent,
		winner:     winner,
		isAccepted: isAccepted,
		matchTime:  matchTime,
	}
}
