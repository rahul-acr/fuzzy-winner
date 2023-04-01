package domain

import "sort"

type LeaderBoard struct {
	players []*Player
}

var leaderBoard LeaderBoard

func (board *LeaderBoard) Len() int {
	return len(board.players)
}

func (board *LeaderBoard) Less(i, j int) bool {
	return board.players[i].wins > board.players[j].wins
}

func (board *LeaderBoard) Swap(i, j int) {
	board.players[i], board.players[j] = board.players[j], board.players[i]
}

func (board *LeaderBoard) Init(players []*Player) {
	board.players = players
}

func (board *LeaderBoard) Update() {
	sort.Sort(board)
}

func (board *LeaderBoard) GetRank(player *Player) int {
	for i := 0; i < board.Len(); i++ {
		if player.id == board.players[i].id {
			return i + 1
		}
	}
	return -1
}
