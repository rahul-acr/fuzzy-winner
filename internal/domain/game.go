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
		match.winner = match.firstPlayer
		loser = match.secondPlayer
	} else {
		match.winner = match.secondPlayer
		loser = match.firstPlayer
	}
	winner.wins += 1
	loser.losses += 1
}

func Between(firstPlayer *Player, secondPlayer *Player) Match {
	return Match{firstPlayer, secondPlayer, nil}
}
