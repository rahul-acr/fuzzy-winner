package domain

type PlayerId int

type Player struct {
	id     PlayerId
	wins   int
	losses int
}

func (player *Player) winAgainst(loser *Player) {
	player.wins += 1
	loser.losses += 1
}
