package domain

type PlayerId int

type Player struct {
	id     PlayerId
	wins   int
	losses int
}

func (winner *Player) winAgainst(loser *Player) {
	winner.wins += 1
	loser.losses += 1
}
