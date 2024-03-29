package domain

import (
	"errors"
	"time"
)

// TODO: make it immutable ?
type Challenge struct {
	Id         any
	challenger Player
	opponent   Player
	winner     Player
	isAccepted bool
	matchTime  *time.Time
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

func (c *Challenge) winBy(winner Player) error {
	if c.winner != (Player{}) {
		return errors.New("WINNER IS ALREADY SET")
	}
	c.winner = winner
	publishChallengeUpdate(*c)
	return nil
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

func LoadChallenge(id any,
	challenger Player,
	opponent Player,
	winner Player,
	isAccepted bool,
	matchTime *time.Time,
) Challenge {
	return Challenge{
		Id:         id,
		challenger: challenger,
		opponent:   opponent,
		winner:     winner,
		isAccepted: isAccepted,
		matchTime:  matchTime,
	}
}
