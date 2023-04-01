package domain

type PlayerId int

type Player struct {
	id     PlayerId
	wins   int
	losses int
}

func (winner *Player) WinAgainst(loser *Player) {
	winner.wins += 1
	loser.losses += 1
	leaderBoard.Update()
}

func (player *Player) GetRank() int {
	return leaderBoard.GetRank(player)
}
