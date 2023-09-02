package domain

import "time"

type Challenge struct {
	id         any
	challenger Player
	opponent   Player
	winner     Player
	isAccepted bool
	time       *time.Time
}

func (c *Challenge) SetId(value any) {
	c.id = value
}

func (c *Challenge) GetId() any {
	return c.id
}

func (c *Challenge) Time() *time.Time {
	return c.time
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


func newChallenge(challenger Player, opponent Player) *Challenge {
	challenge := Challenge{challenger: challenger, opponent: opponent}
	publishChallengeCreate(challenge)
	return &challenge
}

func (c *Challenge) winBy(winner Player) {
	c.winner = winner
	publishChallengeUpdate(*c)
}

func (c *Challenge) acceptBy(acceptedBy Player, agreedTime time.Time) {
	if acceptedBy.id != c.opponent.id {
		panic("challenge can not be accepted by someone other than opponent")
	}
	c.isAccepted = true
	c.time = &agreedTime
	publishChallengeUpdate(*c)
}

func LoadChallenge(id any, challenger Player, opponent Player, winner Player, isAccepted bool, time *time.Time) *Challenge {
	return &Challenge{
		id:         id,
		challenger: challenger,
		opponent:   opponent,
		winner:     winner,
		isAccepted: isAccepted,
		time:       time,
	}
}
