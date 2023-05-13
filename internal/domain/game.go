package domain

import "time"

type PlayerId int

type Player struct {
	id     PlayerId
	wins   int
	losses int
}

func (player *Player) WinAgainst(loser *Player) {
	player.wins += 1
	loser.losses += 1
	leaderBoard.refresh()
}

func CreatePlayer(id PlayerId, wins, looses int) *Player {
	return &Player{id, wins, looses}
}

func (player *Player) Wins() int {
	return player.wins
}

func (player *Player) Losses() int {
	return player.losses
}

func (player *Player) Id() PlayerId {
	return player.id
}

func (player *Player) challenge(otherPlayer *Player) *Challenge {
	challenge := Challenge{challenger: player, opponent: otherPlayer}
	return &challenge
}

func (player *Player) accept(challenge *Challenge, agreedTime time.Time) {
	challenge.isAccepted = true
	challenge.time = agreedTime
}
