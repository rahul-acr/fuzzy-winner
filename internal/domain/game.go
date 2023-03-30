package domain

type PlayerId int

type Player struct {
	id     PlayerId
	wins   int
	losses int
}

type Match struct {
	firstPlayer  *Player
	secondPlayer *Player
	winner       *Player
}

func (match *Match) wonBy(winner *Player) {
	var loser *Player
	if winner == match.firstPlayer {
		match.winner, loser = match.firstPlayer, match.secondPlayer
	} else {
		match.winner, loser = match.secondPlayer, match.firstPlayer
	}
	winner.wins += 1
	loser.losses += 1
}

func Between(firstPlayer *Player, secondPlayer *Player) *Match {
	return &Match{firstPlayer, secondPlayer, nil}
}
